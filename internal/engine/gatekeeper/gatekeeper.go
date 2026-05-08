package gatekeeper

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/gateway"
	"github.com/manboster/manboster/internal/engine/safeguard"
	"github.com/manboster/manboster/internal/hachimi"
	"github.com/manboster/manboster/internal/session/ignorance"
)

type Service struct {
	gatewayService          *gateway.Service
	safeguardService        *safeguard.Service
	ignoranceSessionManager *ignorance.Manager
	hachimiConfig           config.HachimiConfigs
	hachimiProvider         hachimi.Provider
	hachimiLoaded           *bool
}

func NewService(gatewayService *gateway.Service, safeguardService *safeguard.Service, ignoranceSessionManager *ignorance.Manager, hachimiConfig config.HachimiConfigs, hachimiProvider hachimi.Provider, hachimiLoaded *bool) *Service {
	return &Service{
		gatewayService:          gatewayService,
		safeguardService:        safeguardService,
		ignoranceSessionManager: ignoranceSessionManager,
		hachimiConfig:           hachimiConfig,
		hachimiProvider:         hachimiProvider,
		hachimiLoaded:           hachimiLoaded,
	}
}
