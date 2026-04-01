package telegram

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/core/chat"
	"gopkg.in/telebot.v3"
)

// SendMessage sends a message to user.
func (s *Service) SendMessage(ctx context.Context, msg *chat.Message) error {
	recp, err := recipientParser(msg.ChatID)
	if err != nil {
		return err
	}

	opts := &telebot.SendOptions{}

	if msg.MessageID != "" {
		var replyID int
		_, err := fmt.Sscanf(msg.MessageID, "%d", &replyID)
		opts.ReplyTo = &telebot.Message{ID: replyID}
		if err != nil {
			return err
		}
	}
	_, err = s.tgInstance.Send(recp, msg.Text, opts)
	if err != nil {
		return err
	}
	return nil
}
