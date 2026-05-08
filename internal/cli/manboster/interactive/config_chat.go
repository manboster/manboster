package interactive

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/config"
	chatType "github.com/manboster/manboster/spec/chat"
)

func configLandingChatActionForm() (configLandingActionSelection, error) {
	return configLandingActionForm("Add a new chat provider", "Select an existing chat provider", "Quit")
}

func runLandingChatActionForm(ctx context.Context) error {
	defer helper.ClearScreen()

	conf := config.Read()

	printConfigChatProvidersData(ctx)
	se, err := configLandingChatActionForm()
	if err != nil {
		return err
	}
	switch se {
	case configLandingActionAdd:
		providerAvail := chat.AvailProviders()
		var providerList []chatType.Provider

		occupy := make(map[string]bool)
		for _, c := range conf.Chats {
			occupy[c.Provider] = true
		}

		for _, provider := range providerAvail {
			if !occupy[provider] {
				p, err := chat.GetProvider(provider)
				if err != nil {
					color.Yellow(fmt.Sprintf("[Manboster Client] Failed to get Chat provider %s: %q", provider, err))
					continue
				}
				providerList = append(providerList, p)
			}
		}

		helper.ClearScreen()
		if len(providerList) == 0 {
			color.Yellow("[Manboster Client] No new chat providers available to add!")
			time.Sleep(1 * time.Second)
			return nil
		}

		chatProvider, err := SelectChatForm(ctx, providerList, "Please select a provider you want to add:")
		if err != nil {
			return err
		}

		cf, err := RunOnboardConfig(ctx, chatProvider.Config())
		if err != nil {
			return err
		}

		conf.Chats = append(conf.Chats, config.ChatConfig{
			Provider:      chatProvider.Name(),
			Configuration: cf,
		})

		err = config.Write(conf)
		if err != nil {
			return err
		}
	case configLandingActionSelect:
		var providerList []chatType.Provider

		for _, c := range conf.Chats {
			p, err := chat.GetProvider(c.Provider)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Client] Failed to get Chat provider %s: %q", c.Provider, err))
				continue
			}
			providerList = append(providerList, p)
		}
		provider, err := SelectChatForm(ctx, providerList, "Please select a provider:")
		if err != nil {
			return err
		}

		cfg := provider.Config()
		var confData any
		var outputMsg strings.Builder
		for i, cp := range conf.Chats {
			if cp.Provider == provider.Name() {
				// get config
				confData = cp.Configuration
				err = mapstructure.Decode(cp.Configuration, &cfg)
				if err != nil {
					outputMsg.WriteString(fmt.Sprintf("%d) Could not get this!\n", i+1))
					break
				}
				outputMsg.WriteString(fmt.Sprintf("%d) `%s`, config: %s\n", i+1, provider.DisplayName(), cfg))
				break
			}
		}
		action, err := configPageActionSelectForm("Edit This Chat Provider", "Delete This Provider", "Quit")
		if err != nil {
			return err
		}
		switch action {
		case configLandingPageEdit:
			edited, err := RunEditConfig(ctx, cfg, confData)
			if err != nil {
				return err
			}
			for i, c := range conf.Chats {
				if c.Provider == provider.Name() {
					conf.Chats[i].Configuration = edited
					break
				}
			}
			err = config.Write(conf)
			if err != nil {
				return err
			}
			color.Blue("Config Updated!")
			time.Sleep(1 * time.Second)
		case configLandingPageDelete:
			for i, c := range conf.Chats {
				if c.Provider == provider.Name() {
					conf.Chats = append(conf.Chats[:i], conf.Chats[i+1:]...)
					break
				}
			}
			err = config.Write(conf)
			if err != nil {
				return err
			}
			color.Blue("Config Deleted!")
			time.Sleep(1 * time.Second)
		case configLandingPageQuit:
			return nil
		}
	case configLandingActionQuit:
		color.Blue("Bye!")
		return nil
	default:
		return fmt.Errorf("unexpected landing chat-action form: %s", se)
	}
	return nil
}
