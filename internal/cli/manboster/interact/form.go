package interact

import (
	"context"
	"fmt"

	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/spec/cli"
	"github.com/manboster/manboster/spec/config"
)

func EditConfig(ctx context.Context, cliProvider cli.Provider, configProvider config.Provider, conf any) (any, error) {
	return Config(ctx, cliProvider, configProvider, conf)
}

func CreateConfig(ctx context.Context, cliProvider cli.Provider, configProvider config.Provider) (any, error) {
	return Config(ctx, cliProvider, configProvider, nil)
}

func Config(ctx context.Context, cliProvider cli.Provider, configProvider config.Provider, initialValue any) (any, error) {
	form := configProvider.Args().ToCliProvider(cliProvider)
	err := form.Build(initialValue)
	if err != nil {
		return nil, err
	}

	values := form.Collect()
	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "mapstructure",
		WeaklyTypedInput: true,
		Result:           configProvider.GetConfig(),
	})
	err = decoder.Decode(values)
	if err != nil {
		return nil, err
	}

	p, ok := configProvider.(config.ProviderWithSetup)
	if ok {
		err = p.Setup(ctx, cliProvider)
		if err != nil {
			return nil, err
		}
	}

	err = configProvider.Validate()
	if err != nil {
		return nil, err
	}

	return configProvider.GetConfig(), nil
}

func buildConfigStringData[T configurable](ctx context.Context, provider T, conf any) (string, error) {
	cfg := provider.Config()
	if conf == nil {
		return "", fmt.Errorf("no configuration provided")
	}

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "mapstructure",
		WeaklyTypedInput: true,
		Result:           cfg.GetConfig(),
	})

	err := decoder.Decode(conf)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", cfg), nil
}

type configurable interface {
	Config() config.Provider
}
