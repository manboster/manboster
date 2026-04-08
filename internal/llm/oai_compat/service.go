package oai_compat

import (
	"github.com/manboster/manboster/internal/llm"
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
	return "oai-compat"
}

func (s *Service) Model() string {
	return s.cfg.Model.Name
}

func (s *Service) ModelInfo() llm.Model {
	return s.cfg.Model
}

func (s *Service) New() llm.Provider {
	return &Service{}
}
