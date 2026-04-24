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
			return err
		}
	}

	// check whether message forward available or not
	if c.Message().IsForwarded() {
		msg.Forward = &chat.Message{}
		if c.Message().Origin != nil {
			msg.Forward.Username = c.Message().Origin.SenderUsername
			if c.Message().Origin.Sender != nil {
				msg.Forward.Username = c.Message().Origin.Sender.FirstName + " " + c.Message().Origin.Sender.LastName
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
	if c.Message().Origin != nil && c.Message().Origin.SenderUsername != "" {
		if msg.Forward == nil {
			msg.Forward = &chat.Message{}
		}
		msg.Forward.Username = c.Message().Origin.SenderUsername
	}

	// j, _ := json.MarshalIndent(c.Message(), "", " ")
	// fmt.Println(string(j))

	err := s.msgParser(msg, c.Message())
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Failed to parse message data: %q", err.Error()))
		return err
	}

	// TODO: Passthrough all messages from Group, handle it in handleMessage, check.
	if ((msg.ChatType == chat.ChatsGroup || msg.ChatType == chat.ChatsChannel) && strings.HasPrefix(c.Text(), "@"+c.Bot().Me.Username)) && msg.MessageType != chat.MessageCommand {
		msg.MessageType = chat.MessageText
		msg.Text = &chat.TextPayload{
			Text: c.Text()[len("@"+c.Bot().Me.Username)+1:],
		}
	}

	if msg.ChatType == chat.ChatsPersonal || ((msg.ChatType == chat.ChatsGroup || msg.ChatType == chat.ChatsChannel) && ((msg.Reply != nil && c.Message().ReplyTo.Sender.ID == c.Bot().Me.ID) || strings.Contains(c.Text(), "@"+c.Bot().Me.Username))) || msg.MessageType == chat.MessageCommand {
		typingCtx, cancelTyping := context.WithCancel(ctx)
		defer cancelTyping()
		go s.Type(telebot.ChatID(c.Chat().ID), typingCtx)

		// call onMsg on index
		onMsg(msg)
	}
	return nil
}
