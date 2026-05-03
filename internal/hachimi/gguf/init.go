package gguf

import "github.com/manboster/manboster/internal/hachimi"

func init() {
	hachimi.Register("hachimi-gguf", func() hachimi.Provider {
		return &Service{}
	})
}
