package openrouter

import (
	"context"

	"github.com/manboster/manboster/internal/llm"
)

// Chat allows you to chat with your model
func (s *Service) Chat(ctx context.Context, messages []llm.Message) (*llm.Message, error) {
	return s.oaiInstance.Chat(ctx, messages)
}
