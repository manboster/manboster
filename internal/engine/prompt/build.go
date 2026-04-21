package prompt

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/llm"
)

// Build TODO: build from a chat message to a llm message, make it easier to handle in engine
func (s *Service) Build(ctx context.Context, msg *chat.Message, sessionId string) (llm.Message, error) {
	return llm.Message{}, nil
}
