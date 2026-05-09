package loader

import (
	"context"
	"fmt"

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

	err = cProvider.Init(ctx, chatConfig.Configuration)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] Failed to init a chat provider on %s, get error: %q", cProvider.DisplayName(), err))
	}
	go l.RunChat(ctx, cProvider)
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
