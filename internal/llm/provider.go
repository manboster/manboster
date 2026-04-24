package llm

import (
	"context"

	"github.com/manboster/manboster/internal/config"
)

// Provider defines which functions should LLM provides give.
type Provider interface {
	Chat(ctx context.Context, model string, messages []Message) (*Event, error)              // now one.
	ChatStream(ctx context.Context, model string, messages []Message) (<-chan *Event, error) // WIP: New Streaming chat
	Init(ctx context.Context, config any) error
	Stop() error
	Name() string // identifiable name
	Type() string // OpenAI / anthropic / Gemini etc.
	DisplayName() string
	Config() config.Provider
	Models() []Model
	New() Provider
}
