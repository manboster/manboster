package openrouter

import (
	"github.com/manboster/manboster/internal/config/model"
	"github.com/manboster/manboster/internal/llm"
)

// models' information from openrouter.ai

var initialModels = []llm.Model{
	// Openrouter Specific
	{
		DisplayName:     "OpenRouter Auto(Smart Routing, Dynamic Billing)",
		Name:            "openrouter/auto",
		Context:         500000,
		MaxOutputTokens: 8192, // not given, so give 8192
		InputPrice:      0,    // this is not defined, based on model, this may vary, please see openrouter's documentation.
		OutputPrice:     0,    // this is not defined, based on model, this may vary, please see openrouter's documentation.
		Capabilities:    llm.Capabilities{Input: llm.CapabilityAll, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "OpenRouter Free Model",
		Name:            "openrouter/free",
		Context:         131072,
		MaxOutputTokens: 4096, // not given, so give 4096
		InputPrice:      0,    // it's free!
		OutputPrice:     0,    // it's free!
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},
}

func Models() []llm.Model {
	return append(initialModels, model.Models()...)
}

//var models = []string{
//	"qwen/qwen3.5-397b-a17b",
//	"qwen/qwen3.5-flash-02-23",
//}
