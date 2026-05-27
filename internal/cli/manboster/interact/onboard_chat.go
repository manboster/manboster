package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/chat"
	_ "github.com/manboster/manboster/internal/chat/all"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/util"
	chatType "github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/cli"
)

// runOnboardChatConfigs
func runOnboardChatConfigs(p cli.Provider) ([]config.ChatConfig, error) {
	var chatConfigs []config.ChatConfig

	allChatProviders := chat.AllProviders()
	chatProviders := allChatProviders
	for {
		chatConfig, err := runOnboardChatConfig(p, chatProviders)
		if err != nil {
			err := p.Alert(i18n.T(keys.CliWizardTitle), fmt.Sprintf(i18n.T(keys.OnboardChatConfigError), err))
			if err != nil {
				return chatConfigs, err
			}
			continue
		}

		chatConfigs = append(chatConfigs, chatConfig)
		for i, provider := range chatProviders {
			if provider.Name() == chatConfig.Provider {
				chatProviders = append(chatProviders[:i], chatProviders[i+1:]...)
				break
			}
		}

		ok, err := p.Prompt(fmt.Sprintf(i18n.T(keys.OnboardChatAddedCount), len(chatConfigs)), "", i18n.T(keys.UIBtnContinue), i18n.T(keys.UIBtnExit))
		if err != nil {
			return chatConfigs, err
		}
		if len(chatProviders) == 0 || !ok {
			if len(chatProviders) == 0 && ok {
				err := p.Alert(i18n.T(keys.CliWizardTitle), i18n.T(keys.OnboardChatNoMoreProviders))
				if err != nil {
					return chatConfigs, err
				}
			}
			break
		}
	}
	return chatConfigs, nil
}

// runOnboardChatConfig runs chat Config
func runOnboardChatConfig(p cli.Provider, chatProviders []chatType.Provider) (config.ChatConfig, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.ChatConfig{}
	options := util.BuildOptionsForConfig[chatType.Provider](chatProviders, nil)
	chatProviderOption, err := p.Select(i18n.T(keys.OnboardChatSelectPrompt), "", options, "", func(option cli.Option) error {
		for _, provider := range chatProviders {
			if provider.Name() == option.Value {
				return nil
			}
		}
		return fmt.Errorf("no chat provider named %q", option.Value)
	})
	if err != nil {
		return conf, err
	}

	for _, provider := range chatProviders {
		if provider.Name() == chatProviderOption.Value {
			cfg, err := CreateConfig(ctx, p, provider.Config())
			if err != nil {
				return conf, err
			}
			conf.Provider = provider.Name()
			conf.Configuration = cfg
			return conf, nil
		}
	}

	return conf, fmt.Errorf("no chat provider named %q", chatProviderOption.Value)
}
