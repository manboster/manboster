package browser

import (
	"github.com/manboster/manboster/internal/config"
	configType "github.com/manboster/manboster/spec/config"
	"github.com/manboster/manboster/spec/schema"
)

var metadata = schema.MetaData{
	Name:             "dev.manboster.browser",
	DisplayName:      "Manboster Web Browser Tool",
	Description:      "Manboster Web Browser Tool gives unlimited ability for models to search for the Internet, capture webpages, surfing the Internet or download files. Also, you can use CloudFlare's browser use to proxy in order to hide real IP. If there is something you don't know, please use search to search for it.",
	MinEngineVersion: config.APILevel,
	AppVersion:       "0.0.1",
	APIVersion:       1,
	Requires:         nil,
	MinUserType:      "unknown",
}

type Service struct {
	Manager *Manager
	isReady bool
	cfg     *Config
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
