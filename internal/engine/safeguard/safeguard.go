package safeguard

import (
	"github.com/manboster/manboster/internal/engine/onboard"
	"github.com/manboster/manboster/internal/repository"
)

type Service struct {
	repo    repository.Repository
	onboard *onboard.Service
}

// NewService create a safeguard instance
func NewService(repo repository.Repository, onboard *onboard.Service) *Service {
	return &Service{
		repo:    repo,
		onboard: onboard,
	}
}
