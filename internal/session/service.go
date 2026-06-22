package session

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/engine/onboard"
	"github.com/manboster/manboster/internal/engine/soul"
	"github.com/manboster/manboster/internal/repository"
)

type Service struct {
	Manager     *Manager
	repo        repository.Repository
	soulService *soul.Service
	config      *config.Config
	onboard     *onboard.Service
}

func NewService(repo repository.Repository, soulService *soul.Service, conf *config.Config, onboardService *onboard.Service) *Service {
	return &Service{
		Manager:     NewManager(),
		repo:        repo,
		soulService: soulService,
		config:      conf,
		onboard:     onboardService,
	}
}
