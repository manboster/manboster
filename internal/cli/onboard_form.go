package cli

import (
	"context"
	"fmt"
	"os"
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

// OnboardWarningForm provides a warning notice
func OnboardWarningForm(ctx context.Context) error {
	agree := false

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("RISK DISCLOSURE & DISCLAIMER").
				Description(`*PLEASE READ THESE WORDS CAREFULLY:*
Manboster is an AI agent able to chat and control your computers like OpenClaw and IronClaw and currently in MVP stage. By proceeding, you acknowledge:
1. WIP means this project is *Work in Progress*, and *it is expected to encounter bugs, crashes, and breaking changes.*
2. If you run 'manboster start', you open a daemon running in your computer. *The background process has persistent resource access to your computer.*
3. WASM sandboxing plugins is strong, but *3rd-party code still carries risks*.
4. *Hachimi scoring reduces decision fatigue, but cannot fully prevent advanced prompt injections or unsafe LLM behaviors.*
5. *Granting access enables data transmission to LLMs and allows device control. We are not liable for any issues arising from these interactions.*
6. This software is provided "AS IS" under Apache 2.0. *You are strictly prohibited from using this application for any criminal or illegal purposes. We disclaim all liability and responsibility for any unlawful activities conducted using this software.*
`).
				Next(false),
			huh.NewConfirm().
				Title("Do you understand the risks and wish to proceed?").
				Affirmative("I Agree & Continue").
				Negative("Exit Now").
				Value(&agree),
		),
	)

	err := form.Run()
	if !agree {
		os.Exit(0)
	}
	return err
}

func OnboardVersionWarningForm(ctx context.Context) error {
	agree := false

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("UNSTABLE VERSION WARNING").
				Description(`*PLEASE READ THESE WORDS CAREFULLY:*
It seems that you're going to use an unstable version of Manboster. Please note that:
1. It's normal to encounter bugs, crashes, and breaking changes in unstable versions.
2. As this is not a stable version, it's not contain ANY security patches and fixes.
3. This version's configuration may be incompatible with older versions and please aware the configuration changes.
4. If you encounter bugs, we appreciate you to commit to issues and we will fix it as soon as possible.
5. PLEASE DO NOT STORAGE ANY SENSITIVE AND IMPORTANT DATA IN THIS VERSION! As it's unstable and we are unsure that this application will work as is.
`).
				Next(false),
			huh.NewConfirm().
				Title("Do you understand the risks and wish to proceed?").
				Affirmative("I Understand & Continue").
				Negative("Exit Now").
				Value(&agree),
		),
	)

	err := form.Run()
	if !agree {
		os.Exit(0)
	}
	return err
}

func ContinueConfirm(ctx context.Context, content string) bool {
	agree := false

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf("%s\nContinue?", content)).
				Affirmative("Continue").
				Negative("Skip").
				Value(&agree),
		),
	)

	err := form.Run()
	if err != nil {
		os.Exit(0)
	}
	return agree
}

// OnboardConfigurationForm provides a huh form configuration with TUI.
func OnboardConfigurationForm(ctx context.Context) (config.Config, error) {
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
	chatCfg, err := OnboardChatConfigForm(ctx)
	if err != nil {
		return c, err
	}
	c.Chats = append(c.Chats, chatCfg)

	count := 1
	for {
		// Step 2: config LLMs
		llmCfg, err := OnboardLLMConfigForm(ctx)
		if err != nil {
			return c, err
		}
		c.LLMs = append(c.LLMs, llmCfg)
		if !ContinueConfirm(ctx, fmt.Sprintf("You've successfully add %d llm configs!", count)) {
			break
		}
		count++
	}

	// Step 3: config apps
	appCfg, err := OnboardAPPConfigForm(ctx, c.LLMs)
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

// OnboardChatConfigForm returns chats' config data.
func OnboardChatConfigForm(ctx context.Context) (config.ChatConfig, error) {
	chatProviders := chat.AllProviders()

	chatProvider, err := OnboardSelectChatForm(ctx, chatProviders, "First, which platform would you like to use for your Manboster?")
	if err != nil {
		return config.ChatConfig{}, err
	}

	provider := chatProvider.Config()
	conf, err := RunOnboardConfig(ctx, provider)
	if err != nil {
		return config.ChatConfig{}, err
	}

	return config.ChatConfig{
		Configuration: conf,
		Provider:      provider.Name(),
	}, nil
}

// OnboardLLMConfigForm returns LLMs' config data.
func OnboardLLMConfigForm(ctx context.Context) (config.LLMConfig, error) {
	// get providers to generate options
	llmProviders := llm.AllProviders()
	llmProvider, err := OnboardSelectLLMForm(ctx, llmProviders, "Next, let's pick an LLM. Which provider would you like to use?")
	if err != nil {
		return config.LLMConfig{}, err
	}

	provider := llmProvider.Config()
	conf, err := RunOnboardConfig(ctx, provider)
	if err != nil {
		return config.LLMConfig{}, err
	}

	err = llmProvider.Init(ctx, conf)
	if err != nil {
		return config.LLMConfig{}, err
	}

	return config.LLMConfig{
		Configuration: conf,
		Provider:      provider.Name(),
	}, nil
}

func OnboardAPPConfigForm(ctx context.Context, llmConfig []config.LLMConfig) (config.AppConfig, error) {
	provider, err := OnboardLLMProviderInstanceForm(ctx, llmConfig, "Please select the default provider you want to use in the Manboster:")
	if err != nil {
		return config.AppConfig{}, err
	}
	models := provider.Models()
	model, err := OnboardSelectModelForm(ctx, models, "Please select the default model you want to use inthe manboster:")
	if err != nil {
		return config.AppConfig{}, err
	}

	return config.AppConfig{
		DefaultLLMProvider: provider.Name(),
		DefaultLLMModel:    model.Name,
	}, nil
}
