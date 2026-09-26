package gguf

import (
	"github.com/manboster/manboster/internal/hachimi"
	specHachimi "github.com/manboster/manboster/spec/hachimi"
)

func init() {
	hachimi.Register("hachimi-gguf", func() specHachimi.Provider {
		return &Service{}
	})
}
