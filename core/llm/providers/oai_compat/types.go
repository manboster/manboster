package oai_compat

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/core/util"
)

// Config contains what you should enter in application configuration.
type Config struct {
	ApiKey  string            `yaml:"api_key" mapstructure:"api_key" json:"api_key"`    // your apikey
	BaseURL string            `yaml:"base_url" mapstructure:"base_url" json:"base_url"` // this is dynamic when you choose oai_compat systems
	Model   string            `yaml:"model" mapstructure:"model" json:"model"`          // your wanted model like anthropic/claude-sonnet-4.5
	Headers map[string]string `json:"headers" mapstructure:"headers" yaml:"headers"`
}

// ToHuhGroup enables configuration go ahead.
func (c *Config) ToHuhGroup() []*huh.Group {
	return []*huh.Group{
		huh.NewGroup(
			huh.NewInput().Title("API Site URL").Description("The URL used to call API.\nIf you don't have one, please head to your provider and ask for it.").Value(&c.BaseURL),
			huh.NewInput().Title("API Key").Description("Your API Key.\nIf you don't have one, please go to your provider's API Key manage page and create one.").EchoMode(huh.EchoModePassword).Value(&c.ApiKey),
			huh.NewInput().Title("Model").Description("The model name you want to use as Manboster's brain.").Value(&c.Model),
		),
	}
}

// GetConfig returns itself directly to the app.
func (c *Config) GetConfig() any {
	return c
}

// String is used to print sth.
func (c *Config) String() string {
	return fmt.Sprintf("API URL: %s, API Key: %s, Model: %s", c.BaseURL, util.MaskSecret(c.ApiKey), c.Model)
}
