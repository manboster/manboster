package openrouter

import (
	"context"
	"errors"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/llm/oai_compat"
	"github.com/manboster/manboster/internal/util"
)

// Config contains what you should enter in application configuration.
type Config struct {
	ApiKey string `yaml:"api_key" json:"api_key" mapstructure:"api_key"` // your openrouter system's apikey
	// BaseURL string `yaml:"base_url"` // this is fixed so you don't need to enter it.
	Model          []llm.Model `yaml:"model" json:"model" mapstructure:"model"` // your wanted model like anthropic/claude-sonnet-4.5
	inputModelData []string    // internal input keys
}

// ToHuhGroup enables configuration go ahead.
func (c *Config) ToHuhGroup() []*huh.Group {
	var modelOptions []huh.Option[string]
	for _, m := range Models() {
		modelOptions = append(modelOptions, huh.NewOption(m.DisplayName, m.Name))
	}
	modelOptions = append(modelOptions, huh.NewOption("Other Model", "_CustomModel_"))

	return []*huh.Group{
		huh.NewGroup(
			huh.NewMultiSelect[string]().Title("OpenRouter Models").Description("Select the model you want to use as Manboster's brain.").Options(
				modelOptions...,
			).Value(&c.inputModelData),
			huh.NewInput().Title("Your OpenRouter API Key").Description("Your OpenRouter API Key.\nIf you don't have one, please open https://openrouter.ai/workspaces/default/keys to create one.").EchoMode(huh.EchoModePassword).Value(&c.ApiKey),
		),
	}
}

// GetConfig returns its own struct.
func (c *Config) GetConfig() any {
	return c
}

// VerifyAndConvert ensures configuration is valid.
func (c *Config) VerifyAndConvert(ctx context.Context) error {
	if len(c.inputModelData) == 0 {
		return ErrModelNameRequired
	}

	// If you choose Custom Model, you should specify it.
	for _, m := range c.inputModelData {
		if m == oai_compat.CustomModel {
			customModel, err := c.InputCustomModel()
			if err != nil {
				return err
			}
			c.Model = append(c.Model, customModel)
		} else {
			avail := false
			for _, k := range Models() {
				// check if these name is valid or not
				if k.Name == m {
					c.Model = append(c.Model, k)
					avail = true
				}
			}
			if !avail {
				color.Yellow(fmt.Sprintf("Input Model %s is not found in models data", m))
			}
		}
	}

	return nil
}

// String is used to print sth.
func (c *Config) String() string {
	return fmt.Sprintf("API Key: %s, Model: %+v", util.MaskSecret(c.ApiKey), c.Model)
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
