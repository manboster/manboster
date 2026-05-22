package loader

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/hook"
	"github.com/manboster/manboster/internal/tool"
	_ "github.com/manboster/manboster/internal/tool/all"
)

func LoadToolCallProvider(ctx context.Context, provider tool.Provider, conf config.ToolConfig) (tool.Provider, error) {
	err := provider.Init(ctx, conf.Configuration)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] We encountered an problem while loading tool call provider %q: %q", provider.DisplayName(), err))
		return nil, err
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
	return provider, nil
}

func LoadToolCallProviders(ctx context.Context, cfg *config.Config) ([]tool.Provider, error) {
	var toolCallProviders []tool.Provider
	allToolCallProviders := tool.AvailProviders()
	for _, conf := range cfg.Tools {
		provider, err := tool.GetProvider(conf.Name)
		if err != nil {
			// find namespaces and then run
			mark := false
			for _, p := range allToolCallProviders {
				if strings.HasPrefix(p, conf.Name) {
					mark = true
					provider, err := tool.GetProvider(p)
					if err != nil {
						color.Red(fmt.Sprintf("[Manboster Loader] We encountered an problem while loading tool call provider %q: %q", p, err))
					}
					color.Blue(fmt.Sprintf("[Manboster Loader] Loading tool call provider %q...", provider.DisplayName()))
					p, err := LoadToolCallProvider(ctx, provider, conf)
					if err != nil {
						color.Red(fmt.Sprintf("[Manboster Loader] failed to load tool call %s: %q", conf.Name, err))
						continue
					}
					toolCallProviders = append(toolCallProviders, p)
				}
			}
			if !mark {
				color.Red(fmt.Sprintf("[Manboster Loader] We encountered an problem while loading tool call provider %q: %q", conf.Name, err))
			}
			continue
		}
		color.Blue(fmt.Sprintf("[Manboster Loader] Loading tool call provider %q...", provider.DisplayName()))
		p, err := LoadToolCallProvider(ctx, provider, conf)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Loader] failed to load tool call %s: %q", conf.Name, err))
			continue
		}
		toolCallProviders = append(toolCallProviders, p)
	}
	return toolCallProviders, nil
}
