package cli

import (
	"context"

	"github.com/charmbracelet/huh"
	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/spec/config"
)

// RunOnboardConfig runs provider and gets config result
func RunOnboardConfig(ctx context.Context, provider config.Provider) (any, error) {
	form := provider.Args().ToHuhGroup()
	err := huh.NewForm(form.Groups...).Run()
	if err != nil {
		return nil, err
	}

	values := form.Collect()
	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "mapstructure",
		WeaklyTypedInput: true,
		Result:           provider.GetConfig(),
	})
	err = decoder.Decode(values)
	if err != nil {
		return nil, err
	}

	p, ok := provider.(config.ProviderWithSetup)
	if ok {
		err = p.Setup(ctx)
		if err != nil {
			return nil, err
		}
	}

	err = provider.Validate()
	if err != nil {
		return nil, err
	}

	return provider.GetConfig(), nil
}
