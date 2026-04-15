package types

import (
	"time"

	"github.com/manboster/manboster/internal/database/types"
)

type Chat struct {
	ID             uint64
	ChatID         string
	ChatProvider   string
	ChatModel      string
	ChatPermission int16
	SessionID      string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func MapC(chat Chat) types.Chat {
	return types.Chat{
		ID:             chat.ID,
		ChatID:         chat.ChatID,
		ChatProvider:   chat.ChatProvider,
		ChatModel:      chat.ChatModel,
		ChatPermission: chat.ChatPermission,
		SessionID:      chat.SessionID,
		CreatedAt:      chat.CreatedAt,
		UpdatedAt:      chat.UpdatedAt,
	}
}

func MapChat(chat types.Chat) Chat {
	return Chat{
		ID:             chat.ID,
		ChatID:         chat.ChatID,
		ChatProvider:   chat.ChatProvider,
		ChatModel:      chat.ChatModel,
		ChatPermission: chat.ChatPermission,
		SessionID:      chat.SessionID,
		CreatedAt:      chat.CreatedAt,
		UpdatedAt:      chat.UpdatedAt,
	}
}
