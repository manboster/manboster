package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/chat"
	_ "github.com/manboster/manboster/internal/chat/all"
	"github.com/manboster/manboster/internal/config"
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
			err := p.Alert("Manboster Configuration Wizard", fmt.Sprintf("Failed to config %q", err))
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

		ok, err := p.Prompt(fmt.Sprintf("You've added %d chat providers! Do you want to continue?", len(chatConfigs)), "", "Continue", "Exit and go on")
		if err != nil {
			return chatConfigs, err
		}
		if len(chatProviders) == 0 || !ok {
			if len(allChatProviders) == 0 {
				err := p.Alert("Manboster Configuration Wizard", "There is no more providers available for you to config...")
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
	options := util.BuildOptions[chatType.Provider](chatProviders, nil)
	chatProviderOption, err := p.Select("First, which platform would you like to use for your Manboster?", "", options, "", func(option cli.Option) error {
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
