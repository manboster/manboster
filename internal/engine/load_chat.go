package engine

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
)

func (e *Engine) loadChats(ctx context.Context) error {
	for _, chatConfig := range e.config.Chats {
		cProvider, err := chat.GetProvider(chatConfig.Provider)
		if err != nil {
			return err
		}

		go e.RunChat(ctx, cProvider, chatConfig.Configuration)
	}
	return nil
}
