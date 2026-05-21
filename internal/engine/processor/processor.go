package processor

import (
	"context"

	"github.com/manboster/manboster/internal/engine/onboard"
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
	sessionService   *session.Service
	safeguardService *safeguard.Service
	onboard          *onboard.Service
}

func New(engine required, sessionService *session.Service, safeguardService *safeguard.Service, onboardService *onboard.Service) *Service {
	return &Service{
		engine:           engine,
		sessionService:   sessionService,
		safeguardService: safeguardService,
		onboard:          onboardService,
	}
}
