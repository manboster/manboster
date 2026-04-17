package safeguard

import "github.com/manboster/manboster/internal/repository"

type Service struct {
	repo repository.Repository
}

// New create a safeguard instance
func New(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
