package interact

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/cli"
)

// runOnboardLLMConfig runs LLM Config
func runOnboardLLMConfig(p cli.Provider) (config.LLMConfig, error) {
	conf := config.LLMConfig{}
	return conf, nil
}
