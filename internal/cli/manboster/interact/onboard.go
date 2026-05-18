package interact

import (
	"fmt"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/cli"
)

func runOnboardConfig(p cli.Provider) (config.Config, error) {
	conf := config.Config{}

	allowed, err := OnboardWarningPrompt(p)
	if err != nil {
		return conf, err
	}
	if !allowed {
		return conf, fmt.Errorf("you rejected the warning, in order to protect you, we skip your installation progress")
	}

	err = p.Alert("Manboster Configuration Wizard", "Welcome to the Manboster Configuration Wizard. Enjoy your experience with your little Manbo!")
	if err != nil {
		return conf, err
	}

	chatConfig, err := runOnboardChatConfig(p)
	if err != nil {
		return conf, err
	}
	conf.Chats = append(conf.Chats, chatConfig)

	for {
		llmConfig, err := runOnboardLLMConfig(p)
		if err != nil {
			return conf, err
		}
		conf.LLMs = append(conf.LLMs, llmConfig)

		ok, err := p.Prompt(fmt.Sprintf("You've added %d llm providers! Do you want to continue?", len(conf.LLMs)), "", "Continue", "Exit and go on")
		if err != nil {
			return conf, err
		}
		if !ok {
			break
		}
	}

	appConfig, err := runOnboardAPPConfig(p, conf)
	if err != nil {
		return conf, err
	}
	conf.App = appConfig

	toolConfigs, err := runOnboardToolConfig(p)
	if err != nil {
		return conf, err
	}
	conf.Tools = toolConfigs

	err = runOnboardPreview(p, conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
