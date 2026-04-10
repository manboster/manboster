package oai_compat

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/internal/llm"
)

func InputModel() (llm.Model, error) {
	var customModel llm.Model

	// define middleware variables
	var modelContext string
	var modelMaxOutputTokens string
	var modelInputPrice string
	var modelOutputPrice string
	var modelInputCapabilities []llm.CapabilityType
	var modelOutputCapabilities []llm.CapabilityType

	// define overall capabilities, in order to avoid duplicated selection, so it made a factory anonymous function to get heap.
	modelCapabilityOptions := func() []huh.Option[llm.CapabilityType] {
		return []huh.Option[llm.CapabilityType]{
			huh.NewOption("Text", llm.CapabilityText).Selected(true),
			huh.NewOption("ToolCall", llm.CapabilityToolCall).Selected(true),
			huh.NewOption("Image", llm.CapabilityImage),
			huh.NewOption("Audio", llm.CapabilityAudio),
			huh.NewOption("Video", llm.CapabilityVideo),
			huh.NewOption("File", llm.CapabilityFile),
		}
	}

	// generate a new form
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Your Model Name").Description("Please specify the model name. You can copy it by clicking the clipboard icon on OpenRouter's model page.").Value(&customModel.Name),
			huh.NewInput().Title("Model's Context Length").Description("If you don't know what is this, please leave it empty, we will assign a default value.").Value(&modelContext),
			huh.NewInput().Title("Model's Max Output Tokens").Description("If you don't know what is this, please leave it empty, we will assign a default value.").Value(&modelMaxOutputTokens),
			huh.NewInput().Title("Model's Input Price($/1m tokens)").Description("The input prices of this model, if you don't know, you can leave it empty.").Value(&modelInputPrice),
			huh.NewInput().Title("Model's Output Price($/1m tokens)").Description("The output prices of this model, if you don't know, you can leave it empty.").Value(&modelOutputPrice),
			huh.NewMultiSelect[llm.CapabilityType]().Title("Model's Input Capability").Description("The input capabilities of this model, if you don't know, you can leave it default.").Options(
				modelCapabilityOptions()...,
			).Value(&modelInputCapabilities),
			huh.NewMultiSelect[llm.CapabilityType]().Title("Model's Output Capability").Description("The output capabilities of this model, if you don't know, you can leave it default.").Options(
				modelCapabilityOptions()...,
			).Value(&modelOutputCapabilities),
		)).Run()
	if err != nil {
		return llm.Model{}, err
	}

	// convert them into a proper value!
	if customModel.Name == "" {
		return llm.Model{}, ErrModelNameRequired
	}

	customModel.DisplayName = fmt.Sprintf("%s(Custom Model)", customModel.Name)
	customModel.Context = 262144
	customModel.MaxOutputTokens = 8192
	customModel.InputPrice = 0
	customModel.OutputPrice = 0
	customModel.Capabilities.Input = llm.MergeCapabilityFields(modelInputCapabilities)
	customModel.Capabilities.Output = llm.MergeCapabilityFields(modelOutputCapabilities)

	if modelContext != "" {
		modelContextInt, err := strconv.ParseUint(modelContext, 10, 64)
		if err != nil {
			return llm.Model{}, err
		}
		customModel.Context = modelContextInt
	}
	if modelMaxOutputTokens != "" {
		modelMaxOutputTokensInt, err := strconv.ParseUint(modelMaxOutputTokens, 10, 64)
		if err != nil {
			return llm.Model{}, err
		}
		customModel.MaxOutputTokens = modelMaxOutputTokensInt
	}
	if modelInputPrice != "" {
		modelInputPriceFloat, err := strconv.ParseFloat(modelInputPrice, 64)
		if err != nil {
			return llm.Model{}, err
		}
		customModel.InputPrice = modelInputPriceFloat
	}
	if modelOutputPrice != "" {
		modelOutputPriceFloat, err := strconv.ParseFloat(modelOutputPrice, 64)
		if err != nil {
			return llm.Model{}, err
		}
		customModel.OutputPrice = modelOutputPriceFloat
	}
	return customModel, nil
}
