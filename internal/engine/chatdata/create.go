package chatdata

import (
	"context"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
)

// CreateWithSystemPrompt creates chat session with additional system prompt
func (s *Service) CreateWithSystemPrompt(ctx context.Context, sessionId string, prompt string) error {
	textPayload := &llm.MessageTextPayload{
		Text: config.InitialSystemPrompt + prompt, // TODO: prompt engineering
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

// Create creates chat session with system prompt
func (s *Service) Create(ctx context.Context, sessionId string) error {
	return s.CreateWithSystemPrompt(ctx, sessionId, "")
}
