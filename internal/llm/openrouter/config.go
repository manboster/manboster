package openrouter

import (
	"errors"
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/llm"
)

// Config contains what you should enter in application configuration.
type Config struct {
	ApiKey         string      `yaml:"api_key" json:"api_key" mapstructure:"api_key" manboconfig:"required;secret;name:Your OpenRouter APIKey;desc:Your OpenRouter API Key.\nIf you don't have one,please open https://openrouter.ai/workspaces/default/keys to create one."` // your openrouter system's apikey
	Model          []llm.Model `yaml:"model" json:"model" mapstructure:"model" manboconfig:"skip"`                                                                                                                                                                            // your wanted model like anthropic/claude-sonnet-4.5
	ID             int         `yaml:"id" json:"id" mapstructure:"id" manboconfig:"skip"`                                                                                                                                                                                     // if duplicate, what id it is?
	inputModelData []string    // internal input keys
}

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
	var b strings.Builder
	for _, m := range c.Model {
		b.WriteString(m.DisplayName + ",")
	}
	return fmt.Sprintf("API Key: `%s`, Models(%d): `%s`", util.MaskSecret(c.ApiKey), len(c.Model), b.String())

}

func (c *Config) Name() string {
	return "openrouter"
}

func (c *Config) DisplayName() string {
	return "OpenRouter"
}

func (c *Config) Validate() error {
	if c.Model == nil {
		return errors.New("model is empty")
	}
	if c.ApiKey == "" {
		return errors.New("api_key is empty")
	}
	return nil
}
