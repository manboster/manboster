package interactive

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/tool"
)

func configLandingToolActionForm() (configLandingActionSelection, error) {
	return configLandingActionForm("Add a new Tool provider", "Select an existing tool provider", "Quit")
}

func runLandingToolActionForm(ctx context.Context) error {
	defer helper.ClearScreen()

	printConfigToolProvidersData(ctx)
	se, err := configLandingToolActionForm()
	if err != nil {
		return err
	}
	switch se {
	case configLandingActionAdd:
		conf := config.Read()
		var toolProviders []tool.Provider

		occupy := make(map[string]bool)
		for _, c := range conf.Tools {
			occupy[c.Name] = true
		}

		helper.ClearScreen()
		for _, p := range tool.AvailProviders() {
			if !occupy[p] {
				pr, err := tool.GetProvider(p)
				if err != nil {
					color.Yellow(fmt.Sprintf("[Manboster Client] Failed to get tool provider %s: %q", p, err))
				}
				toolProviders = append(toolProviders, pr)
			}
		}

		if len(toolProviders) == 0 {
			color.Yellow("[Manboster Client] No new tool providers available to add!")
			time.Sleep(1 * time.Second)
			return nil
		}

		toolProvider, err := SelectSingleToolForm(ctx, toolProviders, "Please select the Tool provider you want to add:")
		if err != nil {
			return err
		}

		if toolProvider.Config() != nil {
			cf, err := RunOnboardConfig(ctx, toolProvider.Config())
			if err != nil {
				return err
			}
			conf.Tools = append(conf.Tools, config.ToolConfig{
				Name:          toolProvider.Name(),
				Configuration: cf,
			})
		} else {
			conf.Tools = append(conf.Tools, config.ToolConfig{
				Name: toolProvider.Name(),
			})
		}
	case configLandingActionSelect:
	case configLandingActionQuit:
		color.Blue("Bye!")
		return nil
	default:
		return fmt.Errorf("unexpected landing tool-action form: %s", se)
	}
	return nil
}
