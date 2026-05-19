package interactive

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/tool"
)

func configLandingToolActionForm() (configLandingActionSelection, error) {
	return configLandingActionForm("Add a new Tool provider", "Select an existing tool provider", "Quit")
}

func runLandingToolActionForm(ctx context.Context) error {
	defer helper.ClearScreen()

	conf := config.Read()
	printConfigToolProvidersData(ctx)
	se, err := configLandingToolActionForm()
	if err != nil {
		return err
	}
	switch se {
	case configLandingActionAdd:
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
		var toolProviders []tool.Provider
		var confData any
		for _, p := range conf.Tools {
			pr, err := tool.GetProvider(p.Name)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Client] Failed to get tool provider %s: %q", p, err))
			}
			toolProviders = append(toolProviders, pr)
		}

		provider, err := SelectSingleToolForm(ctx, toolProviders, "Please choose the tool Provider:")
		if err != nil {
			return err
		}

		for _, p := range conf.Tools {
			if p.Name == provider.Name() {
				confData = p.Configuration
			}
		}
		cfg := provider.Config()

		helper.ClearScreen()
		var outputMsg strings.Builder
		outputMsg.WriteString(fmt.Sprintf("`%s`", provider.DisplayName()))
		if confData != nil {
			// get config
			err = mapstructure.Decode(confData, &cfg)
			if err != nil {
				outputMsg.WriteString(fmt.Sprintf(" could not get this!\n"))
			}
			outputMsg.WriteString(fmt.Sprintf(", config: %s", cfg))
		}
		outputMsg.WriteString(fmt.Sprintf("\n"))
		helper.DisplayText(outputMsg.String())

		sel, err := configPageActionSelectForm("Edit this tool provider", "Delete this tool provider", "Quit")
		if err != nil {
			return err
		}

		switch sel {
		case configLandingPageEdit:
			if cfg != nil {
				newConfig, err := RunEditConfig(ctx, cfg, confData)
				if err != nil {
					return err
				}
				for i, p := range conf.Tools {
					if provider.Name() == p.Name {
						conf.Tools[i].Configuration = newConfig
					}
				}
				err = config.Write(conf, config.Path("config.yaml"))
				if err != nil {
					return err
				}
				color.Green("Successfully updated the tool provider.")
			} else {
				color.Yellow("No need to configure because there is no configuration available!")
			}
			time.Sleep(1 * time.Second)
		case configLandingPageDelete:
			for i, p := range conf.Tools {
				if provider.Name() == p.Name {
					conf.Tools = append(conf.Tools[:i], conf.Tools[i+1:]...)
				}
			}
			err := config.Write(conf, config.Path("config.yaml"))
			if err != nil {
				return err
			}
			color.Green("Successfully deleted the tool provider.")
			time.Sleep(1 * time.Second)
		case configLandingPageQuit:
			return nil
		}
	case configLandingActionQuit:
		color.Blue("Bye!")
		return nil
	default:
		return fmt.Errorf("unexpected landing tool-action form: %s", se)
	}
	return nil
}
