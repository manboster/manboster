package model

import (
	"github.com/manboster/manboster/spec/llm"
)

// XAIModels defines models from elon mu$k's xAI.
var XAIModels = []llm.Model{
	{
		DisplayName:     "Grok 4.20",
		Name:            "x-ai/grok-4.20",
		Context:         2000000,
		MaxOutputTokens: 256000,
		InputPrice:      2,
		OutputPrice:     6,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityFile, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Grok 4.1 Fast",
		Name:            "x-ai/grok-4.1-fast",
		Context:         2000000,
		MaxOutputTokens: 30000,
		InputPrice:      0.2,
		OutputPrice:     0.5,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityFile, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Grok 4",
		Name:            "x-ai/grok-4",
		Context:         256000,
		MaxOutputTokens: 66000,
		InputPrice:      3,
		OutputPrice:     15,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityFile, Output: llm.CapabilityText},
	},
}
