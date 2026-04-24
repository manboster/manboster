package util

import (
	"context"

	"github.com/manboster/manboster/internal/llm"
)

// GetModelWithFallback gets provider and model's data with fallback provided
func GetModelWithFallback(ctx context.Context, llmProviders map[string]llm.Provider, targetProvider string, targetModel string) (llm.Provider, llm.Model) {
	var provider llm.Provider
	var model llm.Model
	p, avail := llmProviders[targetProvider]
	if !avail {
		for _, pr := range llmProviders {
			provider = pr
			break
		}
	} else {
		provider = p
	}
	for _, m := range provider.Models() {
		if m.Name == targetModel {
			model = m
			break
		}
	}
	return provider, model
}
