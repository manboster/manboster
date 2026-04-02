package registry

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
)

// LoadLLMProviders loads chat's providers
func LoadLLMProviders(providers []llm.Provider, llmConfig []config.LLMConfig) ([]llm.Provider, error) {
	return []llm.Provider{}, nil
}

// LoadLLMConfigProviders loads configuration providers
func LoadLLMConfigProviders(providers []config.Provider) error {
	return nil
}
