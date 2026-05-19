package interactive

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	_ "github.com/manboster/manboster/internal/chat/all"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/hachimi"
	_ "github.com/manboster/manboster/internal/hachimi/all"
	"github.com/manboster/manboster/internal/llm"
	_ "github.com/manboster/manboster/internal/llm/all"
	"github.com/manboster/manboster/internal/tool"
	_ "github.com/manboster/manboster/internal/tool/all"
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
		if !helper.ContinueConfirm(ctx, fmt.Sprintf("You've successfully added %d llm providers!", count)) {
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

	// Step 5: config hachimi
	hachimiCfg, err := OnboardHachimiConfigForm(ctx)
	if err != nil {
		return c, err
	}
	c.Hachimi = hachimiCfg

	// set V and manboster.db path
	c = c.Default()

	// Step 6: See what's entered and start to write configuration.
	confDescription := strings.Builder{}
	confDescription.WriteString("# Before you proceed, you need to review what you have entered. \n")
	confDescription.WriteString("If anything is incorrect, please use Ctrl+C to quit and restart it with 'manboster onboard'.\n\n")

	confDescription.WriteString(fmt.Sprintf("You configured %d chat providers\n\n", len(c.Chats)))
	for i, _ := range c.Chats {
		confDescription.WriteString(fmt.Sprintf("#%d: %s's Configuration:\n\n %s\n\n", i+1, c.Chats[i].Provider, c.Chats[i].Configuration))
	}

	confDescription.WriteString(fmt.Sprintf("You configured %d llm providers\n\n", len(c.LLMs)))
	for i, _ := range c.LLMs {
		confDescription.WriteString(fmt.Sprintf("#%d's Configuration: \n\n%s \n\n", i+1, c.LLMs[i].Configuration))
	}

	confDescription.WriteString(fmt.Sprintf("You configured %d tool providers\n\n", len(c.Tools)))
	for i, _ := range c.Tools {
		confDescription.WriteString(fmt.Sprintf("#%d: %s's Configuration: \n\n%s \n\n", i+1, c.Tools[i].Name, c.Tools[i].Configuration))
	}

	if c.Hachimi.Enabled {
		confDescription.WriteString(fmt.Sprintf("You enabled hachimi features\n\n"))
	} else {
		confDescription.WriteString(fmt.Sprintf("You disabled hachimi feature.\n\n"))
	}

	confDescription.WriteString("If there is no problem, you can continue writing the configuration.\n\n")
	confDesc := confDescription.String()

	err = helper.ConfirmForm(ctx, confDesc, "Do you want to continue?", "Continue")
	if err != nil {
		return c, err
	}

	return c, nil
}

// OnboardChatConfigForm returns chats' config data.
func OnboardChatConfigForm(ctx context.Context) (config.ChatConfig, error) {
	chatProviders := chat.AllProviders()

	chatProvider, err := SelectChatForm(ctx, chatProviders, "First, which platform would you like to use for your Manboster?")
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
	providers, err := SelectToolForm(ctx, tools, "Select and activate the tools you want to use", "Please select the tool you want to use, please be careful to select as they will be the tool call of AIs.")
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
	llmProvider, err := SelectLLMForm(ctx, llmProviders, "Next, let's pick a LLM provider. Which provider would you like to use?")
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
	provider, err := SelectLLMProviderInstanceForm(ctx, llmConfig, "Please select the default provider you want to use in the Manboster:", "The model you select will be the default model of all sessions. If you don't know what's this, please leave it as is.")
	if err != nil {
		return config.AppConfig{}, err
	}
	models := provider.Models()
	model, err := SelectModelForm(ctx, models, "Please select the default model you want to use in the manboster:", "The model you select will be the default model of all sessions. If you don't know what's this, please leave it as is.")
	if err != nil {
		return config.AppConfig{}, err
	}

	return config.AppConfig{
		DefaultLLMProvider: provider.Name(),
		DefaultLLMModel:    model.Name,
	}, nil
}

func OnboardHachimiConfigForm(ctx context.Context) (config.HachimiConfigs, error) {
	var hachimiConf config.HachimiConfigs
	isEnabled, err := HachimiEnablePrompt(ctx, "Do you want to activate hachimi?", "Hachimi is a small language model running on your device side and it can check out LLM's behaviour and evaulate its action.\nIf your device's available memory is lower than 1GB or you don't know what's this, please disable it.")
	if err != nil {
		return hachimiConf, err
	}
	if !isEnabled {
		hachimiConf.Enabled = false
		return hachimiConf, nil
	}
	hachimiProviders := hachimi.AllProviders()
	hachimiProvider, err := SelectHachimiForm(ctx, hachimiProviders, "Please choose your hachimi provider:")
	if err != nil {
		return hachimiConf, err
	}
	provider := hachimiProvider.Config()
	conf, err := RunOnboardConfig(ctx, provider)
	if err != nil {
		return hachimiConf, err
	}
	hachimiConf.Enabled = true
	hachimiConf.Provider = provider.Name()
	hachimiConf.Hachimi = append(hachimiConf.Hachimi, config.HachimiConfig{
		Provider:      provider.Name(),
		Configuration: conf,
	})
	return hachimiConf, nil
}
