package openrouter

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/llm/oai_compat"
	"github.com/manboster/manboster/spec/cli"
)

// Setup runs its first run
func (c *Config) Setup(ctx context.Context, p cli.Provider) error {
	var modelOptions []huh.Option[string]
	for _, m := range Models() {
		modelOptions = append(modelOptions, huh.NewOption(m.DisplayName, m.Name))
	}
	modelOptions = append(modelOptions, huh.NewOption("Other Model", "_CustomModel_"))

	for _, m := range c.Model {
		c.inputModelData = append(c.inputModelData, m.Name)
	}

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
