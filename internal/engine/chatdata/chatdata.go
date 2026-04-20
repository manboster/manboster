package chatdata

import (
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session"
)

type Service struct {
	repo           repository.Repository
	sessionManager *session.Manager
	llmProviders   []llm.Provider
}

func New(repo repository.Repository, sessionManager *session.Manager, providers []llm.Provider) *Service {
	return &Service{
		repo:           repo,
		sessionManager: sessionManager,
		llmProviders:   providers,
	}
}
