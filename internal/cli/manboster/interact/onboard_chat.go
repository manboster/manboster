package interact

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/cli"
)

// runOnboardChatConfig runs LLM Config
func runOnboardChatConfig(p cli.Provider) (config.ChatConfig, error) {
	conf := config.ChatConfig{}
	return conf, nil
}
