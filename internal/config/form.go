package config

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/internal/chat/telegram"
	"github.com/manboster/manboster/internal/llm/oai_compat"
	"github.com/manboster/manboster/internal/llm/openrouter"
	"github.com/manboster/manboster/internal/util"
)

// Form provides a huh form configuration with TUI.
func Form() (Config, error) {
	var c Config

	err := huh.NewNote().
		Title("Manboster Configuration Wizard").
		Description("Welcome to the Manboster Configuration Wizard. Enjoy your experience with your little Manbo!").
		Next(true).Run()
	if err != nil {
		return c, err
	}

	// TODO: Refactor to single functions
	// Step 1: choose Chat Providers
	var chatProvider string
	chatProviderForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("First, which platform would you like to use for your Manboster?").
				Options(
					huh.NewOption("Telegram", "telegram"),
				).Value(&chatProvider),
		),
	)
	err = chatProviderForm.Run()
	if err != nil {
		return c, err
	}
	switch chatProvider {
	// Telegram specific configuration, you can change this in chat/providers/telegram/config.go.
	case "telegram":
		tgProvider := &telegram.Config{}

		err = huh.NewForm(tgProvider.ToHuhGroup()...).Run()
		if err != nil {
			return c, err
		}
		// convert string to int64(uid)
		err = tgProvider.VerifyAndConvert()
		if err != nil {
			return c, err
		}

		c.Chats = append(c.Chats, ChatConfig{
			Provider:      "telegram",
			Configuration: tgProvider.GetConfig(),
		})
	}

	// TODO: Refactor to single function
	var llmProvider string
	// Step 2: choose LLM Providers
	llmProviderForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Next, let's pick an LLM. Which provider would you like to use?").
				Options(
					huh.NewOption("OpenRouter", "openrouter"),
					huh.NewOption("OpenAI-compatible APIs", "oai-compat"),
				).
				Value(&llmProvider),
		),
	)
	err = llmProviderForm.Run()
	if err != nil {
		return c, err
	}

	// TODO: Refactor to single function
	switch llmProvider {
	// openrouter specific configuration
	case "openrouter":
		orProvider := &openrouter.Config{}
		err = huh.NewForm(orProvider.ToHuhGroup()...).Run()
		if err != nil {
			return c, err
		}

		err = orProvider.VerifyAndConvert()
		if err != nil {
			return c, err
		}

		c.LLMs = append(c.LLMs, LLMConfig{
			Provider:      "openrouter",
			Configuration: orProvider.GetConfig(),
		})
	// oai config
	case "oai-compat":
		oaiProvider := &oai_compat.Config{}
		err = huh.NewForm(oaiProvider.ToHuhGroup()...).Run()
		if err != nil {
			return c, err
		}

		c.LLMs = append(c.LLMs, LLMConfig{
			Provider:      "oai-compat",
			Configuration: oaiProvider.GetConfig(),
		})
	}

	// Step 3: See what's entered and start to write configuration.
	confDesc := fmt.Sprintf(`
You need to review what you have entered. 
If anything is incorrect, please use Ctrl+C to quit and restart it with 'manboster config'.
Your Chat Provider: %s
Your Chat Provider's Configuration: %s
Your LLM Provider: %s
Your LLM Provider's Configuration: %s
If there is no problem, you can press enter and we will work on it.
`, c.Chats[0].Provider, c.Chats[0].Configuration, c.LLMs[0].Provider, c.LLMs[0].Configuration)
	confDesc = util.EscapeMarkdown(confDesc)

	err = huh.NewNote().
		Title("Before you proceed...").
		Description(confDesc).
		Next(true).Run()
	if err != nil {
		return c, err
	}

	return c, nil
}
