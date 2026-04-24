package model

import (
	"slices"
	"strings"

	"github.com/manboster/manboster/internal/llm"
)

func Models() []llm.Model {
	return slices.Concat(
		GoogleModels,
		AnthropicModels,
		OpenAIModels,
		MoonshotModels,
		GlmModels,
		StepFunModels,
		MinimaxModels,
		QwenModels,
		XAIModels,
		XiaomiModels,
		DeepSeekModels,
	)
}

// Default returns a default model definition
func Default(m string) llm.Model {
	return llm.Model{
		Name:            m, // to be filled
		DisplayName:     m, // to be filled
		Context:         262144,
		MaxOutputTokens: 8192,
		InputPrice:      0,
		OutputPrice:     0,
		Capabilities: llm.Capabilities{
			Input:  llm.CapabilityText,
			Output: llm.CapabilityText,
		},
	}
}

// Search is used to find proper model for you in the models library
func Search(keyword string) (llm.Model, bool) {
	keywordLower := strings.ToLower(keyword)
	for _, model := range Models() {
		lowerName := strings.ToLower(model.Name)
		if keywordLower == lowerName {
			return model, true
		}
		splitNames := strings.Split(lowerName, "/")
		if len(splitNames) > 1 && splitNames[1] == keywordLower {
			return model, true
		}
	}
	return llm.Model{}, false
}
