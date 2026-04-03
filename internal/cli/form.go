package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/util"

	_ "github.com/manboster/manboster/internal/chat/telegram"
	_ "github.com/manboster/manboster/internal/llm/oai_compat"
	_ "github.com/manboster/manboster/internal/llm/openrouter"
)

// Form provides a huh form configuration with TUI.
func Form() (config.Config, error) {
	// get default configuration's value
	var c config.Config
	c = config.Default(c)

	err := huh.NewNote().
		Title("Manboster Configuration Wizard").
		Description("Welcome to the Manboster Configuration Wizard. Enjoy your experience with your little Manbo!").
		Next(true).Run()
	if err != nil {
		return c, err
	}

	// TODO: Refactor to single functions
	// Step 1: choose Chat Providers
	chatCfg, err := ChatConfigForm()
	if err != nil {
		return c, err
	}
	c.Chats = append(c.Chats, chatCfg)

	// Step 2: config LLMs
	llmCfg, err := LLMConfigForm()
	c.LLMs = append(c.LLMs, llmCfg)

	// Step 3: See what's entered and start to write configuration.
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
func ChatConfigForm() (config.ChatConfig, error) {
	// get providers to generate options
	var chatProviders []config.Provider
	names := config.AvailProviders("chat")
	for _, name := range names {
		provider, err := config.GetProvider(name)
		if err != nil {
			return config.ChatConfig{}, err
		}
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

	provider, err := config.GetProvider("chat:" + chatProvider)
	if err != nil {
		return config.ChatConfig{}, err
	}

	err = huh.NewForm(provider.ToHuhGroup()...).Run()
	if err != nil {
		return config.ChatConfig{}, err
	}

	err = provider.VerifyAndConvert()
	if err != nil {
		return config.ChatConfig{}, err
	}

	return config.ChatConfig{
		Configuration: provider.GetConfig(),
		Provider:      provider.Name(),
	}, nil
}

// LLMConfigForm returns LLMs' config data.
func LLMConfigForm() (config.LLMConfig, error) {
	// get providers to generate options
	var llmProviders []config.Provider
	names := config.AvailProviders("llm")
	for _, name := range names {
		provider, err := config.GetProvider(name)
		if err != nil {
			return config.LLMConfig{}, err
		}
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

	provider, err := config.GetProvider("llm:" + lProvider)
	if err != nil {
		return config.LLMConfig{}, err
	}

	err = huh.NewForm(provider.ToHuhGroup()...).Run()
	if err != nil {
		return config.LLMConfig{}, err
	}

	err = provider.VerifyAndConvert()
	if err != nil {
		return config.LLMConfig{}, err
	}

	return config.LLMConfig{
		Configuration: provider.GetConfig(),
		Provider:      provider.Name(),
	}, nil
}
