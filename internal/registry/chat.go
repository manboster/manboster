package registry

import (
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
)

// LoadChatProviders loads chat's providers
func LoadChatProviders(providers []chat.Provider, chatConfig []config.ChatConfig) ([]chat.Provider, error) {
	return []chat.Provider{}, nil
}

// LoadChatConfigProviders loads configuration providers
func LoadChatConfigProviders(providers []config.Provider) error {
	return nil
}
