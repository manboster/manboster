package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/util"

	_ "github.com/manboster/manboster/internal/chat/telegram"
	_ "github.com/manboster/manboster/internal/llm/oai_compat"
	_ "github.com/manboster/manboster/internal/llm/openrouter"
)

// ConfigurationForm provides a huh form configuration with TUI.
func ConfigurationForm(ctx context.Context) (config.Config, error) {
	// get default configuration's value
	var c config.Config

	err := huh.NewNote().
		Title("Manboster Configuration Wizard").
		Description("Welcome to the Manboster Configuration Wizard. Enjoy your experience with your little Manbo!").
		Next(true).Run()
	if err != nil {
		return c, err
	}

	// TODO: Refactor to single functions
	// Step 1: choose Chat Providers
	chatCfg, err := ChatConfigForm(ctx)
	if err != nil {
		return c, err
	}
	c.Chats = append(c.Chats, chatCfg)

	// Step 2: config LLMs
	llmCfg, err := LLMConfigForm(ctx)
	if err != nil {
		return c, err
	}
	c.LLMs = append(c.LLMs, llmCfg)

	// Step 3: config apps
	appCfg, err := APPConfigForm(ctx, llmCfg)
	if err != nil {
		return c, err
	}
	c.App = appCfg

	// set V and manboster.db path
	c = config.Default(c)

	// Step 4: See what's entered and start to write configuration.
	confDescription := strings.Builder{}
	confDescription.WriteString("You need to review what you have entered. \n")
	confDescription.WriteString("If anything is incorrect, please use Ctrl+C to quit and restart it with 'manboster config'.\n")
	confDescription.WriteString(fmt.Sprintf("Your Chat Provider: %s\n", c.Chats[0].Provider))
	confDescription.WriteString(fmt.Sprintf("Your Chat Provider's Configuration:\n %s\n", c.Chats[0].Configuration))
	confDescription.WriteString(fmt.Sprintf("Your LLM Provider: %s\n", c.LLMs[0].Provider))
	confDescription.WriteString(fmt.Sprintf("Your LLM Provider's Configuration:\n %s\n", c.LLMs[0].Configuration))
	confDescription.WriteString("If there is no problem, you can press enter and we will work on it.\n")
	confDesc := util.EscapeMarkdown(confDescription.String())

	err = huh.NewNote().
		Title("Before you proceed...").
		Description(confDesc).
		Next(true).Run()
	if err != nil {
		return c, err
	}

	return c, nil
}

// ChatConfigForm returns chats' config data.
func ChatConfigForm(ctx context.Context) (config.ChatConfig, error) {
	// get providers to generate options
	var chatProviders []config.Provider
	names := chat.AvailProviders()
	for _, name := range names {
		chatProvider, err := chat.GetProvider(name)
		if err != nil {
			return config.ChatConfig{}, err
		}
		provider := chatProvider.Config()
		chatProviders = append(chatProviders, provider)
	}

	var chatOptions []huh.Option[string]
	for _, c := range chatProviders {
		chatOptions = append(chatOptions, huh.NewOption(c.DisplayName(), c.Name()))
	}

	var chatProvider string
	chatProviderForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("First, which platform would you like to use for your Manboster?").
				Options(
					chatOptions...,
				).Value(&chatProvider),
		),
	)
	err := chatProviderForm.Run()
	if err != nil {
		return config.ChatConfig{}, err
	}

	cProvider, err := chat.GetProvider(chatProvider)
	if err != nil {
		return config.ChatConfig{}, err
	}

	provider := cProvider.Config()
	conf, err := RunConfig(ctx, provider)
	if err != nil {
		return config.ChatConfig{}, err
	}

	return config.ChatConfig{
		Configuration: conf,
		Provider:      provider.Name(),
	}, nil
}

// LLMConfigForm returns LLMs' config data.
func LLMConfigForm(ctx context.Context) (config.LLMConfig, error) {
	// get providers to generate options
	var llmProviders []config.Provider
	names := llm.AvailProviders()
	for _, name := range names {
		lProvider, err := llm.GetProvider(name)
		if err != nil {
			return config.LLMConfig{}, err
		}
		provider := lProvider.Config()
		llmProviders = append(llmProviders, provider)
	}

	var chatOptions []huh.Option[string]
	for _, c := range llmProviders {
		chatOptions = append(chatOptions, huh.NewOption(c.DisplayName(), c.Name()))
	}

	var lProvider string
	lProviderForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Next, let's pick an LLM. Which provider would you like to use?").
				Options(
					chatOptions...,
				).Value(&lProvider),
		),
	)
	err := lProviderForm.Run()
	if err != nil {
		return config.LLMConfig{}, err
	}

	llmProvider, err := llm.GetProvider(lProvider)
	if err != nil {
		return config.LLMConfig{}, err
	}

	provider := llmProvider.Config()
	conf, err := RunConfig(ctx, provider)
	if err != nil {
		return config.LLMConfig{}, err
	}

	return config.LLMConfig{
		Configuration: conf,
		Provider:      provider.Name(),
	}, nil
}

func APPConfigForm(ctx context.Context, llmConfig config.LLMConfig) (config.AppConfig, error) {
	provider, err := llm.GetProvider(llmConfig.Provider)
	if err != nil {
		return config.AppConfig{}, err
	}

	// initialize provider's models immediately
	err = provider.Init(ctx, llmConfig.Configuration)
	if err != nil {
		return config.AppConfig{}, err
	}
	models := provider.Models()
	var options []huh.Option[string]
	for i, modelData := range models {
		option := huh.NewOption(modelData.DisplayName, modelData.Name)
		if i == 0 {
			option = option.Selected(true)
		}
		options = append(options, option)
	}

	var model string
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select the default model").
				Description("The model you select will be the default model of all sessions. If you don't know what's this, please leave it as is.").
				Options(options...).Value(&model),
		),
	).Run()
	if err != nil {
		return config.AppConfig{}, err
	}

	return config.AppConfig{
		DefaultLLMProvider: provider.Name(),
		DefaultLLMModel:    model,
	}, nil
}
