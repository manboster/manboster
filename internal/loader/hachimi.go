package loader

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/hachimi"
)

func LoadHachimiProvider(ctx context.Context, conf config.HachimiConfig, provider hachimi.Provider) (hachimi.Provider, error) {
	newHachimiProvider := provider.New()
	err := newHachimiProvider.Init(ctx, conf.Configuration)
	if err != nil {
		return nil, err
	}
	go func(p hachimi.Provider) {
		err := p.Start(ctx)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Loader] Error starting Hachimi provider: %v", err))
		}
	}(newHachimiProvider)

	return newHachimiProvider, nil
}
