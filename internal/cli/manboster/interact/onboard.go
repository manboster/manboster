package interact

import (
	"fmt"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/cli"
)

type wizardCurrentState int8

const (
	wizardConfigChat wizardCurrentState = iota
	wizardConfigLLM
	wizardConfigApp
	wizardConfigTool
	wizardConfigHachimi
	wizardConfigPreview
	wizardConfigConfig
	wizardConfigWriting
	wizardConfigSuccess
	wizardConfigError
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

	chatConfigs, err := runOnboardChatConfigs(p)
	if err != nil {
		return conf, err
	}
	conf.Chats = chatConfigs

	llmConfigs, err := runOnboardLLMConfigs(p)
	if err != nil {
		return conf, err
	}
	conf.LLMs = llmConfigs

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

	hachimiConfigs, err := runOnboardHachimiConfigs(p)
	if err != nil {
		return conf, err
	}
	conf.Hachimi = hachimiConfigs

	conf = conf.Default()
	err = runOnboardPreview(p, conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
