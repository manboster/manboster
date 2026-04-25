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
	Model            string
	Provider         string
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
	MessagePayload   string // json encoded
	InputCost        float64
	OutputCost       float64
	TotalCost        float64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func MapCD(chatData ChatData) types.ChatData {
	return types.ChatData{
		ID:               chatData.ID,
		SessionID:        chatData.SessionID,
		Role:             string(chatData.Role),
		MessageType:      int16(chatData.MessageType),
		Model:            chatData.Model,
		Provider:         chatData.Provider,
		PromptTokens:     chatData.PromptTokens,
		CompletionTokens: chatData.CompletionTokens,
		TotalTokens:      chatData.TotalTokens,
		MessagePayload:   chatData.MessagePayload,
		InputCost:        chatData.InputCost,
		OutputCost:       chatData.OutputCost,
		TotalCost:        chatData.TotalCost,
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
		Model:            chatData.Model,
		Provider:         chatData.Provider,
		PromptTokens:     chatData.PromptTokens,
		CompletionTokens: chatData.CompletionTokens,
		TotalTokens:      chatData.TotalTokens,
		MessagePayload:   chatData.MessagePayload,
		InputCost:        chatData.InputCost,
		OutputCost:       chatData.OutputCost,
		TotalCost:        chatData.TotalCost,
		CreatedAt:        chatData.CreatedAt,
		UpdatedAt:        chatData.UpdatedAt,
	}
}
