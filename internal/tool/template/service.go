package template

import (
	"context"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/hook"
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

type Service struct{}

func (s *Service) ClientRenderer(args string) string {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Description() string {
	return metadata.Description
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

func (s *Service) RegisterHook(registry *hook.Registry) {}

func (s *Service) Migrate(ctx context.Context, from int, conf any) (any, error) {
	return nil, nil
}

func (s *Service) CacheGroup(args string) string {
	return ""
}
