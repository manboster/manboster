package types

import (
	"github.com/manboster/manboster/internal/database/types"
	"github.com/manboster/manboster/internal/llm"
)

type ChatData struct {
	ID             uint64
	SessionID      string
	Role           llm.RoleType
	MessageType    llm.MessageType
	Tokens         int
	MessagePayload string // json encoded
}

func MapCD(chatData ChatData) types.ChatData {
	return types.ChatData{
		ID:             chatData.ID,
		SessionID:      chatData.SessionID,
		Role:           string(chatData.Role),
		MessageType:    int16(chatData.MessageType),
		Tokens:         chatData.Tokens,
		MessagePayload: chatData.MessagePayload,
	}
}

func MapChatData(chatData types.ChatData) ChatData {
	return ChatData{
		ID:             chatData.ID,
		SessionID:      chatData.SessionID,
		Role:           llm.RoleType(chatData.Role),
		MessageType:    llm.MessageType(chatData.MessageType),
		Tokens:         chatData.Tokens,
		MessagePayload: chatData.MessagePayload,
	}
}
