package chat

import (
	"github.com/manboster/manboster/internal/engine/soul"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session"
)

type Service struct {
	soulService    *soul.Service
	repo           repository.Repository
	sessionManager *session.Manager
}

func NewService(soul *soul.Service, repo repository.Repository, manager *session.Manager) *Service {
	return &Service{
		soulService:    soul,
		repo:           repo,
		sessionManager: manager,
	}
}
