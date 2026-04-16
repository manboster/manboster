package oai_compat

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/util"
	"github.com/sashabaranov/go-openai"
)

// Config contains what you should enter in application configuration.
type Config struct {
	ApiKey  string            `yaml:"api_key" mapstructure:"api_key" json:"api_key"`    // your apikey
	BaseURL string            `yaml:"base_url" mapstructure:"base_url" json:"base_url"` // this is dynamic when you choose oai_compat systems
	Model   []llm.Model       `yaml:"model" mapstructure:"model" json:"model"`          // your wanted model's information like anthropic/claude-sonnet-4.5
	Headers map[string]string `json:"headers" mapstructure:"headers" yaml:"headers"`
}

const CustomModel = "__CustomModel__"

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

func (c *Config) VerifyAndConvert(ctx context.Context) error {
	svc := NewService(&openai.Client{})
	err := svc.Init(ctx, c)
	models, err := svc.FetchModels(ctx)
	var modelValues []string

	if err != nil {
		err := huh.NewForm(huh.NewGroup(huh.NewNote().Title("Oops! We can't get information based on your credentials!").Description("We strongly recommend you check your Internet Connection and credentials. If you have a poor Internet connection or have a strong belief that you're not wrong, please press enter to proceed to add model manually. Otherwise, please press Ctrl + C to exit.").Next(true))).Run()
		modelValues = append(modelValues, CustomModel)
		if err != nil {
			// don't do that, it will continue to configuration acceptance
			os.Exit(1)
		}
	} else {
		var ModelOptions []huh.Option[string]
		for _, m := range models {
			ModelOptions = append(ModelOptions, huh.NewOption(m, m))
		}
		ModelOptions = append(ModelOptions, huh.NewOption("Other Model", CustomModel))

		err = huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().Title("Models").Description("Please select models available from your given API. If you want to change default values, you can later run `manboster config` to do this.").Options(
					ModelOptions...).Value(&modelValues),
			),
		).Run()
		if err != nil {
			return err
		}
	}

	for _, m := range modelValues {
		if m == CustomModel {
			model, err := InputModel()
			if err != nil {
				return err
			}
			c.Model = append(c.Model, model)
		} else {
			// give it a default value, or make user complete? TODO: First read model library and get information.
			c.Model = append(c.Model, llm.Model{
				Name:            m,
				DisplayName:     m,
				Context:         262144,
				MaxOutputTokens: 8192,
				InputPrice:      0,
				OutputPrice:     0,
				Capabilities: llm.Capabilities{
					Input:  llm.CapabilityText,
					Output: llm.CapabilityText,
				},
			})
		}
	}

	if len(c.Model) == 0 {
		return fmt.Errorf("no models selected")
	}

	return nil
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
	return nil
}
