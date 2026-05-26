package gateway

import (
	"context"

	"github.com/manboster/manboster/spec/llm"
)

func (s *Service) LLMQuickChat(ctx context.Context, currentProvider llm.Provider, currentModel llm.Model, system string, user string) (*llm.Event, error) {
	msgList := []llm.Message{
		{
			Type: llm.MessageText,
			Role: llm.RoleSystem,
			Parts: []llm.MessageParts{
				{
					PartsType: llm.MessagePartsText,
					Text: &llm.MessageTextPayload{
						Text: system,
					},
				},
			},
		},
		{
			Type: llm.MessageText,
			Role: llm.RoleUser,
			Parts: []llm.MessageParts{
				{
					PartsType: llm.MessagePartsText,
					Text: &llm.MessageTextPayload{
						Text: user,
					},
				},
			},
		},
	}
	return s.LLMChat(ctx, currentProvider, currentModel, msgList, make(chan int, 10))
}
