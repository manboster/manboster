package openrouter

import (
	"context"

	"github.com/manboster/manboster/internal/llm"
)

// Chat allows you to chat with your model
func (s *Service) Chat(ctx context.Context, messages []llm.Message) (*llm.Event, error) {
	return s.oaiInstance.Chat(ctx, messages)
}

// ChatStream is the next generation async completion function
func (s *Service) ChatStream(ctx context.Context, messages []llm.Message) (<-chan *llm.Event, error) {
	return s.oaiInstance.ChatStream(ctx, messages)
}
