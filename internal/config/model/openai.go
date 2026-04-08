package model

import "github.com/manboster/manboster/internal/llm"

// OpenAIModels defines models from GPT 5.3 series to now.
var OpenAIModels = []llm.Model{
	{
		DisplayName:     "GPT 5.4 Nano",
		Name:            "openai/gpt-5.4-nano",
		Context:         400000,
		MaxOutputTokens: 128000,
		InputPrice:      0.2,
		OutputPrice:     1.25,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityFile, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GPT 5.4 Mini",
		Name:            "openai/gpt-5.4-mini",
		Context:         400000,
		MaxOutputTokens: 128000,
		InputPrice:      0.75,
		OutputPrice:     4.5,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityFile, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GPT 5.4",
		Name:            "openai/gpt-5.4",
		Context:         1050000,
		MaxOutputTokens: 128000,
		InputPrice:      2.5,
		OutputPrice:     15,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityFile, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GPT 5.4 Pro",
		Name:            "openai/gpt-5.4-pro",
		Context:         1050000,
		MaxOutputTokens: 128000,
		InputPrice:      30,
		OutputPrice:     180,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityFile, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GPT 5.3 Codex",
		Name:            "openai/gpt-5.3-codex",
		Context:         400000,
		MaxOutputTokens: 128000,
		InputPrice:      1.75,
		OutputPrice:     14,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityFile, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GPT 5.3 Chat",
		Name:            "openai/gpt-5.3-chat",
		Context:         128000,
		MaxOutputTokens: 1600,
		InputPrice:      1.75,
		OutputPrice:     14,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityFile, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GPT oss 20b",
		Name:            "openai/gpt-oss-20b",
		Context:         131000,
		MaxOutputTokens: 8192,
		InputPrice:      0.03,
		OutputPrice:     0.11,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GPT oss 120b",
		Name:            "openai/gpt-oss-120b",
		Context:         131000,
		MaxOutputTokens: 8192,
		InputPrice:      0.039,
		OutputPrice:     0.19,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
}
