package oai_compat

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config/model"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/llm"
	"github.com/sashabaranov/go-openai"
)

// Config contains what you should enter in application configuration.
type Config struct {
	ProviderName        string            `yaml:"name" mapstructure:"name" json:"name" manboconfig:"required;name:Your Provider's name;desc:The name of your provider.\nYou can enter what you want but no spaces in your string."`
	ProviderDisplayName string            `yaml:"display_name" mapstructure:"display_name" json:"display_name" manboconfig:"required;name:Your Provider's Display Name;desc:The name that will display on your application if you don't know what's this, please leave it empty."`
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

func (c *Config) Setup(ctx context.Context) error {
	if c.ProviderDisplayName == "" {
		c.ProviderDisplayName = c.ProviderName
	}
	if c.ProviderName == "" {
		c.ProviderName = c.Name()
		c.ProviderDisplayName = c.DisplayName()
	}

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
			iModel, err := InputModel()
			if err != nil {
				return err
			}
			c.Model = append(c.Model, iModel)
		} else {
			// give it a default value, or make user complete?
			modelData := llm.Model{}
			modelData, avail := model.Search(m)
			modelData.Name = m // if this is not 'm', it would not get any available messages
			if !avail {
				// default value
				color.Yellow(fmt.Sprintf("[Manboster Configuration Wizard] We can't find %s in our library, we have set this model's params data to default value. If you want to change it, please edit config file.", m))
				modelData = model.Default(m)
			}
			c.Model = append(c.Model, modelData)
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
