package model

import "github.com/manboster/manboster/internal/llm"

// AnthropicModels defines data from Claude Sonnet/Haiku/Opus 4.5 to now
var AnthropicModels = []llm.Model{
	{
		DisplayName:     "Claude Opus 4.7",
		Name:            "anthropic/claude-opus-4.7",
		Context:         1000000,
		MaxOutputTokens: 128000,
		InputPrice:      5,
		OutputPrice:     25,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Claude Sonnet 4.6",
		Name:            "anthropic/claude-sonnet-4.6",
		Context:         1000000,
		MaxOutputTokens: 128000,
		InputPrice:      3,
		OutputPrice:     15,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Claude Opus 4.6",
		Name:            "anthropic/claude-opus-4.6",
		Context:         1000000,
		MaxOutputTokens: 128000,
		InputPrice:      5,
		OutputPrice:     25,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Claude Sonnet 4.5",
		Name:            "anthropic/claude-sonnet-4.5",
		Context:         1000000,
		MaxOutputTokens: 64000,
		InputPrice:      3,
		OutputPrice:     15,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityFile, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Claude Opus 4.5",
		Name:            "anthropic/claude-opus-4.5",
		Context:         200000,
		MaxOutputTokens: 64000,
		InputPrice:      5,
		OutputPrice:     25,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityFile, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Claude Haiku 4.5",
		Name:            "anthropic/claude-haiku-4.5",
		Context:         200000,
		MaxOutputTokens: 64000,
		InputPrice:      1,
		OutputPrice:     5,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},
}
