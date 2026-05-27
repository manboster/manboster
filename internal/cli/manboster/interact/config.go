package interact

import (
	"fmt"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/cli"
)

type configLandingSelection string

const (
	configLandingChat    configLandingSelection = "chat"
	configLandingLLM     configLandingSelection = "llm"
	configLandingTool    configLandingSelection = "tool"
	configLandingHachimi configLandingSelection = "hachimi"
	configLandingApp     configLandingSelection = "app"
	configLandingQuit    configLandingSelection = _QUIT_
)

func (s configLandingSelection) Name() string {
	return string(s)
}

func (s configLandingSelection) DisplayName() string {
	switch s {
	case configLandingChat:
		return i18n.T(keys.CliConfigLandingChat)
	case configLandingLLM:
		return i18n.T(keys.CliConfigLandingLLM)
	case configLandingTool:
		return i18n.T(keys.CliConfigLandingTool)
	case configLandingHachimi:
		return i18n.T(keys.CliConfigLandingHachimi)
	case configLandingApp:
		return i18n.T(keys.CliConfigLandingApp)
	case configLandingQuit:
		return i18n.T(keys.CliConfigLandingQuit)
	default:
		return ""
	}
}

func runConfig(p cli.Provider, cfg config.Config) (config.Config, error) {
	se := []configLandingSelection{configLandingChat, configLandingLLM, configLandingTool, configLandingHachimi, configLandingApp, configLandingQuit}
	options := cli.BuildOptions[configLandingSelection](se, nil)
	mark := false

	form := newConfigForm[configLandingSelection]()
	form.Register(configLandingChat, func() error {
		chatConfigs, err := runChatConfigs(p, cfg)
		if err != nil {
			return err
		}

		cfg.Chats = chatConfigs
		return nil
	})

	form.Register(configLandingLLM, func() error {
		llmConfigs, err := runLLMConfigs(p, cfg)
		if err != nil {
			return err
		}

		cfg.LLMs = llmConfigs
		return nil
	})

	form.Register(configLandingTool, func() error {
		toolConfigs, err := runToolConfigs(p, cfg)
		if err != nil {
			return err
		}
		cfg.Tools = toolConfigs

		return nil
	})

	form.Register(configLandingHachimi, func() error {
		hachimiConfigs, err := runHachimiConfigs(p, cfg)
		if err != nil {
			return err
		}

		cfg.Hachimi = hachimiConfigs
		return nil
	})

	form.Register(configLandingApp, func() error {
		appConfig, err := runAppConfig(p, cfg)
		if err != nil {
			return err
		}
		dbPath := cfg.App.DBPath
		cfg.App = appConfig
		cfg.App.DBPath = dbPath
		return nil
	})

	form.Register(configLandingQuit, func() error {
		mark = true
		return nil
	})

	for {
		err := handle[configLandingSelection](p, form, options, i18n.T(keys.CliConfigLandingSelectPrompt), i18n.T(keys.CliConfigLandingSelectHelp))
		if err != nil {
			err := p.Alert(i18n.T(keys.CliWizardTitle), fmt.Sprintf(i18n.T(keys.CliWizardErrorAlert), err))
			if err != nil {
				return config.Config{}, err
			}
		}
		if mark {
			return cfg, nil
		}
	}
}
