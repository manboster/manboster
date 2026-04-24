package cli

import (
	"context"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/internal/config"
)

// RunOnboardConfig runs provider and gets config result
func RunOnboardConfig(ctx context.Context, provider config.Provider) (any, error) {
	err := huh.NewForm(provider.ToHuhGroup()...).Run()
	if err != nil {
		return nil, err
	}

	err = provider.VerifyAndConvert(ctx)
	if err != nil {
		return nil, err
	}

	return provider.GetConfig(), nil
}
