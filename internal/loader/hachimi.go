package loader

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/hachimi"
	hachimi2 "github.com/manboster/manboster/spec/hachimi"
)

func LoadHachimiProvider(ctx context.Context, conf config.HachimiConfig) (hachimi2.Provider, error) {
	provider, err := hachimi.GetProvider(conf.Provider)
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Loader] Could not load Hachimi Provider: %q", err))
		return nil, err
	}

	newHachimiProvider := provider.New()
	err = newHachimiProvider.Init(ctx, conf.Configuration)
	if err != nil {
		return nil, err
	}
	go func(ctx context.Context, p hachimi2.Provider) {
		err := p.Start(ctx)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Loader] Error starting Hachimi provider: %v", err))
		}

		<-ctx.Done()
		defer func(p hachimi2.Provider) {
			err := p.Stop()
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Loader] Error stopping Hachimi provider: %v", err))
			}
		}(p)
	}(ctx, newHachimiProvider)

	return newHachimiProvider, nil
}
