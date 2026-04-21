package prompt

import "github.com/manboster/manboster/internal/repository"

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
