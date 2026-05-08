package interactive

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/database"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/tool"
	chatType "github.com/manboster/manboster/spec/chat"
	llmType "github.com/manboster/manboster/spec/llm"
)

func configStartupForm() (configSelection, error) {
	var s configSelection
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[configSelection]().Options(
				huh.NewOption("Database\nThis will open your database and manage chats, users and sessions. If you want to manage chat sessions, purge unused sessions or more, please choose this.", configSelectionDatabase),
				huh.NewOption("Configuration\nThis will affect your model and chat provider's configuration and it displays the changes in config.yaml. If you want to manage models, default models, or modify application settings, please choose this.", configSelectionConfig),
				huh.NewOption("Open Configuration yaml file in system's default editor\n(For advanced users only)", configSelectionEditor),
				huh.NewOption("Quit Manboster Configuration Wizard\nBye!", configSelectionQuit),
			).Title("Please select what to configure").Description("Welcome to Manboster Configuration Wizard! Please choose which field you want to configure.").Value(&s),
		)).Run()
	if err != nil {
		return "", err
	}
	return s, nil
}

func configFormRun() error {
	for {
		se, err := configStartupForm()
		if err != nil {
			return err
		}
		err = config.Init()
		if err != nil {
			color.Red("It seems that there is no configuration available in your device, please run 'manboster onboard' first!")
			return err
		}
		switch se {
		case configSelectionEditor:
			configCmdOpenRun(nil, nil)
			color.Blue("Opened via your default Editor, please edit in the editor, save it and then restart the instance!")
			os.Exit(0)
		case configSelectionConfig:
			err := runConfigLandingSelectionForm()
			if err != nil {
				return err
			}
		case configSelectionDatabase:
			cli := database.Client{}
			path := config.Read().App.DBPath
			err := cli.Init(path)
			if err != nil {
				color.Red("It seems that there is no database available in your device, please run 'manboster' first!")
				return err
			}

			repo := repository.New(cli.Instance())
			s := newDatabaseConfigService(repo)
			err = s.configDatabaseLandingForm()
			if err != nil {
				return err
			}
		case configSelectionQuit:
			color.Blue("Bye!")
			os.Exit(0)
		}
	}

	return nil
}

func (s *databaseConfigService) runConfigDatabaseLandingSelectionForm() (databaseConfigLandingSelection, error) {
	var se databaseConfigLandingSelection
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[databaseConfigLandingSelection]().Options(
				huh.NewOption("Sessions\nIt's chat sessions you made in daily chatting routine, you can purge it to save disk or do more.", databaseConfigLandingSession),
				huh.NewOption("Users\nIt helps you to grant a user to an admin or degrade it. It's all up to you to decide.", databaseConfigLandingUser),
				huh.NewOption("Souls\nSouls is the key system prompt for you to personalize your Manbo Lobster.", databaseConfigLandingSoul),
				huh.NewOption("Quit\nBye!", databaseConfigLandingQuit),
			).Title("Please select what to configure in database").Description("Please choose what you want to configure in database field.").Value(&se),
		)).Run()
	if err != nil {
		return "", err
	}
	return se, nil
}

func (s *databaseConfigService) configDatabaseLandingForm() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		se, err := s.runConfigDatabaseLandingSelectionForm()
		if err != nil {
			return err
		}
		switch se {
		case databaseConfigLandingUser:
			return nil
		case databaseConfigLandingSession:
			err := s.runConfigDatabaseSessionSelection(ctx)
			if err != nil {
				return err
			}
			return nil
		case databaseConfigLandingSoul:
			return nil
		case databaseConfigLandingQuit:
			color.Blue("Bye!")
			return nil
		default:
			return fmt.Errorf("unexpected database landing form: %s", se)
		}
	}
}

func configLandingForm() (configLandingSelection, error) {
	var se configLandingSelection
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[configLandingSelection]().Options(
				huh.NewOption("Chat Providers\nAdd, edit or delete your chat providers.", configLandingChat),
				huh.NewOption("LLM Providers\nAdd, edit or delete your llm providers.", configLandingLLM),
				huh.NewOption("Tool Providers\nAdd, edit or delete your system tool providers.", configLandingTool),
				huh.NewOption("Hachimi Settings\nAdd, edit or delete your Hachimi providers or modify Hachimi settings.", configLandingHachimi),
				huh.NewOption("App Settings\nModify Manboster settings.", configLandingApp),
				huh.NewOption("Quit\nBye!", configLandingQuit),
			).Value(&se).Title("Please select what to configure in configuration").Description("Please choose what you want to configure in configuration field."),
		)).Run()
	if err != nil {
		return se, err
	}
	return se, nil
}

func runConfigLandingSelectionForm() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer helper.ClearScreen()

	conf := config.Read()

	se, err := configLandingForm()
	if err != nil {
		return err
	}
	switch se {
	case configLandingChat:
		err := runLandingChatActionForm(ctx)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
		return nil
	case configLandingLLM:
		err := runLandingLLMActionForm(ctx)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
		return nil
	case configLandingTool:
		err := runLandingToolActionForm(ctx)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
		return nil
	case configLandingHachimi:
	case configLandingApp:
		appConf, err := OnboardAPPConfigForm(ctx, conf.LLMs)
		if err != nil {
			return err
		}
		dbpath := conf.App.DBPath
		conf.App = appConf
		conf.App.DBPath = dbpath
		err = config.Write(conf)
		if err != nil {
			return err
		}
		color.Blue("Successfully saved config!")
		time.Sleep(1 * time.Second)
	case configLandingQuit:
		color.Blue("Bye!")
		return nil
	default:
		return fmt.Errorf("unexpected database landing form: %s", se)
	}
	return nil
}

func configLandingChatActionForm() (configLandingActionSelection, error) {
	return configLandingActionForm("Add a new chat provider", "Select an existing chat provider", "Quit")
}

func configLandingLLMActionForm() (configLandingActionSelection, error) {
	return configLandingActionForm("Add a new LLM provider", "Select an existing LLM provider", "Quit")
}

func configLandingToolActionForm() (configLandingActionSelection, error) {
	return configLandingActionForm("Add a new Tool provider", "Select an existing tool provider", "Quit")
}

func configLandingHachimiActionForm() (configLandingActionSelection, error) {
	return configLandingActionForm("Add a new hachimi provider", "Select an existing hachimi provider", "Quit")
}

func configLandingActionForm(add string, sel string, quit string) (configLandingActionSelection, error) {
	var se configLandingActionSelection
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[configLandingActionSelection]().Options(
				huh.NewOption(add, configLandingActionAdd),
				huh.NewOption(sel, configLandingActionSelect),
				huh.NewOption(quit, configLandingActionQuit),
			).Description("What do you want to do?").Value(&se),
		)).Run()
	if err != nil {
		return se, err
	}
	return se, nil
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

func runLandingLLMActionForm(ctx context.Context) error {
	defer helper.ClearScreen()

	conf := config.Read()
	printConfigLLMProvidersData(ctx)
	se, err := configLandingLLMActionForm()
	if err != nil {
		return err
	}
	switch se {
	case configLandingActionAdd:
		var llmProviders []llmType.Provider
		for _, p := range llm.AvailProviders() {
			pr, err := llm.GetProvider(p)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Client] Failed to get LLM provider %s: %q", p, err))
			}
			llmProviders = append(llmProviders, pr)
		}

		helper.ClearScreen()
		llmProvider, err := SelectLLMForm(ctx, llmProviders, "Please select a LLM provider you want to add:")
		if err != nil {
			return err
		}

		cf, err := RunOnboardConfig(ctx, llmProvider.Config())
		if err != nil {
			return err
		}

		conf.LLMs = append(conf.LLMs, config.LLMConfig{
			Provider:      llmProvider.Config().Name(),
			Configuration: cf,
		})

		err = config.Write(conf)
		if err != nil {
			return err
		}
		color.Blue("Config Updated!")
		time.Sleep(1 * time.Second)
	case configLandingActionSelect:
		helper.ClearScreen()
		llmProvider, err := SelectLLMProviderInstanceForm(ctx, conf.LLMs, "Please select a LLM provider:", "")
		if err != nil {
			return err
		}

		oldName := llmProvider.Name()

		llmProviders, confData := GetSelectedLLMConfig(ctx, conf.LLMs, llmProvider.Name())
		llmProvidersMap := map[string]llmType.Provider{}
		for _, p := range llmProviders {
			llmProvidersMap[p.Name()] = p
		}

		sel, err := configPageActionSelectForm("Edit this LLM Provider", "Delete this LLM Provider", "Quit")
		if err != nil {
			return err
		}
		switch sel {
		case configLandingPageEdit:
			edited, err := RunEditConfig(ctx, llmProvider.Config(), confData)
			if err != nil {
				return err
			}
			for i, c := range conf.LLMs {
				p, ok := llmProvidersMap[c.Provider]
				if ok && p.Name() == oldName {
					conf.LLMs[i].Configuration = edited
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
				p, ok := llmProvidersMap[c.Provider]
				if ok && p.Name() == llmProvider.Name() {
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
		}
	case configLandingActionQuit:
		color.Blue("Bye!")
		return nil
	default:
		return fmt.Errorf("unexpected landing LLM-action form: %s", se)
	}
	return nil
}

func runLandingToolActionForm(ctx context.Context) error {
	defer helper.ClearScreen()

	printConfigToolProvidersData(ctx)
	se, err := configLandingToolActionForm()
	if err != nil {
		return err
	}
	switch se {
	case configLandingActionAdd:
		conf := config.Read()
		var toolProviders []tool.Provider

		occupy := make(map[string]bool)
		for _, c := range conf.Tools {
			occupy[c.Name] = true
		}

		helper.ClearScreen()
		for _, p := range tool.AvailProviders() {
			if !occupy[p] {
				pr, err := tool.GetProvider(p)
				if err != nil {
					color.Yellow(fmt.Sprintf("[Manboster Client] Failed to get tool provider %s: %q", p, err))
				}
				toolProviders = append(toolProviders, pr)
			}
		}

		if len(toolProviders) == 0 {
			color.Yellow("[Manboster Client] No new tool providers available to add!")
			time.Sleep(1 * time.Second)
			return nil
		}

		toolProvider, err := SelectSingleToolForm(ctx, toolProviders, "Please select the Tool provider you want to add:")
		if err != nil {
			return err
		}

		if toolProvider.Config() != nil {
			cf, err := RunOnboardConfig(ctx, toolProvider.Config())
			if err != nil {
				return err
			}
			conf.Tools = append(conf.Tools, config.ToolConfig{
				Name:          toolProvider.Name(),
				Configuration: cf,
			})
		} else {
			conf.Tools = append(conf.Tools, config.ToolConfig{
				Name: toolProvider.Name(),
			})
		}
	case configLandingActionSelect:
	case configLandingActionQuit:
		color.Blue("Bye!")
		return nil
	default:
		return fmt.Errorf("unexpected landing tool-action form: %s", se)
	}
	return nil
}

func runLandingHachimiActionForm(ctx context.Context) error {
	printConfigHachimiProvidersData(ctx)
	se, err := configLandingChatActionForm()
	if err != nil {
		return err
	}
	switch se {
	case configLandingActionAdd:
	case configLandingActionSelect:
	case configLandingActionQuit:
		color.Blue("Bye!")
		return nil
	default:
		return fmt.Errorf("unexpected landing hachimi-action form: %s", se)
	}
	return nil
}

func configPageActionSelectForm(edit string, del string, quit string) (configLandingPageActionSelection, error) {
	var se configLandingPageActionSelection
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[configLandingPageActionSelection]().Options(
				huh.NewOption(edit, configLandingPageEdit),
				huh.NewOption(del, configLandingPageDelete),
				huh.NewOption(quit, configLandingPageQuit),
			).Description("What do you want to do?").Value(&se),
		)).Run()
	if err != nil {
		return se, err
	}
	return se, nil
}
