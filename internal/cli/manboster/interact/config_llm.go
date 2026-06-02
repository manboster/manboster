package interact

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/llm"
	_ "github.com/manboster/manboster/internal/llm/all"
	"github.com/manboster/manboster/spec/cli"
	llmType "github.com/manboster/manboster/spec/llm"
)

type llmConfigAction string

const (
	llmConfigDelete llmConfigAction = _DELETE_
	llmConfigEdit   llmConfigAction = _EDIT_
	llmConfigQuit   llmConfigAction = _QUIT_
)

func (a llmConfigAction) Name() string {
	return string(a)
}

func (a llmConfigAction) DisplayName() string {
	switch a {
	case llmConfigDelete:
		return i18n.T(keys.CliConfigActionDeleteProvider)
	case llmConfigEdit:
		return i18n.T(keys.CliConfigActionEditProvider)
	case llmConfigQuit:
		return i18n.T(keys.CliConfigActionQuit)
	default:
		return ""
	}
}

func runLLMConfigs(p cli.Provider, cfg config.Config) ([]config.LLMConfig, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var option cli.Option
	for {
		// reload on every iteration so changes are reflected
		var llmProviders []llmType.Provider
		for _, c := range cfg.LLMs {
			provider, err := llm.GetProvider(c.Provider)
			if err != nil {
				return nil, err
			}

			err = provider.Init(ctx, c.Configuration)
			if err != nil {
				continue
			}

			llmProviders = append(llmProviders, provider)
		}

		options := cli.BuildOptions[llmType.Provider](llmProviders, nil)
		options = append(options, addOption, quitOption)

		var err error
		option, err = p.Select(i18n.T(keys.CliConfigLLMSelectPrompt), i18n.T(keys.CliConfigLLMSelectHelp), options, option.Value, func(option cli.Option) error {
			for _, o := range options {
				if o.Value == option.Value {
					return nil
				}
			}
			return fmt.Errorf("unknown LLM provider selected: %s", option.Value)
		})
		if err != nil {
			return nil, err
		}

		if option.Value == _QUIT_ {
			return cfg.LLMs, nil
		}

		if option.Value == _ADD_ {
			llmConfig, err := runOnboardLLMConfig(p, llm.AllProviders())
			if err != nil {
				return nil, err
			}
			cfg.LLMs = append(cfg.LLMs, llmConfig)
			continue
		}

		var selectedConfig config.LLMConfig
		var selectedProvider llmType.Provider
		selectedIndex := -1
		for i, c := range cfg.LLMs {
			pr, err := llm.GetProvider(c.Provider)
			if err != nil {
				return nil, err
			}
			err = pr.Init(ctx, c.Configuration)
			if err != nil {
				continue
			}
			if pr.Name() == option.Value {
				selectedConfig = c
				selectedIndex = i
				selectedProvider = pr
				break
			}
		}

		if selectedIndex == -1 {
			return nil, fmt.Errorf("unknown LLM provider selected: %s", option.Value)
		}

		se := []llmConfigAction{llmConfigEdit, llmConfigDelete, llmConfigQuit}
		opts := cli.BuildOptions[llmConfigAction](se, nil)
		form := newConfigForm[llmConfigAction]()

		form.Register(llmConfigDelete, func() error {
			confirm, err := p.Prompt(i18n.Te(keys.CliConfigLLMDeleteConfirm, selectedConfig.Provider, nil), "Do you want to continue?", "Yes", "No")
			if err != nil {
				return err
			}
			if !confirm {
				return fmt.Errorf("cancelled")
			}
			cfg.LLMs = append(cfg.LLMs[:selectedIndex], cfg.LLMs[selectedIndex+1:]...)
			if err := p.Alert(i18n.T(keys.CliWizardTitle), i18n.Te(keys.CliConfigLLMDeleteSuccess, selectedConfig.Provider, nil)); err != nil {
				return err
			}
			return errQuit
		})

		form.Register(llmConfigEdit, func() error {
			conf, err := EditConfig(ctx, p, selectedProvider.Config(), selectedConfig.Configuration)
			if err != nil {
				return err
			}
			selectedConfig.Configuration = conf
			cfg.LLMs[selectedIndex] = selectedConfig
			return errQuit
		})

		form.Register(llmConfigQuit, nilFunc)

		err = handleWithPrompt[llmConfigAction](p, form, opts, fmt.Sprintf("This LLM provider %s's info:\n\n%s", selectedProvider.DisplayName(), selectedConfig.Configuration), i18n.T(keys.CliConfigActionWhatToDo))
		if err != nil {
			return nil, err
		}
	}
}
