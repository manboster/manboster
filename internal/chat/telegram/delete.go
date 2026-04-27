package telegram

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

// DeleteMessage TODO:
func (s *Service) DeleteMessage(ctx context.Context, msg *chat.Message) error {
	if msg.MessageType&chat.MessageUnknown == 0 {
		return ErrInvalidMessageType
	}
	return nil
}
