package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

// modelIndexWithFallback gets provider and model's index with fallback provided
func (e *Engine) modelIndexWithFallback(ctx context.Context, targetProvider string, targetModel string) (int, int) {
	availProvider := false
	availModel := false
	providerIndex := 0
	modelIndex := 0

	for i, provider := range e.llmProviders {
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
		}
		break
	}

	if availProvider && availModel {
		return providerIndex, modelIndex
	}

	if !availModel && availProvider {
		color.Yellow(fmt.Sprintf("[Manboster Engine] Not found LLM model. We changed model to %s", e.llmProviders[providerIndex].Models()[0].DisplayName))
		return providerIndex, 0
	}

	return providerIndex, modelIndex
}
