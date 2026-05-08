package gguf

import (
	"github.com/manboster/manboster/internal/hachimi"
	"github.com/manboster/manboster/spec/config"
)

type Service struct {
	avail      bool
	availModel bool
}

func (s *Service) Name() string {
	return "hachimi-gguf"
}

func (s *Service) DisplayName() string {
	return "hachimi gguf runtime"
}

func (s *Service) New() hachimi.Provider {
	return &Service{}
}

func (s *Service) Config() config.Provider {
	return nil
}
