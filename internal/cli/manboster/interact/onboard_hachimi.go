package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/hachimi"
	_ "github.com/manboster/manboster/internal/hachimi/all"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/cli"
)

// runOnboardHachimiConfigs runs hachimi Config
func runOnboardHachimiConfigs(p cli.Provider) (config.HachimiConfigs, error) {
	conf := config.HachimiConfigs{}
	var hachimiConfigs []config.HachimiConfig

	confirm, err := p.Prompt(i18n.T(keys.OnboardHachimiFeaturePrompt), i18n.T(keys.OnboardHachimiEnableQuestion), i18n.T(keys.BtnYes), i18n.T(keys.BtnNo))
	if err != nil {
		return conf, err
	}

	if !confirm {
		conf.Enabled = false
		return conf, nil
	}

	conf.Enabled = true
	allHachimiProviders := hachimi.AllProviders()
	hachimiProviders := allHachimiProviders

	for {
		hachimiConfig, err := runOnboardHachimiConfig(p, hachimiProviders)
		if err != nil {
			err := p.Alert(i18n.T(keys.WizardTitle), fmt.Sprintf(i18n.T(keys.OnboardHachimiConfigError), err))
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

		ok, err := p.Prompt(fmt.Sprintf(i18n.T(keys.OnboardHachimiAddedCount), len(hachimiConfigs)), "", i18n.T(keys.BtnContinue), i18n.T(keys.BtnExit))
		if err != nil {
			return conf, err
		}
		if len(hachimiProviders) == 0 || !ok {
			if len(hachimiProviders) == 0 && ok {
				err := p.Alert(i18n.T(keys.WizardTitle), i18n.T(keys.OnboardHachimiNoMore))
				if err != nil {
					return conf, err
				}
			}
			conf.Hachimi = hachimiConfigs

			// select default provider
			if len(hachimiConfigs) == 1 {
				conf.Provider = hachimiConfigs[0].Provider
			} else {
				var defaultOptions []cli.Option
				for _, hc := range hachimiConfigs {
					defaultOptions = append(defaultOptions, cli.Option{
						Key:   hc.Provider,
						Value: hc.Provider,
					})
				}
				selected, err := p.Select(i18n.T(keys.OnboardHachimiSelectDefault), i18n.T(keys.OnboardHachimiSelectHelp), defaultOptions, "", func(option cli.Option) error {
					for _, o := range defaultOptions {
						if o.Value == option.Value {
							return nil
						}
					}
					return fmt.Errorf("unknown provider %q", option.Value)
				})
				if err != nil {
					return conf, err
				}
				conf.Provider = selected.Value
			}

			return conf, nil
		}
	}
}

func runOnboardHachimiConfig(p cli.Provider, hachimiProviders []hachimi.Provider) (config.HachimiConfig, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.HachimiConfig{}
	options := util.BuildOptionsForConfig[hachimi.Provider](hachimiProviders, nil)
	hachimiProviderOption, err := p.Select(i18n.T(keys.OnboardHachimiSelectProvider), "", options, "", func(option cli.Option) error {
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
