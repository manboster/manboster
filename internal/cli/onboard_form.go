package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"

	_ "github.com/manboster/manboster/internal/chat/telegram"
	_ "github.com/manboster/manboster/internal/llm/oai_compat"
	_ "github.com/manboster/manboster/internal/llm/openrouter"
	_ "github.com/manboster/manboster/internal/tool/datetime"
	_ "github.com/manboster/manboster/internal/tool/memory"
)

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
		if !ContinueConfirm(ctx, fmt.Sprintf("You've successfully add %d llm providers!", count)) {
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

	// Step 4: config tools(first is tool call)
	toolProvidersName := tool.AvailProviders()
	var toolProviders []tool.Provider
	for _, provider := range toolProvidersName {
		p, err := tool.GetProvider(provider)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Client] We encountered an error while loading tool call provider %q: %q", provider, err))
		}
		toolProviders = append(toolProviders, p)
	}
	toolCfg, err := OnboardToolConfigForm(ctx, toolProviders)
	if err != nil {
		return c, err
	}
	c.Tools = toolCfg

	// set V and manboster.db path
	c = config.Default(c)

	// Step 4: See what's entered and start to write configuration.
	confDescription := strings.Builder{}
	confDescription.WriteString("You need to review what you have entered. \n")
	confDescription.WriteString("If anything is incorrect, please use Ctrl+C to quit and restart it with 'manboster config'.\n")

	confDescription.WriteString(fmt.Sprintf("You configured %d chat providers\n", len(c.Chats)))
	for i, _ := range c.Chats {
		confDescription.WriteString(fmt.Sprintf("#%d: %s's Configuration:\n %s\n", i+1, c.Chats[i].Provider, c.Chats[i].Configuration))
	}

	confDescription.WriteString(fmt.Sprintf("You configured %d llm providers\n", len(c.LLMs)))
	for i, _ := range c.LLMs {
		confDescription.WriteString(fmt.Sprintf("#%d's Configuration: \n%s \n", i+1, c.LLMs[i].Configuration))
	}

	confDescription.WriteString(fmt.Sprintf("You configured %d tool providers\n", len(c.Tools)))
	for i, _ := range c.Tools {
		confDescription.WriteString(fmt.Sprintf("#%d: %s's Configuration: \n%s \n", i+1, c.Tools[i].Name, c.Tools[i].Configuration))
	}

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

func OnboardToolConfigForm(ctx context.Context, tools []tool.Provider) ([]config.ToolConfig, error) {
	providers, err := OnboardSelectToolForm(ctx, tools, "Select and activate the tools you want to use", "Please select the tool you want to use, please be careful to select as they will be the tool call of AIs.")
	if err != nil {
		return nil, err
	}

	var toolConfigs []config.ToolConfig
	for _, provider := range providers {
		var conf any
		if provider.Config() != nil {
			conf, err = RunOnboardConfig(ctx, provider.Config())
			if err != nil {
				return nil, err
			}
		}
		toolConfigs = append(toolConfigs, config.ToolConfig{
			Name:          provider.Name(),
			Configuration: conf,
		})
	}

	return toolConfigs, nil
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
	model, err := OnboardSelectModelForm(ctx, models, "Please select the default model you want to use in the manboster:")
	if err != nil {
		return config.AppConfig{}, err
	}

	return config.AppConfig{
		DefaultLLMProvider: provider.Name(),
		DefaultLLMModel:    model.Name,
	}, nil
}
