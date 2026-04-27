package safeguard

import "github.com/manboster/manboster/internal/repository"

type Service struct {
	repo repository.Repository
}

// NewService create a safeguard instance
func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
