package engine

import (
	"context"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
)

func (e *Engine) newChatData(ctx context.Context, sessionId string) error {
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

	e.sessionManager.AppendEvent(sessionId, event)

	err := e.writeChatData(ctx, event, sessionId)
	if err != nil {
		return err
	}

	return nil
}
