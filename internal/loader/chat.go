package loader

import (
	"context"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
	chatType "github.com/manboster/manboster/spec/chat"
)

func (l *Loader) LoadChat(ctx context.Context, chatConfig config.ChatConfig) (chatType.Provider, error) {
	cProvider, err := chat.GetProvider(chatConfig.Provider)
	if err != nil {
		return nil, err
	}

	go l.RunChat(ctx, cProvider, chatConfig.Configuration)
	return cProvider, nil
}

func (l *Loader) LoadChats(ctx context.Context, chatConfigs []config.ChatConfig) ([]chatType.Provider, error) {
	color.Blue("[Manboster Loader] Loading chat providers...")
	var providers []chatType.Provider
	for _, chatConfig := range chatConfigs {
		provider, err := l.LoadChat(ctx, chatConfig)
		if err != nil {
			color.Yellow("[Manboster Loader] Loading chat provider failed: %s", err)
			continue
		}
		providers = append(providers, provider)
	}
	return providers, nil
}
