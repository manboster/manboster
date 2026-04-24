package loader

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
)

func LoadLLMProvider(ctx context.Context, llmConfig config.LLMConfig, provider llm.Provider) (llm.Provider, error) {
	newLProvider := provider.New()
	err := newLProvider.Init(ctx, llmConfig.Configuration)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Loader] Activate %q LLM Provider Error! Message: %q", newLProvider.DisplayName(), err))
	} else {
		color.Green(fmt.Sprintf("[Manboster Loader] Activate LLM Provider %q successful!", newLProvider.DisplayName()))
	}
	return newLProvider, nil
}

func LoadLLMProviders(ctx context.Context, llmConf []config.LLMConfig) ([]llm.Provider, error) {
	llmProviders := make([]llm.Provider, 0, len(llmConf))
	color.Cyan("[Manboster Loader] Loading LLM Providers...")
	// configure and init LLM providers
	for _, llmConfig := range llmConf {
		lProvider, err := llm.GetProvider(llmConfig.Provider)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Loader] There is no provider named %q when importing llm providers. Please check your configuration.", llmConfig.Provider))
			return nil, err
		}
		newLProvider, err := LoadLLMProvider(ctx, llmConfig, lProvider)
		// append it into array!
		llmProviders = append(llmProviders, newLProvider)
	}
	return llmProviders, nil
}
