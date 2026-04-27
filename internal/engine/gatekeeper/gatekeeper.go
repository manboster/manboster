package gatekeeper

import (
	"github.com/manboster/manboster/internal/engine/gateway"
	"github.com/manboster/manboster/internal/engine/safeguard"
	"github.com/manboster/manboster/internal/session/ignorance"
)

type Service struct {
	gatewayService          *gateway.Service
	safeguardService        *safeguard.Service
	ignoranceSessionManager *ignorance.Manager
}

func NewService(gatewayService *gateway.Service, safeguardService *safeguard.Service, ignoranceSessionManager *ignorance.Manager) *Service {
	return &Service{
		gatewayService:          gatewayService,
		safeguardService:        safeguardService,
		ignoranceSessionManager: ignoranceSessionManager,
	}
}
