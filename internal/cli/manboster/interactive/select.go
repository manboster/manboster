package interactive

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/chat"
	llmType "github.com/manboster/manboster/spec/llm"
)

func SelectLLMProviderInstanceForm(ctx context.Context, llmConfigs []config.LLMConfig, title string, prompt string) (llmType.Provider, error) {
	var llmProviders []llmType.Provider
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
			huh.NewSelect[string]().Title(title).Description(prompt).
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

func SelectModelForm(ctx context.Context, models []llmType.Model, title string, prompt string) (llmType.Model, error) {
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
				Title(title).
				Description(prompt).
				Options(options...).Value(&model),
		),
	).Run()
	if err != nil || model == "" {
		return llmType.Model{}, err
	}
	for _, m := range models {
		if m.Name == model {
			return m, nil
		}
	}
	return llmType.Model{}, fmt.Errorf("no such model %s", model)
}

func SelectLLMForm(ctx context.Context, llmProviders []llmType.Provider, prompt string) (llmType.Provider, error) {
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

func SelectChatForm(ctx context.Context, chatProviders []chat.Provider, prompt string) (chat.Provider, error) {
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

func SelectToolForm(ctx context.Context, toolProviders []tool.Provider, title string, prompt string) ([]tool.Provider, error) {
	var toolOptions []huh.Option[string]
	toolsMap := map[string]tool.Provider{}

	for _, toolData := range toolProviders {
		toolsMap[toolData.Name()] = toolData

		display := fmt.Sprintf("%s(%s)\n%s", toolData.DisplayName(), toolData.Name(), toolData.MetaData().Description)
		// to check the compatibility...

		option := huh.NewOption(display, toolData.Name())
		toolOptions = append(toolOptions, option)
	}
	var tools []string
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().Title(title).Description(prompt).
				Options(toolOptions...).Value(&tools),
		)).Run()
	if err != nil || tools == nil {
		return nil, err
	}

	var respTool []tool.Provider
	for _, toolName := range tools {
		if t, ok := toolsMap[toolName]; ok {
			respTool = append(respTool, t)
		}
	}

	return respTool, nil
}
