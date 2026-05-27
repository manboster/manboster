package oai_compat

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config/model"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
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
		err = p.Alert(i18n.T(keys.OAICompatSetupCredentialError), i18n.T(keys.OAICompatSetupCredentialErrorMsg))
		modelValues = append(modelValues, CustomModel)
		if err != nil {
			return err
		}
	} else {
		options := cli.BuildStringOptions(models, modelValues)
		options = append(options, cli.Option{
			Key:   i18n.T(keys.OAICompatSetupOtherModel),
			Value: CustomModel,
		})

		modelOptions, err := p.MultiSelect(i18n.T(keys.OAICompatSetupModelSelectPrompt), i18n.T(keys.OAICompatSetupModelSelectHelp), options, modelValues, func(options []cli.Option) error {
			for _, option := range options {
				opt := option
				mark := false
				for _, modelOption := range options {
					if opt == modelOption {
						mark = true
						break
					}
				}
				if !mark {
					return fmt.Errorf("option '%s' not found in options", opt.Value)
				}
			}
			return nil
		})
		if err != nil {
			return err
		}

		modelValues = []string{}
		for _, option := range modelOptions {
			modelValues = append(modelValues, option.Value)
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
			modelData := llm.Model{}
			modelData, avail := model.Search(m)
			modelData.Name = m
			if !avail {
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
