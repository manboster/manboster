package loader

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/hook"
	"github.com/manboster/manboster/internal/tool"

	_ "github.com/manboster/manboster/internal/tool/all"
)

func LoadToolCallProviders(ctx context.Context, cfg *config.Config) ([]tool.Provider, error) {
	var toolCallProviders []tool.Provider
	for _, conf := range cfg.Tools {
		provider, err := tool.GetProvider(conf.Name)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Loader] We encountered an problem while loading tool call provider %q: %q", conf.Name, err))
			continue
		}
		color.Blue(fmt.Sprintf("[Manboster Loader] Loading tool call provider %q...", provider.DisplayName()))
		err = provider.Init(ctx, conf.Configuration)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Loader] We encountered an problem while loading tool call provider %q: %q", provider.DisplayName(), err))
			continue
		}
		provider.RegisterHook(hook.Reg)
		go func(provider tool.Provider) {
			err := provider.Start(ctx)
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Loader] We encountered an problem while polling tool call provider %q: %q", provider.DisplayName(), err))
			}

			defer func(provider tool.Provider) {
				err := provider.Stop()
				if err != nil {
					color.Red(fmt.Sprintf("[Manboster Loader] We encountered an problem while stopping tool call provider %q: %q", provider.DisplayName(), err))
				}
			}(provider)
		}(provider)
		toolCallProviders = append(toolCallProviders, provider)
	}
	return toolCallProviders, nil
}
