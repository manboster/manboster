package file

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/manboster/manboster/internal/config"
	configType "github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/schema"
)

var metadata = schema.MetaData{
	Name:             "dev.manboster.file",
	DisplayName:      "Manboster File Helper",
	Description:      "This tool can read, write, list files, or get info about file. However, it can only read/write/getinfo within session's workspace or a shared public directory in workspace.",
	MinEngineVersion: config.APILevel,
	AppVersion:       "0.0.1",
	APIVersion:       1,
	Requires:         nil,
	MinUserType:      "admin",
}

type Service struct {
	cfg *Config
}

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
	return &Config{}
}

func (s *Service) Migrate(ctx context.Context, from int, conf any) (any, error) {
	return nil, nil
}

func (s *Service) CacheGroup(args string) string {
	arg := RunArgs{}
	var respStr strings.Builder
	if json.Unmarshal([]byte(args), &arg) == nil {
		switch arg.Name {
		case NameInfo, NameList, NameDir:
			respStr.WriteString("list")
		case NameRead:
			respStr.WriteString("read")
		case NameWrite, NameDelete:
			respStr.WriteString("write")
		}
		respStr.WriteString(":")
		jsonify, _ := json.Marshal(arg.FilePath)
		respStr.WriteString(string(jsonify))
	}
	return respStr.String()
}
