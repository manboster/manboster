package telegram

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
	"gopkg.in/telebot.v3"
)

func recipientParser(chatID string) (telebot.ChatID, error) {
	id, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return telebot.ChatID(0), fmt.Errorf("invalid chat id: %w", err)
	}

	// define recipients
	recp := telebot.ChatID(id)
	return recp, nil
}

// build up a message from the ground up
func (s *Service) msgParser(msg *chat.Message, m *telebot.Message) error {
	// define chat type
	var chatType chat.ChatsType
	switch m.Chat.Type {
	case telebot.ChatGroup:
		chatType = chat.ChatsGroup
	case telebot.ChatSuperGroup:
		chatType = chat.ChatsGroup
	case telebot.ChatChannel:
		chatType = chat.ChatsChannel
	case telebot.ChatPrivate:
		chatType = chat.ChatsPersonal
	default:
		chatType = chat.ChatsUnknown
	}
	msg.ChatType = chatType

	text := m.Text
	if m.Entities != nil && len(m.Entities) > 0 {
		for _, e := range m.Entities {
			substr := m.EntityText(e)
			switch e.Type {
			case telebot.EntityCommand:
				// process "/xxxxxx xxxx" and "/xxxx@xxxxbot xxxxxxx"
				var command string
				var args []string
				// if it is bot's username
				if strings.Contains(substr, "@"+s.tgInstance.Me.Username) {
					command = strings.ToLower(strings.Replace(substr, "@"+s.tgInstance.Me.Username, "", -1))[1:]
					args = strings.Split(text, " ")[1:]
				} else {
					command = substr[1:]
					if len(strings.Split(text, " ")) > 1 {
						args = strings.Split(text, " ")[1:]
					}
				}
				msg.MessageType = chat.MessageCommand
				msg.Command = &chat.CommandPayload{
					CommandType: chat.CommandType(command),
					CommandArgs: args,
				}
				return nil
			case telebot.EntityBold:
				text = strings.Replace(text, substr, "**"+substr+"**", -1)
			case telebot.EntityItalic:
				text = strings.Replace(text, substr, "*"+substr+"*", -1)
			case telebot.EntityCode:
				text = strings.Replace(text, substr, "`"+substr+"`", -1)
			case telebot.EntityBlockquote:
				text = strings.Replace(text, substr, "> "+substr, -1)
			case telebot.EntityCodeBlock:
				text = strings.Replace(text, substr, "```"+e.Language+"\n"+substr+"```", -1)
			case telebot.EntitySpoiler:
				text = strings.Replace(text, substr, "|| "+substr+" ||", -1)
			case telebot.EntityMention:
				if e.User != nil {
					uName := e.User.FirstName + " " + e.User.LastName + " " + e.User.Username
					text = strings.Replace(text, substr, fmt.Sprintf("[[!@{%s,%s-%d}]]", uName, s.Name(), e.User.ID), -1)
				}
			case telebot.EntityStrikethrough:
				text = strings.Replace(text, substr, "~~"+substr+"~~", -1)
			case telebot.EntityUnderline:
				text = strings.Replace(text, substr, "_"+substr+"_", -1)
			case telebot.EntityTextLink:
				text = strings.Replace(text, substr, "["+substr+"]("+e.URL+")", -1)
			}
		}
	}

	msg.MessageType = chat.MessageText
	if m.Text != "" {
		msg.Text = &chat.TextPayload{
			Text: m.Text,
		}
		msg.MessageType |= chat.MessageText
	}

	if m.Photo != nil {
		if m.Photo.Caption != "" {
			msg.Text = &chat.TextPayload{
				Text: m.Text,
			}
			msg.MessageType |= chat.MessageText
		}
	}

	if m.Video != nil {
		if m.Video.Caption != "" {
			msg.Text = &chat.TextPayload{
				Text: m.Video.Caption,
			}
			msg.MessageType |= chat.MessageText
		}
	}

	if m.Audio != nil {
		if m.Audio.Caption != "" {
			msg.Text = &chat.TextPayload{
				Text: m.Audio.Caption,
			}
			msg.MessageType |= chat.MessageText
		}
	}

	if m.Document != nil {
		if m.Document.Caption != "" {
			msg.Text = &chat.TextPayload{
				Text: m.Document.Caption,
			}
			msg.MessageType |= chat.MessageText
		}
	}

	if m.Sticker != nil {
		if m.Sticker.Emoji != "" {
			msg.Text = &chat.TextPayload{
				Text: "[A sticker with emoji " + m.Sticker.Emoji + "]",
			}
			msg.MessageType |= chat.MessageText
		}
	}

	if msg.Text != nil {
		msg.Text.Text = strings.Replace(msg.Text.Text, "@"+s.tgInstance.Me.Username, "[[!@{Assistant}]]", -1)
	}

	// fmt.Printf(m.Text)
	// handle reply to quote message
	if msg.Reply != nil && m.Quote != nil {
		msg.Reply.Text = &chat.TextPayload{
			Text: m.Quote.Text,
		}
	}

	return nil
}

func (s *Service) msgBaseParser(msg *chat.Message, c telebot.Context) {
	// define things all we know.
	msg.MessageID = fmt.Sprintf("%d", c.Message().ID)
	msg.Username = c.Sender().FirstName + " " + c.Sender().LastName
	if c.Sender().Username != "" {
		msg.Username += "(" + c.Sender().Username + ")"
	}
	msg.ChatName = c.Chat().Title // Only Group available
	msg.UserID = fmt.Sprintf("%d", c.Sender().ID)
	msg.ChatID = fmt.Sprintf("%d", c.Chat().ID)
	msg.CreatedAt = time.Now()
	msg.Provider = s.Name()

	// check whether replies available or not
	if c.Message().ReplyTo != nil {
		msg.Reply = &chat.Message{}
		msg.Reply.Username = c.Message().ReplyTo.Sender.FirstName + " " + c.Message().ReplyTo.Sender.LastName
		if s.tgInstance.Me.ID == c.Message().ReplyTo.Sender.ID {
			msg.Reply.Username = "Assistant"
		}
		msg.Reply.MessageID = fmt.Sprintf("%d", c.Message().ReplyTo.ID)
		msg.Reply.UserID = fmt.Sprintf("%d", c.Message().ReplyTo.Sender.ID)
		msg.Reply.ChatID = fmt.Sprintf("%d", c.Message().ReplyTo.Chat.ID)
		msg.Reply.CreatedAt = c.Message().ReplyTo.Time()
		err := s.msgParser(msg.Reply, c.Message().ReplyTo)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Failed to parse message data: %q", err.Error()))
		}
	}

	// check whether message forward available or not
	if c.Message().IsForwarded() {
		msg.Forward = &chat.Message{}
		if c.Message().Origin != nil {
			msg.Forward.Username = c.Message().Origin.SenderUsername
			if c.Message().Origin.Sender != nil {
				msg.Forward.Username = c.Message().Origin.Sender.FirstName + " " + c.Message().Origin.Sender.LastName
				if c.Message().Origin.Sender.Username != "" {
					msg.Forward.Username += "(" + c.Message().Origin.Sender.Username + ")"
				}
				if s.tgInstance.Me.ID == c.Message().Origin.Sender.ID {
					msg.Forward.Username = "Assistant"
				}
				msg.Forward.UserID = fmt.Sprintf("%d", c.Message().Origin.Sender.ID)
			}
			if c.Message().Origin.Chat != nil {
				msg.Forward.ChatName = c.Message().Origin.Chat.Title
			}
		}
	}

	// process sender data
	if c.Message().Origin != nil && c.Message().Origin.SenderUsername != "" {
		if msg.Forward == nil {
			msg.Forward = &chat.Message{}
		}
		msg.Forward.Username = c.Message().Origin.SenderUsername
	}

	// parse message data
	err := s.msgParser(msg, c.Message())
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Failed to parse message data: %q", err.Error()))
	}
}
