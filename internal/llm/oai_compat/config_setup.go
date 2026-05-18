package oai_compat

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config/model"
	"github.com/manboster/manboster/spec/cli"
	"github.com/manboster/manboster/spec/llm"
	"github.com/sashabaranov/go-openai"
)

func (c *Config) Setup(ctx context.Context, p cli.Provider) error {
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
	for _, m := range c.Model {
		modelValues = append(modelValues, m.Name)
	}

	if err != nil {
		err = p.Alert("Oops! We can't get information based on your credentials!", "We strongly recommend you check your Internet Connection and credentials. If you have a poor Internet connection or have a strong belief that you're not wrong, please press enter to proceed to add model manually. Otherwise, please press Ctrl + C to exit.")
		modelValues = append(modelValues, CustomModel)
		if err != nil {
			return err
		}
	} else {
		var ModelOptions []huh.Option[string]
		isFirst := true
		for _, m := range models {
			option := huh.NewOption(m, m)
			if isFirst {
				option.Selected(true)
				isFirst = false
			}
			ModelOptions = append(ModelOptions, option)
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
