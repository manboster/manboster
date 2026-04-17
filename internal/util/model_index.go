package util

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/llm"
)

// GetModelIndexWithFallback gets provider and model's index with fallback provided
func GetModelIndexWithFallback(ctx context.Context, llmProviders []llm.Provider, targetProvider string, targetModel string) (int, int) {
	availProvider := false
	availModel := false
	providerIndex := 0
	modelIndex := 0

	for i, provider := range llmProviders {
		if provider.Name() == targetProvider {
			models := provider.Models()
			availProvider = true
			providerIndex = i

			for j, model := range models {
				if model.Name == targetModel {
					availModel = true
					modelIndex = j
					break
				}
			}
			break
		}
	}

	if availProvider && availModel {
		return providerIndex, modelIndex
	}

	if availProvider {
		color.Yellow(fmt.Sprintf("[Manboster Engine] Not found LLM model. We changed model to %s", llmProviders[providerIndex].Models()[0].DisplayName))
		return providerIndex, 0
	}

	return providerIndex, modelIndex
}
