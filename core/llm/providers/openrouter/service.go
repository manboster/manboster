package openrouter

import (
	"github.com/manboster/manboster/core/llm"
	"github.com/sashabaranov/go-openai"
)

type Service struct {
	cli *openai.Client
	cfg Config
}

func NewService(cli *openai.Client) *Service {
	return &Service{cli: cli}
}

func (s *Service) Name() string {
	return "openrouter"
}

func (s *Service) Model() string {
	return s.cfg.Model
}

func (s *Service) New() llm.Provider {
	return &Service{}
}
