package providers

import (
	"github.com/manboster/manboster/internal/chat/telegram"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm/oai_compat"
	"github.com/manboster/manboster/internal/llm/openrouter"
)

// GetLLMConfigProviders returns configurations of LLMs.
func GetLLMConfigProviders() []config.Provider {
	return []config.Provider{
		&oai_compat.Config{},
		&openrouter.Config{},
	}
}

// GetChatConfigProviders returns configurations of Chats.
func GetChatConfigProviders() []config.Provider {
	return []config.Provider{
		&telegram.Config{},
	}
}
