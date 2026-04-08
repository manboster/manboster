package openrouter

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/util"
)

// Config contains what you should enter in application configuration.
type Config struct {
	ApiKey string `yaml:"api_key" json:"api_key" mapstructure:"api_key"` // your openrouter system's apikey
	// BaseURL string `yaml:"base_url"` // this is fixed so you don't need to enter it.
	Model llm.Model `yaml:"model" json:"model" mapstructure:"model"` // your wanted model like anthropic/claude-sonnet-4.5
}

// ToHuhGroup enables configuration go ahead.
func (c *Config) ToHuhGroup() []*huh.Group {
	var modelOptions []huh.Option[string]
	for _, m := range Models() {
		modelOptions = append(modelOptions, huh.NewOption(m.DisplayName, m.Name))
	}
	modelOptions = append(modelOptions, huh.NewOption("Other Model", "CustomModel"))

	return []*huh.Group{
		huh.NewGroup(
			huh.NewSelect[string]().Title("OpenRouter Models").Description("Select the model you want to use as Manboster's brain.").Options(
				modelOptions...,
			).Value(&c.Model.Name),
			huh.NewInput().Title("Your OpenRouter API Key").Description("Your OpenRouter API Key.\nIf you don't have one, please open https://openrouter.ai/workspaces/default/keys to create one.").EchoMode(huh.EchoModePassword).Value(&c.ApiKey),
		),
	}
}

// GetConfig returns its own struct.
func (c *Config) GetConfig() any {
	return c
}

// VerifyAndConvert ensures configuration is valid.
func (c *Config) VerifyAndConvert() error {
	// If you choose Custom Model, you should specify it.
	if c.Model.Name == "CustomModel" {
		err := huh.NewForm(huh.NewGroup(huh.NewInput().Title("Your Model Name").Description("Please specify the model name. You can copy it by clicking the clipboard icon on OpenRouter's model page.").Value(&c.Model.Name))).Run()
		if err != nil {
			return err
		}
	}
	return nil
}

// String is used to print sth.
func (c *Config) String() string {
	return fmt.Sprintf("API Key: %s, Model: %s", util.MaskSecret(c.ApiKey), c.Model)
}

func (c *Config) Name() string {
	return "openrouter"
}

func (c *Config) DisplayName() string {
	return "OpenRouter"
}
