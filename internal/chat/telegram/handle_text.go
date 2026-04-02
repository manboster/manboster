package telegram

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/chat"
	"gopkg.in/telebot.v3"
)

func (s *Service) HandleText(ctx context.Context, c telebot.Context, onMsg func(msg *chat.Message)) error {
	msg := &chat.Message{
		Text:        c.Text(),
		MessageType: chat.MessageText,
		MessageID:   fmt.Sprintf("%d", c.Message().ID),
		Username:    c.Sender().FirstName + " " + c.Sender().LastName,
		UserID:      fmt.Sprintf("%d", c.Sender().ID),
		ChatID:      fmt.Sprintf("%d", c.Chat().ID),
		Provider:    "telegram",
	}

	typingCtx, cancelTyping := context.WithCancel(ctx)
	defer cancelTyping()
	go s.Type(telebot.ChatID(c.Chat().ID), typingCtx)

	// call onMsg on index
	onMsg(msg)
	return nil
}
