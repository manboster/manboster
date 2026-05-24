package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/spec/cli"
	llmType "github.com/manboster/manboster/spec/llm"
)

// runOnboardAPPConfig runs app Config
func runOnboardAPPConfig(p cli.Provider, cfg config.Config) (config.AppConfig, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.AppConfig{}
	var activatedLLMProviders []llmType.Provider
	for _, l := range cfg.LLMs {
		provider, err := llm.GetProvider(l.Provider)
		if err != nil {
			return conf, err
		}

		err = provider.Init(ctx, l.Configuration)
		if err != nil {
			return conf, err
		}

		activatedLLMProviders = append(activatedLLMProviders, provider)
	}

	options := cli.BuildOptions[llmType.Provider](activatedLLMProviders, nil)
	selected, err := p.Select(i18n.T(keys.OnboardAppSelectProvider), i18n.T(keys.OnboardAppSelectHelp), options, cfg.App.DefaultLLMProvider, func(option cli.Option) error {
		for _, provider := range activatedLLMProviders {
			if provider.Name() == option.Value {
				return nil
			}
		}
		return fmt.Errorf("unknown provider %s", option.Value)
	})
	if err != nil {
		return conf, err
	}

	for _, provider := range activatedLLMProviders {
		if provider.Name() == selected.Value {
			conf.DefaultLLMProvider = selected.Value
			modelOptions := cli.BuildModelOptions[llmType.Model](provider.Models(), nil)
			selectedModel, err := p.Select(i18n.T(keys.OnboardAppSelectModel), i18n.T(keys.OnboardAppSelectHelp), modelOptions, cfg.App.DefaultLLMModel, func(option cli.Option) error {
				for _, model := range modelOptions {
					if model.Value == option.Value {
						return nil
					}
				}
				return fmt.Errorf("unknown model %s", option.Value)
			})
			if err != nil {
				return conf, err
			}
			conf.DefaultLLMModel = selectedModel.Value
			return conf, nil
		}
	}

	return conf, fmt.Errorf("unknown provider %s", selected.Value)
}
