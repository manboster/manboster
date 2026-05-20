package openrouter

import (
	"context"
	"fmt"

	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/llm/oai_compat"
	"github.com/manboster/manboster/spec/cli"
	"github.com/manboster/manboster/spec/config"
)

// Setup runs its first run
func (c *Config) Setup(ctx context.Context, p cli.Provider) error {
	s := oai_compat.Service{}
	err := s.InitWithConfig(ctx, &oai_compat.Config{
		ProviderName:        c.Name(),
		ProviderDisplayName: c.DisplayName(),
		BaseURL:             openrouterBaseurl,
		ApiKey:              c.ApiKey,
		Model:               nil,
	})
	if err != nil {
		return err
	}

	sp, ok := s.Config().(config.ProviderWithSetup)
	if !ok {
		return fmt.Errorf("config is not a setup-able provider")
	}

	err = sp.Setup(ctx, p)
	if err != nil {
		return err
	}

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "mapstructure",
		WeaklyTypedInput: true,
		Result:           &c.Config,
	})

	c.ApiKey = "PLEASE_EDIT_API_KEY_IN_CONFIG_FIELD_BELOW"
	return decoder.Decode(sp.GetConfig())
}
