package engine

import (
	"context"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
)

func (e *Engine) loadChats(ctx context.Context) error {
	color.Blue("[Manboster Engine] Loading chat providers...")
	for _, chatConfig := range e.config.Chats {

		cProvider, err := chat.GetProvider(chatConfig.Provider)
		if err != nil {
			return err
		}

		go e.RunChat(ctx, cProvider, chatConfig.Configuration)
	}
	return nil
}
