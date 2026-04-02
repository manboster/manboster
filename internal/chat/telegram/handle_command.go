package telegram

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
	"gopkg.in/telebot.v3"
)

func (s *Service) HandleCommand(ctx context.Context, c telebot.Context, onMsg func(msg *chat.Message)) error {
	return nil
}
