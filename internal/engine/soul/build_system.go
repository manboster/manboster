package soul

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/llm"
)

// BuildSystem returns system prompt message
func (s *Service) BuildSystem(ctx context.Context, msg *chat.Message, sessionId string) (llm.Message, error) {
	return llm.Message{}, nil
}
