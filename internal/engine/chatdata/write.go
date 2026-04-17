package chatdata

import (
	"context"
	"encoding/json"

	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository/types"
)

func (s *Service) Write(ctx context.Context, event llm.Event, sessionId string) error {
	var chatData types.ChatData

	if (event.EventType&llm.EventMessage == 0) && (event.EventType&llm.EventUsage == 0) {
		return nil
	}

	chatData.SessionID = sessionId
	if event.EventType&llm.EventMessage != 0 && event.Message != nil {
		jsonify, err := json.Marshal(event.Message)
		if err != nil {
			return err
		}

		chatData.MessageType = event.Message.Type
		chatData.MessagePayload = string(jsonify)
		chatData.Role = event.Message.Role
	}

	if event.EventType&llm.EventUsage != 0 && event.Usage != nil {
		chatData.PromptTokens = event.Usage.PromptTokens
		chatData.CompletionTokens = event.Usage.CompletionTokens
		chatData.TotalTokens = event.Usage.TotalTokens
	}

	return s.repo.CreateChatData(ctx, chatData)
}
