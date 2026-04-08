package model

import "github.com/manboster/manboster/internal/llm"

// GoogleModels defines Gemini 3 & Gemma 4 to now
var GoogleModels = []llm.Model{
	{
		DisplayName:     "Gemini 3.1 Pro Preview",
		Name:            "google/gemini-3.1-pro-preview",
		Context:         1048576,
		MaxOutputTokens: 66000,
		InputPrice:      2,
		OutputPrice:     12,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityAll, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Gemini 3 Flash Preview",
		Name:            "google/gemini-3-flash-preview",
		Context:         1048576,
		MaxOutputTokens: 66000,
		InputPrice:      0.5,
		OutputPrice:     3,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityAll, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Google Gemma 4 26B A4B(Free Model)",
		Name:            "google/gemma-4-26b-a4b-it:free",
		Context:         262144,
		MaxOutputTokens: 8192, // max output token defines as 262144 in openrouter, we hard limit it to 8192.
		InputPrice:      0,
		OutputPrice:     0,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Google Gemma 4 26B A4B",
		Name:            "google/gemma-4-26b-a4b-it",
		Context:         262144,
		MaxOutputTokens: 8192, // max output token defines as 262144 in openrouter, we hard limit it to 8192.
		InputPrice:      0.13,
		OutputPrice:     0.4,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Google Gemma 4 31B(Free Model)",
		Name:            "google/gemma-4-31b-it:free",
		Context:         262144,
		MaxOutputTokens: 33000,
		InputPrice:      0,
		OutputPrice:     0,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Google Gemma 4 31B",
		Name:            "google/gemma-4-31b-it",
		Context:         262144,
		MaxOutputTokens: 33000,
		InputPrice:      0.14,
		OutputPrice:     0.4,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
}
