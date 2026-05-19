package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/config"
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
			err := p.Alert("Manboster Configuration Wizard", fmt.Sprintf("Failed to config %q", err))
			if err != nil {
				return nil, err
			}
			continue
		}
		confs = append(confs, llmConfig)

		ok, err := p.Prompt(fmt.Sprintf("You've added %d llm providers! Do you want to continue?", len(confs)), "", "Continue", "Exit and go on")
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
	options := util.BuildOptions[llmType.Provider](llmProviders, nil)
	llmProviderOption, err := p.Select("Next, let's pick a LLM provider. Which provider would you like to use?", "", options, "", func(option cli.Option) error {
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
