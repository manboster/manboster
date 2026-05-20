package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
	chatType "github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/cli"
)

type chatConfigAction string

const (
	chatConfigDelete chatConfigAction = _DELETE_
	chatConfigEdit   chatConfigAction = _EDIT_
	chatConfigQuit   chatConfigAction = _QUIT_
)

func (a chatConfigAction) Name() string {
	return string(a)
}

func (a chatConfigAction) DisplayName() string {
	switch a {
	case chatConfigDelete:
		return "Delete this provider"
	case chatConfigEdit:
		return "Edit this provider"
	case chatConfigQuit:
		return "Quit"
	default:
		return ""
	}
}

func runChatConfigs(p cli.Provider, cfg config.Config) ([]config.ChatConfig, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var chatProviders []chatType.Provider
	for _, c := range cfg.Chats {
		provider, err := chat.GetProvider(c.Provider)
		if err != nil {
			return nil, err
		}

		err = provider.Init(ctx, c.Configuration)
		if err != nil {
			continue
		}

		chatProviders = append(chatProviders, provider)
	}

	options := cli.BuildOptions[chatType.Provider](chatProviders, nil)
	options = append(options, quitOption)

	allChatProviders := chat.AllProviders()
	for i, provider := range chatProviders {
		for _, ch := range allChatProviders {
			if ch.Config().Name() == provider.Config().Name() {
				allChatProviders = append(allChatProviders[:i], allChatProviders[i+1:]...)
				break
			}
		}
	}
	if len(allChatProviders) > 0 {
		options = append(options, addOption)
	}

	var option cli.Option
	for {
		var err error
		option, err = p.Select("Select a chat provider to configure.", "Please select a chat provider to configure.", options, option.Value, func(option cli.Option) error {
			for _, o := range options {
				if o.Value == option.Value {
					return nil
				}
			}
			return fmt.Errorf("unknown chat provider selected: %s", option.Value)
		})
		if err != nil {
			return nil, err
		}

		if option.Value == _QUIT_ {
			return cfg.Chats, nil
		}

		if option.Value == _ADD_ {
			chatConfig, err := runOnboardChatConfig(p, allChatProviders)
			if err != nil {
				return nil, err
			}
			cfg.Chats = append(cfg.Chats, chatConfig)
		}

		var selectedConfig config.ChatConfig
		var selectedProvider chatType.Provider
		selectedIndex := -1
		for i, c := range cfg.Chats {
			if c.Provider == option.Value {
				selectedConfig = c
				selectedIndex = i
				p, err := chat.GetProvider(c.Provider)
				if err != nil {
					return nil, err
				}
				selectedProvider = p
				break
			}
		}
		if selectedIndex == -1 {
			return nil, fmt.Errorf("unknown chat provider selected: %s", option.Value)
		}

		se := []chatConfigAction{chatConfigEdit, chatConfigDelete, chatConfigQuit}
		opts := cli.BuildOptions[chatConfigAction](se, nil)
		form := newConfigForm[chatConfigAction]()

		form.Register(chatConfigDelete, func() error {
			confirm, err := p.Prompt(fmt.Sprintf("Do you want to delete %q?\n\nYour action is IRREVERSIBLE!", selectedConfig.Provider), "Do you want to continue?", "Yes", "No")
			if err != nil {
				return err
			}
			if !confirm {
				return fmt.Errorf("cancelled")
			}
			cfg.Chats = append(cfg.Chats[:selectedIndex], cfg.Chats[selectedIndex+1:]...)
			return errQuit
		})

		form.Register(chatConfigEdit, func() error {
			conf, err := EditConfig(ctx, p, selectedProvider.Config(), selectedConfig.Configuration)
			if err != nil {
				return err
			}
			selectedConfig.Configuration = conf
			cfg.Chats[selectedIndex] = selectedConfig
			return errQuit
		})

		form.Register(chatConfigQuit, nilFunc)

		err = handleWithPrompt[chatConfigAction](p, form, opts, fmt.Sprintf("This chat provider %s's info:\n\n%s", selectedProvider.DisplayName(), selectedConfig.Configuration), "What do you want to do with it?")
		if err != nil {
			return nil, err
		}
	}
}
