package providers

import (
	"github.com/manboster/manboster/core/llm"
	"github.com/manboster/manboster/core/llm/providers/oai_compat"
	"github.com/manboster/manboster/core/llm/providers/openrouter"
)

// GetLLMProviders returned providers in LLMs
func GetLLMProviders() []llm.Provider {
	return []llm.Provider{
		&openrouter.Service{},
		&oai_compat.Service{},
	}
}
