package telegram

import (
	"context"
	"fmt"
	"strconv"

	"github.com/manboster/manboster/core/chat"
	"gopkg.in/telebot.v3"
)

// SendMessage sends a message to user.
func (s *Service) SendMessage(ctx context.Context, msg *chat.Message) error {
	id, err := strconv.ParseInt(msg.ChatID, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid chat id: %w", err)
	}

	// define recipients
	recp := telebot.ChatID(id)

	_, err = s.tgInstance.Send(recp, msg.Text)
	if err != nil {
		return err
	}
	return nil
}
