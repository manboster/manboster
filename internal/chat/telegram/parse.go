package telegram

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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

	// commands, help to process commands which prefixes started with "/".
	if strings.HasPrefix(m.Text, "/") {
		// process "/xxxxxx xxxx" and "/xxxx@xxxxbot xxxxxxx"
		var command string
		var args []string
		// if it is bot's username
		if len(strings.Split(m.Text, "@"+s.tgInstance.Me.Username)) > 1 {
			command = strings.ToLower(strings.Split(m.Text, "@"+s.tgInstance.Me.Username)[0][1:])
			args = strings.Split(m.Text, " ")[1:]
		} else {
			command = strings.ToLower(strings.Split(m.Text, " ")[0][1:])
			if len(strings.Split(m.Text, " ")) > 1 {
				args = strings.Split(m.Text, " ")[1:]
			}
		}
		msg.MessageType = chat.MessageCommand
		msg.Command = &chat.CommandPayload{
			CommandType: chat.CommandType(command),
			CommandArgs: args,
		}
	} else {
		msg.MessageType = chat.MessageText
		msg.Text = &chat.TextPayload{
			Text: m.Text,
		}
		// fmt.Printf(m.Text)
		// handle reply to quote message
		if m.Quote != nil {
			msg.Text = &chat.TextPayload{
				Text: m.Quote.Text,
			}
		}
	}
	return nil
}

func (s *Service) msgBaseParser(msg *chat.Message, c telebot.Context) {
	// define things all we know.
	msg.MessageID = fmt.Sprintf("%d", c.Message().ID)
	msg.Username = c.Sender().FirstName + " " + c.Sender().LastName
	msg.ChatName = c.Chat().Title // Only Group available
	msg.UserID = fmt.Sprintf("%d", c.Sender().ID)
	msg.ChatID = fmt.Sprintf("%d", c.Chat().ID)
	msg.CreatedAt = time.Now()
	msg.Provider = s.Name()
}
