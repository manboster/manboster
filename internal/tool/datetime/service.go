package datetime

import (
	"github.com/manboster/manboster/internal/config"
	configType "github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/schema"
)

var metadata = schema.MetaData{
	Name:             "dev.manboster.datetime",
	DisplayName:      "Get DateTime",
	Description:      "Datetime allows you to get current time and date of this machine.",
	MinEngineVersion: config.APILevel,
	AppVersion:       "0.0.1",
	APIVersion:       1,
	Requires:         nil,
	MinUserType:      "unknown",
}

type Service struct {
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
