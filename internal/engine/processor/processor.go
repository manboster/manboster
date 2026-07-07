package processor

import (
	"context"

	chatEngine "github.com/manboster/manboster/internal/engine/chat"
	"github.com/manboster/manboster/internal/engine/safeguard"
	"github.com/manboster/manboster/internal/session"
	"github.com/manboster/manboster/spec/chat"
)

type required interface {
	Distribute(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error
}

// Service processes information passed from the application then go to the handler.
type Service struct {
	engine           required
	sessionManager   *session.Manager
	chatService      *chatEngine.Service
	safeguardService *safeguard.Service
}

func New(engine required, sessionManager *session.Manager, safeguardService *safeguard.Service, chatService *chatEngine.Service) *Service {
	return &Service{
		engine:           engine,
		sessionManager:   sessionManager,
		safeguardService: safeguardService,
		chatService:      chatService,
	}
}
