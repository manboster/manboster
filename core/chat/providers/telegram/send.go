package telegram

import (
	"context"

	"github.com/manboster/manboster/core/chat"
)

// SendMessage sends a message to user.
func (s *Service) SendMessage(ctx context.Context, msg *chat.Message) error {
	recp, err := recipientParser(msg.ChatID)
	if err != nil {
		return err
	}
	_, err = s.tgInstance.Send(recp, msg.Text)
	if err != nil {
		return err
	}
	return nil
}
