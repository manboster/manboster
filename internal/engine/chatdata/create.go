package chatdata

import (
	"context"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
)

// Create creates chat session with system prompt
func (s *Service) Create(ctx context.Context, sessionId string) error {
	textPayload := &llm.MessageTextPayload{
		Text: config.InitialSystemPrompt, // TODO: prompt engineering
	}

	event := llm.Event{
		EventType: llm.EventMessage,
		Message: &llm.Message{
			Role: llm.RoleSystem,
			Parts: []llm.MessageParts{
				{
					PartsType: llm.MessagePartsText,
					Text:      textPayload,
				},
			},
			Type: llm.MessageText,
		},
		Usage: nil,
	}

	s.sessionManager.AppendEvent(sessionId, event)

	err := s.Write(ctx, event, sessionId)
	if err != nil {
		return err
	}

	return nil
}
