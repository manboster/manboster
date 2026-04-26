package cli

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/spec/chat"
)

func OnboardSelectLLMForm(ctx context.Context, llmProviders []llm.Provider, prompt string) (llm.Provider, error) {
	var options []huh.Option[string]
	for _, c := range llmProviders {
		options = append(options, huh.NewOption(c.Config().DisplayName(), c.Config().Name()))
	}
	var lProvider string
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title(prompt).
				Options(
					options...,
				).Value(&lProvider),
		),
	).Run()
	if err != nil || lProvider == "" {
		return nil, err
	}
	for _, c := range llmProviders {
		if c.Config().Name() == lProvider {
			return c, nil
		}
	}
	return nil, fmt.Errorf("no such provider %s", lProvider)
}

func OnboardLLMProviderInstanceForm(ctx context.Context, llmConfigs []config.LLMConfig, prompt string) (llm.Provider, error) {
	var llmProviders []llm.Provider
	for _, c := range llmConfigs {
		p, err := llm.GetProvider(c.Provider)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Configuration Wizard] Failed to get llm provider %q: %q\n", c.Provider, err))
			continue
		}
		err = p.Init(ctx, c.Configuration)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Configuration Wizard] Failed to init llm provider %q: %q\n", c.Provider, err))
			continue
		}
		llmProviders = append(llmProviders, p)
	}
	var llmOptions []huh.Option[string]
	for _, c := range llmProviders {
		llmOptions = append(llmOptions, huh.NewOption(c.DisplayName(), c.Name()))
	}
	var lProvider string
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title(prompt).
				Options(
					llmOptions...,
				).Value(&lProvider),
		),
	).Run()
	if err != nil || lProvider == "" {
		return nil, err
	}
	for _, c := range llmProviders {
		if c.Name() == lProvider {
			return c, nil
		}
	}
	return nil, fmt.Errorf("no such provider %s", lProvider)
}

func OnboardSelectChatForm(ctx context.Context, chatProviders []chat.Provider, prompt string) (chat.Provider, error) {
	var chatOptions []huh.Option[string]
	for _, c := range chatProviders {
		chatOptions = append(chatOptions, huh.NewOption(c.DisplayName(), c.Name()))
	}

	var chatProvider string
	chatProviderForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title(prompt).
				Options(
					chatOptions...,
				).Value(&chatProvider),
		),
	)
	err := chatProviderForm.Run()
	if err != nil || chatProvider == "" {
		return nil, err
	}
	for _, c := range chatProviders {
		if c.Name() == chatProvider {
			return c, nil
		}
	}
	return nil, fmt.Errorf("no such provider %s", chatProvider)
}

func OnboardSelectModelForm(ctx context.Context, models []llm.Model, prompt string) (llm.Model, error) {
	var options []huh.Option[string]
	for i, modelData := range models {
		option := huh.NewOption(modelData.DisplayName, modelData.Name)
		if i == 0 {
			option = option.Selected(true)
		}
		options = append(options, option)
	}

	var model string
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select the default model").
				Description("The model you select will be the default model of all sessions. If you don't know what's this, please leave it as is.").
				Options(options...).Value(&model),
		),
	).Run()
	if err != nil || model == "" {
		return llm.Model{}, err
	}
	for _, m := range models {
		if m.Name == model {
			return m, nil
		}
	}
	return llm.Model{}, fmt.Errorf("no such model %s", model)
}
