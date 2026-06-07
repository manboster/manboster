package telegram

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	manbofs "github.com/manboster/manboster/internal/fs"
	"github.com/manboster/manboster/internal/util"
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
func (s *Service) msgParser(msg *chat.Message, m *telebot.Message, onMsg func(msg *chat.Message)) error {
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
			Text: text,
		}
		msg.MessageType |= chat.MessageText
	}

	if m.Caption != "" {
		msg.Text = &chat.TextPayload{
			Text: text,
		}
		msg.MessageType |= chat.MessageText
	}

	if m.Document != nil {
		reader, err := s.tgInstance.File(&m.Document.File)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Failed to base64 image photo: %v", err))
		} else {
			content := util.RandomString(8)
			err = manbofs.Write(content, reader)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Failed to base64 image photo: %v", err))
			} else {
				msg.MessageType |= chat.MessageFile
				msg.File = &chat.FilePayload{
					Content: []string{
						content,
					},
				}
			}
		}
	}

	if m.Photo != nil {
		reader, err := s.tgInstance.File(&m.Photo.File)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Failed to base64 image photo: %v", err))
		} else {
			content := util.RandomString(8)
			err = manbofs.Write(content, reader)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Failed to base64 image photo: %v", err))
			} else {
				msg.MessageType |= chat.MessageImage
				msg.Image = &chat.ImagePayload{
					Content: []string{
						content,
					},
				}
			}
		}
	}

	if m.AlbumID != "" {
		id := m.AlbumID
		_, _ = SetTimer(id, onMsg)
		ms, avail := GetData(id)
		if !avail {
			SetData(id, msg)
		} else {
			if ms.Image != nil {
				ms.Image.Content = append(ms.Image.Content, msg.Image.Content...)
			} else if msg.Image.Content != nil {
				ms.Image.Content = msg.Image.Content
			}

			if m.Caption != "" {
				if ms.Text == nil {
					ms.Text = &chat.TextPayload{
						Text: m.Caption,
					}
				} else {
					ms.Text.Text += "\n" + m.Caption
				}
			}

			SetData(id, ms)
		}
		return ErrImageNoNeedToTrigger
	}

	if m.Sticker != nil {
		if m.Sticker.Emoji != "" {
			msg.Text = &chat.TextPayload{
				Text: "[A sticker with emoji " + m.Sticker.Emoji + "]",
			}
			msg.MessageType |= chat.MessageText
		}

		reader, err := s.tgInstance.File(&m.Sticker.File)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Failed to get sticker image: %v", err))
		} else {
			content := util.RandomString(8)
			err = manbofs.Write(content, reader)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Failed to base64 image photo: %v", err))
			} else {
				msg.MessageType |= chat.MessageImage
				msg.Image = &chat.ImagePayload{
					Content: []string{
						content,
					},
				}
			}
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
