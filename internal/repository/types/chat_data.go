package types

import (
	"time"

	"github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/llm"
)

type ChatData struct {
	ID               uint64
	SessionID        string
	Role             llm.RoleType
	MessageType      llm.MessageType
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
	MessagePayload   string // json encoded
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func MapCD(chatData ChatData) types.ChatData {
	return types.ChatData{
		ID:               chatData.ID,
		SessionID:        chatData.SessionID,
		Role:             string(chatData.Role),
		MessageType:      int16(chatData.MessageType),
		PromptTokens:     chatData.PromptTokens,
		CompletionTokens: chatData.CompletionTokens,
		TotalTokens:      chatData.TotalTokens,
		MessagePayload:   chatData.MessagePayload,
		CreatedAt:        chatData.CreatedAt,
		UpdatedAt:        chatData.UpdatedAt,
	}
}

func MapChatData(chatData types.ChatData) ChatData {
	return ChatData{
		ID:               chatData.ID,
		SessionID:        chatData.SessionID,
		Role:             llm.RoleType(chatData.Role),
		MessageType:      llm.MessageType(chatData.MessageType),
		PromptTokens:     chatData.PromptTokens,
		CompletionTokens: chatData.CompletionTokens,
		TotalTokens:      chatData.TotalTokens,
		MessagePayload:   chatData.MessagePayload,
		CreatedAt:        chatData.CreatedAt,
		UpdatedAt:        chatData.UpdatedAt,
	}
}
