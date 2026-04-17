package oai_compat

import "github.com/manboster/manboster/internal/llm"

func init() {
	llm.Register("oai-compat", func() llm.Provider {
		return &Service{}
	})
}
