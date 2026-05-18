package interact

import (
	"fmt"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/cli"
)

func runOnboardLLMConfigs(p cli.Provider) ([]config.LLMConfig, error) {
	var confs []config.LLMConfig
	for {
		llmConfig, err := runOnboardLLMConfig(p)
		if err != nil {
			return confs, err
		}
		confs = append(confs, llmConfig)

		ok, err := p.Prompt(fmt.Sprintf("You've added %d llm providers! Do you want to continue?", len(confs)), "", "Continue", "Exit and go on")
		if err != nil {
			return confs, err
		}
		if !ok {
			break
		}
	}
	return confs, nil
}

// runOnboardLLMConfig runs LLM Config
func runOnboardLLMConfig(p cli.Provider) (config.LLMConfig, error) {
	conf := config.LLMConfig{}
	return conf, nil
}
