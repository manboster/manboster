package model

import "github.com/manboster/manboster/internal/llm"

// DeepSeekModels defines models from DeepSeek v3.2 (Along with R1) to now
var DeepSeekModels = []llm.Model{
	{
		DisplayName:     "DeepSeek V3.2",
		Name:            "deepseek/deepseek-v3.2",
		Context:         163840,
		MaxOutputTokens: 16384,
		InputPrice:      0.26,
		OutputPrice:     0.38,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "DeepSeek R1 0528",
		Name:            "deepseek/deepseek-r1-0528",
		Context:         163840,
		MaxOutputTokens: 16384,
		InputPrice:      0.45,
		OutputPrice:     2.15,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
}
