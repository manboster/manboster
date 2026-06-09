package request

import (
	"context"
	"encoding/json"
	"net"
	"net/url"
	"strings"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/hook"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	configType "github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/schema"
	"golang.org/x/net/publicsuffix"
)

var metadata = schema.MetaData{
	Name:               "dev.manboster.request",
	DisplayName:        i18n.T(keys.RequestMachineDisplayName),
	Description:        i18n.T(keys.RequestMachineDescription),
	DisplayNameForUser: i18n.T(keys.RequestDisplayName),
	DescriptionForUser: i18n.T(keys.RequestDescription),
	MinEngineVersion:   config.APILevel,
	AppVersion:         "0.0.1",
	APIVersion:         1,
	Requires:           nil,
	Represent:          "🌐",
	MinUserType:        schema.UserAdmin,
}

type Service struct{}

func (s *Service) ClientRenderer(args string) string {
	arg := RunArgs{}
	if json.Unmarshal([]byte(args), &arg) == nil {
		return arg.URL
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
		parse, err := url.Parse(arg.URL)
		if err != nil {
			return ""
		}
		hostname := parse.Hostname()

		if ip := net.ParseIP(hostname); ip != nil {
			return hostname
		}

		rootDomain, err := publicsuffix.EffectiveTLDPlusOne(hostname)
		if err != nil {
			return ""
		}

		return rootDomain
	}
	return respStr.String()
}
