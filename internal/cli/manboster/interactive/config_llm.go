package interactive

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	llmType "github.com/manboster/manboster/spec/llm"
)

func configLandingLLMActionForm() (configLandingActionSelection, error) {
	return configLandingActionForm("Add a new LLM provider", "Select an existing LLM provider", "Quit")
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

		err = config.Write(conf, config.Path("config.yaml"))
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
			err = config.Write(conf, config.Path("config.yaml"))
			if err != nil {
				return err
			}
			color.Blue("Config Updated!")
			time.Sleep(1 * time.Second)
		case configLandingPageDelete:
			for i, c := range conf.LLMs {
				p, ok := llmProvidersMap[c.Provider]
				if ok && p.Name() == llmProvider.Name() {
					conf.LLMs = append(conf.LLMs[:i], conf.LLMs[i+1:]...)
					break
				}
			}
			err = config.Write(conf, config.Path("config.yaml"))
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
