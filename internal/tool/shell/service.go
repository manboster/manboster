package shell

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/hook"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	configType "github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/schema"
)

var metadata = schema.MetaData{
	Name:               "dev.manboster.shell",
	DisplayName:        i18n.T(keys.ShellMachineDisplayName),
	Description:        i18n.T(keys.ShellMachineDescription),
	DisplayNameForUser: i18n.T(keys.ShellDisplayName),
	DescriptionForUser: i18n.T(keys.ShellDescription),
	MinEngineVersion:   config.APILevel,
	AppVersion:         "0.0.1",
	APIVersion:         1,
	Requires:           nil,
	Represent:          "💻",
	MinUserType:        schema.UserRoot,
}

type Service struct{}

func (s *Service) ClientRenderer(args string) string {
	arg := RunArgs{}
	if json.Unmarshal([]byte(args), &arg) == nil {
		return arg.Shell
	}
	return ""
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
	arg := RunArgs{}
	var respStr strings.Builder
	if json.Unmarshal([]byte(args), &arg) == nil {
		str := strings.Split(arg.Shell, " ")
		respStr.WriteString(str[0])
	}
	return respStr.String()
}
