package model

import "github.com/manboster/manboster/internal/llm"

var QwenModels = []llm.Model{
	{
		DisplayName:     "Qwen 3.6 Plus",
		Name:            "qwen/qwen3.6-plus",
		Context:         1000000,
		MaxOutputTokens: 66000,
		InputPrice:      0.325,
		OutputPrice:     1.95,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Qwen 3.5 Flash",
		Name:            "qwen/qwen3.5-flash-02-23",
		Context:         1000000,
		MaxOutputTokens: 66000,
		InputPrice:      0.065,
		OutputPrice:     0.26,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Qwen 3.5 397B A17B",
		Name:            "qwen/qwen3.5-397b-a17b",
		Context:         262144,
		MaxOutputTokens: 66000,
		InputPrice:      0.39,
		OutputPrice:     2.34,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
}
