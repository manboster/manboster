package interact

import (
	"fmt"

	"github.com/manboster/manboster/internal/config"
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
		return "Chat Providers\nAdd, edit or delete your chat providers."
	case configLandingLLM:
		return "LLM Providers\nAdd, edit or delete your llm providers."
	case configLandingTool:
		return "Tool Providers\nAdd, edit or delete your system tool providers."
	case configLandingHachimi:
		return "Hachimi Settings\nAdd, edit or delete your Hachimi providers or modify Hachimi settings."
	case configLandingApp:
		return "App Settings\nModify Manboster settings."
	case configLandingQuit:
		return "Quit\nBye!"
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
		err := handle[configLandingSelection](p, form, options, "Please select what to configure in configuration", "Please choose what you want to configure in configuration field.")
		if err != nil {
			err := p.Alert("Manboster Configuration Wizard", fmt.Sprintf("We encountered an error while configuring: %q", err))
			if err != nil {
				return config.Config{}, err
			}
		}
		if mark {
			return cfg, nil
		}
	}
}
