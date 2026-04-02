package providers

import (
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/chat/telegram"
)

func GetChatProviders() []chat.Provider {
	return []chat.Provider{
		&telegram.Service{},
	}
}
