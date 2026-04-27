package gatekeeper

import (
	"github.com/manboster/manboster/internal/engine/gateway"
	"github.com/manboster/manboster/internal/engine/safeguard"
)

type Service struct {
	gatewayService   *gateway.Service
	safeguardService *safeguard.Service
}

func NewService(gatewayService *gateway.Service, safeguardService *safeguard.Service) *Service {
	return &Service{
		gatewayService:   gatewayService,
		safeguardService: safeguardService,
	}
}
