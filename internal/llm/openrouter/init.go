package openrouter

import (
	"github.com/manboster/manboster/internal/llm"
	llmType "github.com/manboster/manboster/spec/llm"
)

func init() {
	llm.Register("openrouter", func() llmType.Provider {
		return &Service{}
	})
}
