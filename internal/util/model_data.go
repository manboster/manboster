package util

import (
	"context"

	"github.com/manboster/manboster/spec/llm"
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

	isOK := false
	for _, m := range provider.Models() {
		if m.Name == targetModel {
			model = m
			isOK = true
			break
		}
	}
	if !isOK {
		if len(provider.Models()) == 0 {
			return provider, model
		}
		model = provider.Models()[0]
	}

	return provider, model
}
