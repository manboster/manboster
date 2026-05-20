package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/tool"
	_ "github.com/manboster/manboster/internal/tool/all"
	"github.com/manboster/manboster/spec/cli"
)

type toolConfigAction string

const (
	toolConfigDelete toolConfigAction = _DELETE_
	toolConfigEdit   toolConfigAction = _EDIT_
	toolConfigQuit   toolConfigAction = _QUIT_
)

func (a toolConfigAction) Name() string {
	return string(a)
}

func (a toolConfigAction) DisplayName() string {
	switch a {
	case toolConfigDelete:
		return "Delete this provider"
	case toolConfigEdit:
		return "Edit this provider"
	case toolConfigQuit:
		return "Quit"
	default:
		return ""
	}
}

func runToolConfigs(p cli.Provider, cfg config.Config) ([]config.ToolConfig, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var option cli.Option
	for {
		// reload on every iteration so changes are reflected
		var toolProviders []tool.Provider
		for _, c := range cfg.Tools {
			provider, err := tool.GetProvider(c.Name)
			if err != nil {
				return nil, err
			}
			toolProviders = append(toolProviders, provider)
		}

		occupy := make(map[string]bool)
		for _, c := range cfg.Tools {
			occupy[c.Name] = true
		}
		var availProviders []tool.Provider
		for _, name := range tool.AvailProviders() {
			if !occupy[name] {
				pr, err := tool.GetProvider(name)
				if err != nil {
					continue
				}
				availProviders = append(availProviders, pr)
			}
		}

		options := cli.BuildOptionsWithDescription[tool.Provider](toolProviders, nil)
		options = append(options, quitOption)
		if len(availProviders) > 0 {
			options = append([]cli.Option{addOption}, options...)
		}

		var err error
		option, err = p.Select("Select a tool provider to configure.", "Please select a tool provider to configure.", options, option.Value, func(option cli.Option) error {
			for _, o := range options {
				if o.Value == option.Value {
					return nil
				}
			}
			return fmt.Errorf("unknown tool provider selected: %s", option.Value)
		})
		if err != nil {
			return nil, err
		}

		if option.Value == _QUIT_ {
			return cfg.Tools, nil
		}

		if option.Value == _ADD_ {
			newTools, err := runOnboardToolConfig(p)
			if err != nil {
				return nil, err
			}
			cfg.Tools = append(cfg.Tools, newTools...)
			continue
		}

		var selectedConfig config.ToolConfig
		var selectedProvider tool.Provider
		selectedIndex := -1
		for i, c := range cfg.Tools {
			if c.Name == option.Value {
				selectedConfig = c
				selectedIndex = i
				pr, err := tool.GetProvider(c.Name)
				if err != nil {
					return nil, err
				}
				selectedProvider = pr
				break
			}
		}
		if selectedIndex == -1 {
			return nil, fmt.Errorf("unknown tool provider selected: %s", option.Value)
		}

		se := []toolConfigAction{toolConfigEdit, toolConfigDelete, toolConfigQuit}
		opts := cli.BuildOptions[toolConfigAction](se, nil)
		form := newConfigForm[toolConfigAction]()

		form.Register(toolConfigDelete, func() error {
			confirm, err := p.Prompt(fmt.Sprintf("Do you want to delete %q?\n\nYour action is IRREVERSIBLE!", selectedConfig.Name), "Do you want to continue?", "Yes", "No")
			if err != nil {
				return err
			}
			if !confirm {
				return fmt.Errorf("cancelled")
			}
			cfg.Tools = append(cfg.Tools[:selectedIndex], cfg.Tools[selectedIndex+1:]...)
			if err := p.Alert("Manboster Configuration Wizard", fmt.Sprintf("Tool provider %q deleted successfully!", selectedConfig.Name)); err != nil {
				return err
			}
			return errQuit
		})

		form.Register(toolConfigEdit, func() error {
			providerCfg := selectedProvider.Config()
			if providerCfg == nil {
				return p.Alert("No configuration needed", fmt.Sprintf("%s does not require any configuration.", selectedProvider.DisplayName()))
			}
			conf, err := EditConfig(ctx, p, providerCfg, selectedConfig.Configuration)
			if err != nil {
				return err
			}
			selectedConfig.Configuration = conf
			cfg.Tools[selectedIndex] = selectedConfig
			return errQuit
		})

		form.Register(toolConfigQuit, nilFunc)

		err = handleWithPrompt[toolConfigAction](p, form, opts, fmt.Sprintf("This tool provider %s's info:\n\n%s", selectedProvider.DisplayName(), selectedConfig.Configuration), "What do you want to do with it?")
		if err != nil {
			return nil, err
		}
	}
}
