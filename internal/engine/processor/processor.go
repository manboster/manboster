package processor

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

type required interface {
	HandleMessage(ctx context.Context, instance chat.Provider, msg *chat.Message) error
}

// Service processes information passed from the application then go to the handler.
type Service struct {
	engine required
}

func New(engine required) *Service {
	return &Service{engine: engine}
}
