package types

import "github.com/manboster/manboster/internal/database/types"

type Chat struct {
	ID             uint64
	ChatID         string
	ChatProvider   string
	ChatModel      string
	ChatPermission int16
	SessionID      string
}

func MapC(chat Chat) types.Chat {
	return types.Chat{
		ID:             chat.ID,
		ChatID:         chat.ChatID,
		ChatProvider:   chat.ChatProvider,
		ChatModel:      chat.ChatModel,
		ChatPermission: chat.ChatPermission,
		SessionID:      chat.SessionID,
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
	}
}
