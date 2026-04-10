package oai_compat

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/util"
)

// Config contains what you should enter in application configuration.
type Config struct {
	ApiKey  string            `yaml:"api_key" mapstructure:"api_key" json:"api_key"`    // your apikey
	BaseURL string            `yaml:"base_url" mapstructure:"base_url" json:"base_url"` // this is dynamic when you choose oai_compat systems
	Model   []llm.Model       `yaml:"model" mapstructure:"model" json:"model"`          // your wanted model's information like anthropic/claude-sonnet-4.5
	Headers map[string]string `json:"headers" mapstructure:"headers" yaml:"headers"`
}

// ToHuhGroup enables configuration go ahead.
func (c *Config) ToHuhGroup() []*huh.Group {
	return []*huh.Group{
		huh.NewGroup(
			huh.NewInput().Title("API Site URL").Description("The URL used to call API.\nIf you don't have one, please head to your provider and ask for it.").Value(&c.BaseURL),
			huh.NewInput().Title("API Key").Description("Your API Key.\nIf you don't have one, please go to your provider's API Key manage page and create one.").EchoMode(huh.EchoModePassword).Value(&c.ApiKey),
		),
	}
}

// GetConfig returns itself directly to the app.
func (c *Config) GetConfig() any {
	return c
}

// String is used to print sth.
func (c *Config) String() string {
	return fmt.Sprintf("API URL: %s, API Key: %s, Model: %+v", c.BaseURL, util.MaskSecret(c.ApiKey), c.Model)
}

func (c *Config) Name() string {
	return "oai-compat"
}

func (c *Config) DisplayName() string {
	return "OpenAI compatible API"
}

func (c *Config) VerifyAndConvert() error {

	model, err := InputModel()
	if err != nil {
		return err
	}
	c.Model = append(c.Model, model)
	return nil
}
