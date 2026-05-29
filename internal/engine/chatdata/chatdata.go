package chatdata

import (
	"github.com/manboster/manboster/internal/engine/gateway"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/session"
	"github.com/manboster/manboster/internal/session/chat_session"
	"github.com/manboster/manboster/spec/llm"
)

type Service struct {
	repo           repository.Repository
	sessionManager *chat_session.Manager
	sessionService *session.Service
	llmProviders   map[string]llm.Provider
	gateway        *gateway.Service
}

func NewService(repo repository.Repository, sessionManager *chat_session.Manager, sessionService *session.Service, providers map[string]llm.Provider, gateway *gateway.Service) *Service {
	return &Service{
		repo:           repo,
		sessionManager: sessionManager,
		sessionService: sessionService,
		llmProviders:   providers,
		gateway:        gateway,
	}
}
