package telegram

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"gopkg.in/telebot.v3"
)

func (s *Service) HandleText(ctx context.Context, c telebot.Context, onMsg func(msg *chat.Message)) error {
	// var msg = chat.BuildMessage(s.New())
	var msg *chat.Message
	msg = msg.Build(&Service{})

	color.Cyan("[Manboster Telegram Provider] Got a text message.")

	// define things all we know.
	msg.MessageID = fmt.Sprintf("%d", c.Message().ID)
	msg.Username = c.Sender().FirstName + " " + c.Sender().LastName
	msg.ChatName = c.Chat().Title // Only Group available
	msg.UserID = fmt.Sprintf("%d", c.Sender().ID)
	msg.ChatID = fmt.Sprintf("%d", c.Chat().ID)
	msg.CreatedAt = time.Now()
	msg.Provider = s.Name()

	// check whether replies available or not
	if c.Message().ReplyTo != nil {
		msg.Reply = &chat.Message{}
		msg.Reply.Username = c.Message().ReplyTo.Sender.FirstName + " " + c.Message().ReplyTo.Sender.LastName
		msg.Reply.MessageID = fmt.Sprintf("%d", c.Message().ReplyTo.ID)
		msg.Reply.UserID = fmt.Sprintf("%d", c.Message().ReplyTo.Sender.ID)
		msg.Reply.ChatID = fmt.Sprintf("%d", c.Message().ReplyTo.Chat.ID)
	}

	// define chat type
	var chatType chat.ChatsType
	switch c.Chat().Type {
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
	if strings.HasPrefix(c.Text(), "/") {
		// process "/xxxxxx xxxx" and "/xxxx@xxxxbot xxxxxxx"
		var command string
		var args []string
		if len(strings.Split(c.Text(), "@"+c.Bot().Me.Username)) > 1 {
			command = strings.ToLower(strings.Split(c.Text(), "@"+c.Bot().Me.Username)[0][1:])
			args = strings.Split(c.Text(), " ")[1:]
		} else {
			command = strings.ToLower(strings.Split(c.Text(), " ")[0][1:])
			if len(strings.Split(c.Text(), " ")) > 1 {
				args = strings.Split(c.Text(), " ")[1:]
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
			Text: c.Text(),
		}
	}

	if (msg.ChatType == chat.ChatsGroup || msg.ChatType == chat.ChatsChannel) && strings.HasPrefix(c.Text(), "@"+c.Bot().Me.Username) {
		msg.MessageType = chat.MessageText
		msg.Text = &chat.TextPayload{
			Text: c.Text()[len("@"+c.Bot().Me.Username)+1:],
		}
	}

	if msg.ChatType == chat.ChatsPersonal || ((msg.ChatType == chat.ChatsGroup || msg.ChatType == chat.ChatsChannel) && (msg.Reply != nil || strings.Contains(c.Text(), "@"+c.Bot().Me.Username))) {
		typingCtx, cancelTyping := context.WithCancel(ctx)
		defer cancelTyping()
		go s.Type(telebot.ChatID(c.Chat().ID), typingCtx)

		// call onMsg on index
		onMsg(msg)
	}
	return nil
}
