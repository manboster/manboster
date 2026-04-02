package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
)

func loadLLM(ctx context.Context, llmConf []config.LLMConfig) ([]llm.Provider, error) {
	llmProviders := make([]llm.Provider, 0, len(llmConf))
	// configure and init LLM providers
	for _, llmConfigs := range llmConf {
		lProvider, err := llm.GetProvider(llmConfigs.Provider)
		if err != nil {
			color.Red(fmt.Sprintf("There is no provider named %q when importing llm providers. Please check your configuration.", llmConfigs.Provider))
			return nil, err
		}
		newLProvider := lProvider.New()
		err = newLProvider.Init(ctx, llmConfigs.Configuration)
		if err != nil {
			color.Red("Activate ", lProvider.Name(), " LLM API Error! Message:", err.Error())
		}

		// append it into array!
		llmProviders = append(llmProviders, newLProvider)
	}
	return llmProviders, nil
}
