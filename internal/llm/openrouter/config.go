package openrouter

import (
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/llm/oai_compat"
	"github.com/manboster/manboster/spec/config"
)

// Config contains what you should enter in application configuration.
type Config struct {
	ApiKey             string `yaml:"api_key" json:"api_key" mapstructure:"api_key" manboconfig:"required;secret;id:llm.openrouter.api_key" validation:"^[a-zA-Z0-9_-]+$"` // your openrouter system's apikey
	*oai_compat.Config `mapstructure:"config"`
}

const openrouterBaseurl = "https://openrouter.ai/api/v1" // fixed openrouter api calls

// Args returns args from struct Config
func (c *Config) Args() *config.Args {
	return config.ArgsFromStruct(Config{})
}

// GetConfig returns its own struct.
func (c *Config) GetConfig() any {
	return c
}

// String is used to print sth.
func (c *Config) String() string {
	return c.Config.String()
}

func (c *Config) Name() string {
	return "openrouter"
}

func (c *Config) DisplayName() string {
	return i18n.T(keys.LLMOpenRouterProvider)
}

func (c *Config) Validate() error {
	return c.Config.Validate()
}
