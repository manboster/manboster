package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/llm"
	_ "github.com/manboster/manboster/internal/llm/all"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/cli"
	llmType "github.com/manboster/manboster/spec/llm"
)

func runOnboardLLMConfigs(p cli.Provider) ([]config.LLMConfig, error) {
	var confs []config.LLMConfig

	llmProviders := llm.AllProviders()
	for {
		llmConfig, err := runOnboardLLMConfig(p, llmProviders)
		if err != nil {
			err := p.Alert(i18n.T(keys.CliWizardTitle), fmt.Sprintf(i18n.T(keys.OnboardLLMConfigError), err))
			if err != nil {
				return nil, err
			}
			continue
		}
		confs = append(confs, llmConfig)

		ok, err := p.Prompt(fmt.Sprintf(i18n.T(keys.OnboardLLMAddedCount), len(confs)), "", i18n.T(keys.BtnContinue), i18n.T(keys.BtnExit))
		if err != nil {
			return confs, err
		}
		if !ok {
			break
		}
	}
	return confs, nil
}

// runOnboardLLMConfig runs LLM Config
func runOnboardLLMConfig(p cli.Provider, llmProviders []llmType.Provider) (config.LLMConfig, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.LLMConfig{}
	options := util.BuildOptionsForConfig[llmType.Provider](llmProviders, nil)
	llmProviderOption, err := p.Select(i18n.T(keys.OnboardLLMSelectPrompt), "", options, "", func(option cli.Option) error {
		for _, provider := range llmProviders {
			if provider.Config().Name() == option.Value {
				return nil
			}
		}
		return fmt.Errorf("no chat provider named %q", option.Value)
	})
	if err != nil {
		return conf, err
	}

	for _, provider := range llmProviders {
		if provider.Config().Name() == llmProviderOption.Value {
			cfg, err := CreateConfig(ctx, p, provider.Config())
			if err != nil {
				return conf, err
			}
			conf.Provider = provider.Config().Name()
			conf.Configuration = cfg
			return conf, nil
		}
	}

	return conf, nil
}
