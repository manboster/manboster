package chatdata

import (
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session"
	"github.com/manboster/manboster/spec/llm"
)

type Service struct {
	repo           repository.Repository
	sessionManager *session.Manager
	llmProviders   map[string]llm.Provider
}

func New(repo repository.Repository, sessionManager *session.Manager, providers map[string]llm.Provider) *Service {
	return &Service{
		repo:           repo,
		sessionManager: sessionManager,
		llmProviders:   providers,
	}
}
