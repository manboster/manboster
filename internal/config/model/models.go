package model

import (
	"slices"

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
	)
}
