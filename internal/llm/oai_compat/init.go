package oai_compat

import (
	"github.com/manboster/manboster/internal/llm"
	llmType "github.com/manboster/manboster/spec/llm"
)

func init() {
	llm.Register("oai-compat", func() llmType.Provider {
		return &Service{}
	})
}
