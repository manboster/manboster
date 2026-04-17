package telegram

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/chat"
	"gopkg.in/telebot.v3"
)

func (s *Service) Notify(ctx context.Context, chatID string, action chat.ActionType) error {
	switch action {
	case chat.ActionTyping:
		recipient, err := recipientParser(chatID)
		if err != nil {
			return err
		}
		return s.tgInstance.Notify(recipient, telebot.ChatAction(action))
	default:
		return fmt.Errorf("invalid action type: %v", action)
	}
}
