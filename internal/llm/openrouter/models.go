package openrouter

import (
	"github.com/manboster/manboster/internal/config/model"
	"github.com/manboster/manboster/spec/llm"
)

// models' information from openrouter.ai

var initialModels = []llm.Model{
	// Openrouter Specific
	{
		displayName:     "OpenRouter Auto(Smart Routing, Dynamic Billing)",
		name:            "openrouter/auto",
		Context:         500000,
		MaxOutputTokens: 8192, // not given, so give 8192
		InputPrice:      0,    // this is not defined, based on model, this may vary, please see openrouter's documentation.
		OutputPrice:     0,    // this is not defined, based on model, this may vary, please see openrouter's documentation.
		Capabilities:    llm.Capabilities{Input: llm.CapabilityAll, Output: llm.CapabilityText},
	},
	{
		displayName:     "OpenRouter Free Model",
		name:            "openrouter/free",
		Context:         131072,
		MaxOutputTokens: 4096, // not given, so give 4096
		InputPrice:      0,    // it's free!
		OutputPrice:     0,    // it's free!
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},
	{
		displayName:     "Minimax M2.5(Free Model)",
		name:            "minimax/minimax-m2.5:free",
		Context:         197000,
		MaxOutputTokens: 66000,
		InputPrice:      0,
		OutputPrice:     0,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		displayName:     "Google Gemma 4 26B A4B(Free Model)",
		name:            "google/gemma-4-26b-a4b-it:free",
		Context:         262144,
		MaxOutputTokens: 8192, // max output token defines as 262144 in openrouter, we hard limit it to 8192.
		InputPrice:      0,
		OutputPrice:     0,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
	{
		displayName:     "Google Gemma 4 31B(Free Model)",
		name:            "google/gemma-4-31b-it:free",
		Context:         262144,
		MaxOutputTokens: 33000,
		InputPrice:      0,
		OutputPrice:     0,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
	{
		displayName:     "GPT oss 20b(Free Model)",
		name:            "openai/gpt-oss-20b:free",
		Context:         131000,
		MaxOutputTokens: 8192,
		InputPrice:      0,
		OutputPrice:     0,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		displayName:     "GPT oss 120b(Free Model)",
		name:            "openai/gpt-oss-120b:free",
		Context:         131000,
		MaxOutputTokens: 8192,
		InputPrice:      0,
		OutputPrice:     0,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
}

func Models() []llm.Model {
	return append(initialModels, model.Models()...)
}

//var models = []string{
//	"qwen/qwen3.5-397b-a17b",
//	"qwen/qwen3.5-flash-02-23",
//}
