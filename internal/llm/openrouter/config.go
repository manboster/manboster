package openrouter

import (
	"context"
	"errors"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/llm/oai_compat"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/llm"
)

// Config contains what you should enter in application configuration.
type Config struct {
	ApiKey         string      `yaml:"api_key" json:"api_key" mapstructure:"api_key" manboconfig:"required,secret,name:Your OpenRouter APIKey,desc:Your OpenRouter API Key.\nIf you don't have one please open https://openrouter.ai/workspaces/default/keys to create one."` // your openrouter system's apikey
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

// Setup runs its first run
func (c *Config) Setup(ctx context.Context) error {
	var modelOptions []huh.Option[string]
	for _, m := range Models() {
		modelOptions = append(modelOptions, huh.NewOption(m.DisplayName, m.Name))
	}
	modelOptions = append(modelOptions, huh.NewOption("Other Model", "_CustomModel_"))

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().Title("OpenRouter Models").Description("Select the model you want to use as Manboster's brain.").Options(
				modelOptions...,
			).Value(&c.inputModelData))).Run()
	if err != nil {
		return err
	}

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
