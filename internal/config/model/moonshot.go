package model

import "github.com/manboster/manboster/internal/llm"

var MoonshotModels = []llm.Model{
	//	"moonshotai/kimi-k2.6",
	{
		DisplayName:     "Kimi K2.6",
		Name:            "moonshotai/kimi-k2.6",
		Context:         262000,
		MaxOutputTokens: 66000,
		InputPrice:      0.8,
		OutputPrice:     3.5,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},
	//	"moonshotai/kimi-k2.5",
	{
		DisplayName:     "Kimi K2.5",
		Name:            "moonshotai/kimi-k2.5",
		Context:         262000,
		MaxOutputTokens: 66000,
		InputPrice:      0.3827,
		OutputPrice:     1.72,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},
	// kimi k2 thinking
	{
		DisplayName:     "Kimi K2 Thinking",
		Name:            "moonshotai/kimi-k2-thinking",
		Context:         262000,
		MaxOutputTokens: 16384,
		InputPrice:      0.6,
		OutputPrice:     2.5,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
}
