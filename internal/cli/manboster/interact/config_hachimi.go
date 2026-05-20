package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/hachimi"
	_ "github.com/manboster/manboster/internal/hachimi/all"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/cli"
)

type hachimiConfigAction string

const (
	hachimiConfigDelete hachimiConfigAction = _DELETE_
	hachimiConfigEdit   hachimiConfigAction = _EDIT_
	hachimiConfigQuit   hachimiConfigAction = _QUIT_
)

func (a hachimiConfigAction) Name() string {
	return string(a)
}

func (a hachimiConfigAction) DisplayName() string {
	switch a {
	case hachimiConfigDelete:
		return "Delete this provider"
	case hachimiConfigEdit:
		return "Edit this provider"
	case hachimiConfigQuit:
		return "Quit"
	default:
		return ""
	}
}

func runHachimiConfigs(p cli.Provider, cfg config.Config) (config.HachimiConfigs, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var hachimiProviders []hachimi.Provider
	for _, c := range cfg.Hachimi.Hachimi {
		provider, err := hachimi.GetProvider(c.Provider)
		if err != nil {
			return cfg.Hachimi, err
		}
		hachimiProviders = append(hachimiProviders, provider)
	}

	options := util.BuildOptionsForConfig[hachimi.Provider](hachimiProviders, nil)
	options = append(options, quitOption)

	allHachimiProviders := hachimi.AllProviders()
	for i, provider := range hachimiProviders {
		for _, hp := range allHachimiProviders {
			if hp.Name() == provider.Name() {
				allHachimiProviders = append(allHachimiProviders[:i], allHachimiProviders[i+1:]...)
				break
			}
		}
	}
	if len(allHachimiProviders) > 0 {
		options = append([]cli.Option{addOption}, options...)
	}

	var option cli.Option
	for {
		var err error
		option, err = p.Select("Select a Hachimi provider to configure.", "Please select a Hachimi provider to configure.", options, option.Value, func(option cli.Option) error {
			for _, o := range options {
				if o.Value == option.Value {
					return nil
				}
			}
			return fmt.Errorf("unknown Hachimi provider selected: %s", option.Value)
		})
		if err != nil {
			return cfg.Hachimi, err
		}

		if option.Value == _QUIT_ {
			return cfg.Hachimi, nil
		}

		if option.Value == _ADD_ {
			hachimiConfig, err := runOnboardHachimiConfig(p, allHachimiProviders)
			if err != nil {
				return cfg.Hachimi, err
			}
			cfg.Hachimi.Hachimi = append(cfg.Hachimi.Hachimi, hachimiConfig)
			continue
		}

		var selectedConfig config.HachimiConfig
		var selectedProvider hachimi.Provider
		selectedIndex := -1
		for i, c := range cfg.Hachimi.Hachimi {
			if c.Provider == option.Value {
				selectedConfig = c
				selectedIndex = i
				pr, err := hachimi.GetProvider(c.Provider)
				if err != nil {
					return cfg.Hachimi, err
				}
				selectedProvider = pr
				break
			}
		}
		if selectedIndex == -1 {
			return cfg.Hachimi, fmt.Errorf("unknown Hachimi provider selected: %s", option.Value)
		}

		se := []hachimiConfigAction{hachimiConfigEdit, hachimiConfigDelete, hachimiConfigQuit}
		opts := cli.BuildOptions[hachimiConfigAction](se, nil)
		form := newConfigForm[hachimiConfigAction]()

		form.Register(hachimiConfigDelete, func() error {
			confirm, err := p.Prompt(fmt.Sprintf("Do you want to delete %q?\n\nYour action is IRREVERSIBLE!", selectedConfig.Provider), "Do you want to continue?", "Yes", "No")
			if err != nil {
				return err
			}
			if !confirm {
				return fmt.Errorf("cancelled")
			}
			cfg.Hachimi.Hachimi = append(cfg.Hachimi.Hachimi[:selectedIndex], cfg.Hachimi.Hachimi[selectedIndex+1:]...)
			return errQuit
		})

		form.Register(hachimiConfigEdit, func() error {
			conf, err := EditConfig(ctx, p, selectedProvider.Config(), selectedConfig.Configuration)
			if err != nil {
				return err
			}
			selectedConfig.Configuration = conf
			cfg.Hachimi.Hachimi[selectedIndex] = selectedConfig
			return errQuit
		})

		form.Register(hachimiConfigQuit, nilFunc)

		err = handleWithPrompt[hachimiConfigAction](p, form, opts, fmt.Sprintf("This Hachimi provider %s's info:\n\n%s", selectedProvider.DisplayName(), selectedConfig.Configuration), "What do you want to do with it?")
		if err != nil {
			return cfg.Hachimi, err
		}
	}
}
