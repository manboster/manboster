package file

import (
	"context"

	"github.com/manboster/manboster/internal/config"
	configType "github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/schema"
)

var metadata = schema.MetaData{
	Name:             "dev.manboster.file",
	DisplayName:      "Manboster File Helper",
	Description:      "You can use this tool to read, write, list files, or get info about file. However, it can only read/write/getinfo within session's workspace directory.",
	MinEngineVersion: config.APILevel,
	AppVersion:       "0.0.1",
	APIVersion:       1,
	Requires:         nil,
	MinUserType:      "admin",
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

func (s *Service) Migrate(ctx context.Context, from int, conf any) (any, error) {
	return nil, nil
}
