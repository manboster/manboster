package chatdata

import (
	"github.com/manboster/manboster/internal/engine/gateway"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session/chat_session"
	"github.com/manboster/manboster/spec/llm"
)

type Service struct {
	repo           repository.Repository
	sessionManager *chat_session.Manager
	llmProviders   map[string]llm.Provider
	gateway        *gateway.Service
}

func NewService(repo repository.Repository, sessionManager *chat_session.Manager, providers map[string]llm.Provider, gateway *gateway.Service) *Service {
	return &Service{
		repo:           repo,
		sessionManager: sessionManager,
		llmProviders:   providers,
		gateway:        gateway,
	}
}
