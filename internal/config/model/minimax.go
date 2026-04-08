package model

import "github.com/manboster/manboster/internal/llm"

var MinimaxModels = []llm.Model{
	{
		DisplayName:     "Minimax M2.7",
		Name:            "minimax/minimax-m2.7",
		Context:         204800,
		MaxOutputTokens: 131000,
		InputPrice:      0.3,
		OutputPrice:     1.2,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Minimax M2.5",
		Name:            "minimax/minimax-m2.5",
		Context:         197000,
		MaxOutputTokens: 66000,
		InputPrice:      0.118,
		OutputPrice:     0.99,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Minimax M2-her",
		Name:            "minimax/minimax-m2-her",
		Context:         66000,
		MaxOutputTokens: 2048,
		InputPrice:      0.3,
		OutputPrice:     1.2,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
}
