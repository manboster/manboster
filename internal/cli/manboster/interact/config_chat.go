package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
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
		return i18n.T(keys.CliConfigActionDeleteProvider)
	case chatConfigEdit:
		return i18n.T(keys.CliConfigActionEditProvider)
	case chatConfigQuit:
		return i18n.T(keys.CliConfigActionQuit)
	default:
		return ""
	}
}

func runChatConfigs(p cli.Provider, cfg config.Config) ([]config.ChatConfig, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var option cli.Option
	for {
		// reload on every iteration so changes are reflected
		var chatProviders []chatType.Provider
		for _, c := range cfg.Chats {
			provider, err := chat.GetProvider(c.Provider)
			if err != nil {
				return nil, err
			}
			if err := provider.Init(ctx, c.Configuration); err != nil {
				continue
			}
			chatProviders = append(chatProviders, provider)
		}

		allChatProviders := chat.AllProviders()
		for i, provider := range chatProviders {
			for _, ch := range allChatProviders {
				if ch.Config().Name() == provider.Config().Name() {
					allChatProviders = append(allChatProviders[:i], allChatProviders[i+1:]...)
					break
				}
			}
		}

		options := cli.BuildOptions[chatType.Provider](chatProviders, nil)
		options = append(options, quitOption)
		if len(allChatProviders) > 0 {
			options = append([]cli.Option{addOption}, options...)
		}

		var err error
		option, err = p.Select(i18n.T(keys.CliConfigChatSelectPrompt), i18n.T(keys.CliConfigChatSelectHelp), options, option.Value, func(option cli.Option) error {
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
			continue
		}

		var selectedConfig config.ChatConfig
		var selectedProvider chatType.Provider
		selectedIndex := -1
		for i, c := range cfg.Chats {
			if c.Provider == option.Value {
				selectedConfig = c
				selectedIndex = i
				pr, err := chat.GetProvider(c.Provider)
				if err != nil {
					return nil, err
				}
				selectedProvider = pr
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
			confirm, err := p.Prompt(i18n.Te(keys.CliConfigChatDeleteConfirm, selectedConfig.Provider, nil), "Do you want to continue?", "Yes", "No")
			if err != nil {
				return err
			}
			if !confirm {
				return fmt.Errorf("cancelled")
			}
			cfg.Chats = append(cfg.Chats[:selectedIndex], cfg.Chats[selectedIndex+1:]...)
			if err := p.Alert(i18n.T(keys.CliWizardTitle), i18n.Te(keys.CliConfigChatDeleteSuccess, selectedConfig.Provider, nil)); err != nil {
				return err
			}
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

		err = handleWithPrompt[chatConfigAction](p, form, opts, fmt.Sprintf("This chat provider %s's info:\n\n%s", selectedProvider.DisplayName(), selectedConfig.Configuration), i18n.T(keys.CliConfigActionWhatToDo))
		if err != nil {
			return nil, err
		}
	}
}
