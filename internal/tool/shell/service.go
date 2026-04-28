package shell

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/hook"
	configType "github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/schema"
)

var metadata = schema.MetaData{
	Name:             "dev.manboster.shell",
	DisplayName:      "Shell Execution Tool",
	Description:      "[THIS IS A DANGEROUS TOOL, IF YOU DONT KNOW WHAT YOU ARE DOING PLEASE DO NOT USE OR INSTALL IT] shell execution tool enables model to execute native shells in the system. Root Access only.",
	MinEngineVersion: config.APILevel,
	AppVersion:       "0.0.1",
	APIVersion:       1,
	Requires:         nil,
	MinUserType:      "root",
}

type Service struct{}

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

func (s *Service) RegisterHook(registry *hook.Registry) {}
