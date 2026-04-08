package llm

import "context"

// Provider defines which functions should LLM provides give.
type Provider interface {
	Chat(ctx context.Context, messages []Message) (*Event, error)              // now one.
	ChatStream(ctx context.Context, messages []Message) (<-chan *Event, error) // WIP: New Streaming chat
	Init(ctx context.Context, config any) error
	Name() string
	Model() string
	New() Provider
}
