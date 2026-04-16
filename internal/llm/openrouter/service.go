package openrouter

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/llm/oai_compat"
)

type Service struct {
	oaiInstance *oai_compat.Service
}

func NewService(cli *oai_compat.Service) *Service {
	return &Service{oaiInstance: cli}
}

func (s *Service) Name() string {
	return "openrouter"
}

func (s *Service) Models() []llm.Model {
	return s.oaiInstance.Models()
}

func (s *Service) New() llm.Provider {
	return &Service{}
}

func (s *Service) DisplayName() string { return "OpenRouter" }

func (s *Service) Stop() error { return s.oaiInstance.Stop() }

func (s *Service) Config() config.Provider { return &Config{} }
