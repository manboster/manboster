package providers

import (
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/llm/oai_compat"
	"github.com/manboster/manboster/internal/llm/openrouter"
)

// GetLLMProviders returned providers in LLMs
func GetLLMProviders() []llm.Provider {
	return []llm.Provider{
		&openrouter.Service{},
		&oai_compat.Service{},
	}
}
