package loader

import (
	"context"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
)

func (l *Loader) LoadChat(ctx context.Context, chatConfig config.ChatConfig) error {
	cProvider, err := chat.GetProvider(chatConfig.Provider)
	if err != nil {
		return err
	}

	go l.RunChat(ctx, cProvider, chatConfig.Configuration)
	return nil
}

func (l *Loader) LoadChats(ctx context.Context, chatConfigs []config.ChatConfig) error {
	color.Blue("[Manboster Loader] Loading chat providers...")
	for _, chatConfig := range chatConfigs {
		err := l.LoadChat(ctx, chatConfig)
		if err != nil {
			return err
		}
	}
	return nil
}
