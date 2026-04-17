package chatdata

import (
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session"
)

type Service struct {
	repo           repository.Repository
	sessionManager *session.Manager
}

func New(repo repository.Repository, sessionManager *session.Manager) *Service {
	return &Service{
		repo:           repo,
		sessionManager: sessionManager,
	}
}
