package gatekeeper

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/gateway"
	"github.com/manboster/manboster/internal/engine/safeguard"
	"github.com/manboster/manboster/internal/hachimi"
	"github.com/manboster/manboster/internal/session"
	"github.com/manboster/manboster/spec/llm"
)

type Service struct {
	gatewayService   *gateway.Service
	safeguardService *safeguard.Service
	sessionService   *session.Service
	hachimiConfig    config.HachimiConfigs
	hachimiProvider  hachimi.Provider
	hachimiLoaded    *bool
	llmProviders     map[string]llm.Provider
}

func NewService(gatewayService *gateway.Service, safeguardService *safeguard.Service, hachimiConfig config.HachimiConfigs, hachimiProvider hachimi.Provider, hachimiLoaded *bool, sessionService *session.Service, llmProviders map[string]llm.Provider) *Service {
	return &Service{
		gatewayService:   gatewayService,
		safeguardService: safeguardService,
		hachimiConfig:    hachimiConfig,
		hachimiProvider:  hachimiProvider,
		hachimiLoaded:    hachimiLoaded,
		sessionService:   sessionService,
		llmProviders:     llmProviders,
	}
}
