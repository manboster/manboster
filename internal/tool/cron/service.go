package cron

import (
	"context"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/hook"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/repository"
	configType "github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/plugin"
	"github.com/manboster/manboster/spec/schema"
	"github.com/robfig/cron/v3"
)

var metadata = schema.MetaData{
	Name:             "dev.manboster.cron",
	DisplayName:      i18n.T(keys.CronDisplayName),
	Description:      i18n.T(keys.CronDescription),
	MinEngineVersion: config.APILevel,
	AppVersion:       "0.0.1",
	APIVersion:       1,
	Requires:         nil,
	MinUserType:      "admin",
}

var svc *Service

type Service struct {
	cronRepo repository.CronRepository
	manager  *Manager
	cron     *cron.Cron
}

func (s *Service) ClientRenderer(args string) string {
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
	return ""
}

func (s *Service) Run(ctx context.Context, args string) (*plugin.RunResponse, error) {
	return nil, nil
}

func (s *Service) Continue(ctx context.Context, session string) (*plugin.RunResponse, error) {
	return nil, nil
}
