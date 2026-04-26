package openrouter

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm/oai_compat"
	"github.com/manboster/manboster/spec/llm"
)

type Service struct {
	oaiInstance *oai_compat.Service
}

func NewService(cli *oai_compat.Service) *Service {
	return &Service{oaiInstance: cli}
}

func (s *Service) Name() string {
	return s.oaiInstance.Name()
}

func (s *Service) Models() []llm.Model {
	return s.oaiInstance.Models()
}

func (s *Service) New() llm.Provider {
	return &Service{}
}

func (s *Service) DisplayName() string { return s.oaiInstance.DisplayName() }

func (s *Service) Stop() error { return s.oaiInstance.Stop() }

func (s *Service) Config() config.Provider { return &Config{} }

func (s *Service) Type() string { return s.oaiInstance.Type() }
