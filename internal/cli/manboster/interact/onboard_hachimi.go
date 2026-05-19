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

// runOnboardHachimiConfigs runs hachimi Config
func runOnboardHachimiConfigs(p cli.Provider) (config.HachimiConfigs, error) {
	conf := config.HachimiConfigs{}
	var hachimiConfigs []config.HachimiConfig

	confirm, err := p.Prompt(`
Hachimi is a small language model running on your device side and it can check out LLM's behaviour and evaluate its action.
It wouldn't active in order to save memory until you called it to handle requests in moment.

If your device's available memory is lower than 1GB or you don't know what's this, please disable it.

If you want to activate hachimi feature, please ensure your device have a valid Internet connection and 2GB free disk spaces.
`, "Do you want to activate Hachimi feature?", "Yes", "No")
	if err != nil {
		return conf, err
	}

	if !confirm {
		conf.Enabled = false
		return conf, nil
	}

	allHachimiProviders := hachimi.AllProviders()
	hachimiProviders := allHachimiProviders

	for {
		hachimiConfig, err := runOnboardHachimiConfig(p, hachimiProviders)
		if err != nil {
			err := p.Alert("Manboster Configuration Wizard", fmt.Sprintf("Failed to config %q", err))
			if err != nil {
				return conf, err
			}
			continue
		}

		hachimiConfigs = append(hachimiConfigs, hachimiConfig)
		for i, provider := range hachimiProviders {
			if provider.Name() == hachimiConfig.Provider {
				hachimiProviders = append(hachimiProviders[:i], hachimiProviders[i+1:]...)
				break
			}
		}

		ok, err := p.Prompt(fmt.Sprintf("You've added %d hachimi providers! Do you want to continue?", len(hachimiConfigs)), "", "Continue", "Exit and go on")
		if err != nil {
			return conf, err
		}
		if len(hachimiProviders) == 0 || !ok {
			if len(hachimiProviders) == 0 {
				err := p.Alert("Manboster Configuration Wizard", "There are no more providers available for you to config...")
				if err != nil {
					return conf, err
				}
			}
			break
		}
	}

	return conf, nil
}

func runOnboardHachimiConfig(p cli.Provider, hachimiProviders []hachimi.Provider) (config.HachimiConfig, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.HachimiConfig{}
	options := util.BuildOptionsForConfig[hachimi.Provider](hachimiProviders, nil)
	hachimiProviderOption, err := p.Select("Please select your hachimi's provider", "", options, "", func(option cli.Option) error {
		for _, provider := range hachimiProviders {
			if provider.Name() == option.Value {
				return nil
			}
		}
		return fmt.Errorf("no hachimi provider named %q", option.Value)
	})
	if err != nil {
		return conf, err
	}

	for _, provider := range hachimiProviders {
		if provider.Name() == hachimiProviderOption.Value {
			cfg, err := CreateConfig(ctx, p, provider.Config())
			if err != nil {
				return conf, err
			}
			conf.Provider = provider.Name()
			conf.Configuration = cfg
			return conf, nil
		}
	}

	return conf, fmt.Errorf("no hachimi provider named %q", hachimiProviderOption.Value)
}
