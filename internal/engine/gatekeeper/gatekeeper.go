package gatekeeper

import "github.com/manboster/manboster/internal/engine/gateway"

type Service struct {
	gatewayService *gateway.Service
}

func NewService(gatewayService *gateway.Service) *Service {
	return &Service{
		gatewayService: gatewayService,
	}
}
