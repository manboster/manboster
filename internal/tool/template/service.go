package template

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/repository"
	configType "github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/schema"
)

var metadata = schema.MetaData{
	Name:             "dev.manboster.template",
	DisplayName:      "Tool Template",
	Description:      "This is a template of tool call, you can copy and paste its code.",
	MinEngineVersion: config.APILevel,
	AppVersion:       "0.0.0",
	APIVersion:       -1,
	Requires:         nil,
	MinUserType:      "",
}

type Service struct {
	memDB repository.MemoryRepository
}

func (s *Service) Name() string {
	return metadata.Name
}

func (s *Service) DisplayName() string {
	return metadata.DisplayName
}

func (s *Service) MetaData() schema.MetaData {
	return metadata
}

func (s *Service) Requires() []schema.RequirementData {
	return metadata.Requires
}

func (s *Service) Config() configType.Provider {
	return nil
}
