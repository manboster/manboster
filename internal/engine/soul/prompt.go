package soul

import (
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
)

type Service struct {
	repo    repository.Repository
	soulMap map[string]types.Soul
}

func New(repo repository.Repository) *Service {
	return &Service{
		repo:    repo,
		soulMap: make(map[string]types.Soul),
	}
}
