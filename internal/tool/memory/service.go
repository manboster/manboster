package memory

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/spec/schema"
)

var metadata = schema.MetaData{
	Name:             "dev.manboster.memory",
	DisplayName:      "Memory Tools",
	Description:      "Memory allows you to storage the memory into database and read or write it anytime.",
	MinEngineVersion: config.APILevel,
	AppVersion:       "0.0.1",
	APIVersion:       1,
	Requires:         nil,
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
