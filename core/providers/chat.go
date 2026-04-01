package providers

import (
	"github.com/manboster/manboster/core/chat"
	"github.com/manboster/manboster/core/chat/providers/telegram"
)

func GetChatProviders() []chat.Provider {
	return []chat.Provider{
		&telegram.Service{},
	}
}
