package openrouter

import "github.com/manboster/manboster/internal/llm"

func init() {
	llm.Register("openrouter", func() llm.Provider {
		return &Service{}
	})
}
