package interact

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/cli"
)

type wizardCurrentState int8

const (
	wizardConfigHello wizardCurrentState = iota
	wizardConfigChat
	wizardConfigLLM
	wizardConfigApp
	wizardConfigTool
	wizardConfigHachimi
	wizardConfigPreview
	wizardConfigWrite
	wizardConfigError
	wizardConfigSuccess
)

func runOnboardConfig(p cli.Provider) (config.Config, error) {
	conf := config.Config{}

	if err := config.Init(); err == nil {
		prompt, err := p.Prompt(i18n.T(keys.OnboardExistingConfig), i18n.T(keys.OnboardExistingContinue), "Continue", "Exit")
		if err != nil {
			return config.Config{}, err
		}
		if !prompt {
			return config.Config{}, fmt.Errorf(i18n.T(keys.OnboardUserCancelled))
		}
	}

	allowed, err := OnboardWarningPrompt(p)
	if err != nil {
		return conf, err
	}
	if !allowed {
		return conf, fmt.Errorf(i18n.T(keys.OnboardWarningRejected))
	}

	state := wizardConfigHello
	lastState := wizardConfigHello
	var reportedError error

	for state != wizardConfigSuccess {
		switch state {
		case wizardConfigHello:
			err = p.Alert(i18n.T(keys.WizardTitle), i18n.T(keys.WizardWelcome))
			if err != nil {
				lastState = state
				state = wizardConfigError
				reportedError = err
				continue
			}

			state = wizardConfigChat
		case wizardConfigChat:
			chatConfigs, err := runOnboardChatConfigs(p)
			if err != nil {
				lastState = state
				state = wizardConfigError
				reportedError = err
				continue
			}

			conf.Chats = chatConfigs
			state = wizardConfigLLM
		case wizardConfigLLM:
			llmConfigs, err := runOnboardLLMConfigs(p)
			if err != nil {
				lastState = state
				state = wizardConfigError
				reportedError = err
				continue
			}

			conf.LLMs = llmConfigs
			state = wizardConfigApp
		case wizardConfigApp:
			appConfig, err := runOnboardAPPConfig(p, conf)
			if err != nil {
				lastState = state
				state = wizardConfigError
				reportedError = err
				continue
			}

			conf.App = appConfig
			state = wizardConfigTool
		case wizardConfigTool:
			toolConfigs, err := runOnboardToolConfig(p)
			if err != nil {
				lastState = state
				state = wizardConfigError
				reportedError = err
				continue
			}
			conf.Tools = toolConfigs
			state = wizardConfigHachimi
		case wizardConfigHachimi:
			hachimiConfigs, err := runOnboardHachimiConfigs(p)
			if err != nil {
				lastState = state
				state = wizardConfigError
				reportedError = err
				continue
			}

			conf.Hachimi = hachimiConfigs
			state = wizardConfigPreview
		case wizardConfigPreview:
			conf = conf.Default()
			confirm, err := runOnboardPreview(p, conf)
			if err != nil {
				lastState = state
				state = wizardConfigError
				reportedError = err
				continue
			}

			if !confirm {
				conf, err = runConfig(p, conf)
				if err != nil {
					return conf, err
				}
			}
			state = wizardConfigWrite
		case wizardConfigWrite:
			err = runConfigWrite(p, conf)
			if err != nil {
				lastState = state
				state = wizardConfigError
				reportedError = err
				continue
			}
			state = wizardConfigSuccess
		// If you are using GoLand or other JetBrains IDEs, please ignore this `condition is always true` error.
		case wizardConfigError:
			confirm, err := p.Prompt(fmt.Sprintf(i18n.T(keys.WizardErrorRetry), reportedError), "Do you want to retry?", "Retry", "Exit")
			if err != nil {
				color.Red(fmt.Sprintf(i18n.T(keys.WizardConfigError), reportedError))
			}
			if !confirm {
				return conf, reportedError
			}
			state = lastState
		case wizardConfigSuccess:
			err := p.Alert(i18n.T(keys.WizardTitle), i18n.T(keys.WizardSuccess))
			if err != nil {
				return config.Config{}, err
			}
		default:
			return conf, fmt.Errorf("unknown wizard state: %d", state)
		}

	}

	return conf, nil
}
