package gguf

import (
	"sync"
	"time"

	"github.com/hybridgroup/yzma/pkg/llama"
	"github.com/manboster/manboster/internal/hachimi"
	"github.com/manboster/manboster/spec/config"
)

type Service struct {
	manager      *Manager
	cfg          *Config
	ready        chan struct{}
	sampler      llama.Sampler
	vocab        llama.Vocab
	chatTemplate string
	chatLock     sync.Mutex
	lastUse      time.Time
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
	return &Config{}
}
