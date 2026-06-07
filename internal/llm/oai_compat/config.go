package oai_compat

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/llm"
)

var re = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

// Config contains what you should enter in application configuration.
type Config struct {
	ProviderName        string            `yaml:"name" mapstructure:"name" json:"name" manboconfig:"required;id:llm.oai_compat.name" validation:"^[a-zA-Z0-9_-]+$"`
	ProviderDisplayName string            `yaml:"display_name" mapstructure:"display_name" json:"display_name" manboconfig:"id:llm.oai_compat.display_name"`
	BaseURL             string            `yaml:"base_url" mapstructure:"base_url" json:"base_url" manboconfig:"required;id:llm.oai_compat.base_url" validation:"^https?://"`          // this is dynamic when you choose oai_compat systems
	ApiKey              string            `yaml:"api_key" mapstructure:"api_key" json:"api_key" manboconfig:"required;secret;id:llm.oai_compat.api_key" validation:"^[a-zA-Z0-9_-]+$"` // your apikey
	Model               []llm.Model       `yaml:"model" mapstructure:"model" json:"model" manboconfig:"skip"`                                                                          // your wanted model's information like anthropic/claude-sonnet-4.5
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
	return i18n.T(keys.LLMOAICompatProvider)
}

func (c *Config) Validate() error {
	if len(c.Model) == 0 {
		return fmt.Errorf("no models selected")
	}
	duplicateMap := map[string]bool{}
	for _, m := range c.Model {
		if a, ok := duplicateMap[m.Name]; ok && a {
			return fmt.Errorf("duplicate model name: %s", m.Name)
		}
		duplicateMap[m.Name] = true
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
