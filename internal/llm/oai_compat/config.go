package oai_compat

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/llm"
)

var re = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

// Config contains what you should enter in application configuration.
type Config struct {
	ProviderName        string            `yaml:"name" mapstructure:"name" json:"name" manboconfig:"required;name:Your Provider's name;desc:The name of your provider.\nYou can enter what you want but no spaces in your string."`
	ProviderDisplayName string            `yaml:"display_name" mapstructure:"display_name" json:"display_name" manboconfig:"required;name:Your Provider's Display name;desc:The name that will display on your application if you don't know what's this, please leave it empty."`
	BaseURL             string            `yaml:"base_url" mapstructure:"base_url" json:"base_url" manboconfig:"required;name:API BaseURL;desc:The URL used to call API.\nIf you don't have one please head to your provider and ask for it."`          // this is dynamic when you choose oai_compat systems
	ApiKey              string            `yaml:"api_key" mapstructure:"api_key" json:"api_key" manboconfig:"required;secret;name:API Key;desc:Your API Key.\nIf you don't have one, please go to your provider's API Key manage page and create one."` // your apikey
	Model               []llm.Model       `yaml:"model" mapstructure:"model" json:"model" manboconfig:"skip"`                                                                                                                                           // your wanted model's information like anthropic/claude-sonnet-4.5
	Headers             map[string]string `json:"headers" mapstructure:"headers" yaml:"headers" manboconfig:"skip"`
}

const CustomModel = "__CustomModel__"

// Args returns args which configuration go ahead.
func (c *Config) Args() *config.Args {
	return config.ArgsFromStruct(Config{})
}

// GetConfig returns itself directly to the app.
func (c *Config) GetConfig() any {
	return c
}

// String is used to print sth.
func (c *Config) String() string {
	var b strings.Builder
	for _, m := range c.Model {
		b.WriteString(m.DisplayName + ",")
	}
	return fmt.Sprintf("Name: `%s`, API URL: `%s`, API Key: `%s`\nmodels(%d): `%s`", c.ProviderDisplayName, c.BaseURL, util.MaskSecret(c.ApiKey), len(c.Model), b.String())
}

func (c *Config) Name() string {
	return "oai-compat"
}

func (c *Config) DisplayName() string {
	return "OpenAI compatible API"
}

func (c *Config) Validate() error {
	if len(c.Model) == 0 {
		return fmt.Errorf("no models selected")
	}
	if c.BaseURL == "" {
		return fmt.Errorf("no Base URL provided")
	}
	if c.ApiKey == "" {
		return fmt.Errorf("no API key provided")
	}
	if !re.MatchString(c.ProviderName) {
		return fmt.Errorf("invalid provider name")
	}
	return nil
}
