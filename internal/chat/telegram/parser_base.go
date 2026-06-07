package telegram

import (
	"errors"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
	"gopkg.in/telebot.v3"
)

func (s *Service) msgBaseParser(msg *chat.Message, c telebot.Context, onMsg func(msg *chat.Message)) error {
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
		err := s.msgParser(msg.Reply, c.Message().ReplyTo, onMsg)
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
	err := s.msgParser(msg, c.Message(), onMsg)
	if errors.Is(err, ErrImageNoNeedToTrigger) {
		return err
	}

	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Telegram Provider] Failed to parse message data: %q", err.Error()))
	}
	return err
}
