package openrouter

import (
	"context"

	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/tool"
)

// Chat allows you to chat with your model
func (s *Service) Chat(ctx context.Context, model string, tools []tool.Provider, messages []llm.Message) (*llm.Event, error) {
	return s.oaiInstance.Chat(ctx, model, tools, messages)
}

// ChatStream is the next generation async completion function
func (s *Service) ChatStream(ctx context.Context, model string, tools []tool.Provider, messages []llm.Message) (<-chan *llm.Event, error) {
	return s.oaiInstance.ChatStream(ctx, model, tools, messages)
}
