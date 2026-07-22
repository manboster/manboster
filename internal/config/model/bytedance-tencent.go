package model

import "github.com/manboster/manboster/spec/llm"

var ByteDanceModels = []llm.Model{
	{
		DisplayName:     "Doubao Seed 2.1 Pro",
		Name:            "bytedance/doubao-seed-2-1-pro",
		Context:         262144,
		MaxOutputTokens: 262144,
		InputPrice:      0.9,
		OutputPrice:     4.43,
		CachedPrice:     0.18,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Doubao Seed 2.1 Turbo",
		Name:            "bytedance/doubao-seed-2-1-turbo",
		Context:         262144,
		MaxOutputTokens: 262144,
		InputPrice:      0.45,
		OutputPrice:     2.23,
		CachedPrice:     0.09,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Doubao Seed 2.0 Pro",
		Name:            "bytedance/doubao-seed-2-0-pro",
		Context:         262144,
		MaxOutputTokens: 262144,
		InputPrice:      0.87,
		OutputPrice:     4.33,
		CachedPrice:     0.17,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},

	{
		DisplayName:     "Doubao Seed 2.0 Code",
		Name:            "bytedance/doubao-seed-2-0-code",
		Context:         262144,
		MaxOutputTokens: 262144,
		InputPrice:      0.87,
		OutputPrice:     4.33,
		CachedPrice:     0.17,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},

	{
		DisplayName:     "Doubao Seed 2.0 Lite",
		Name:            "bytedance/doubao-seed-2-0-lite",
		Context:         262144,
		MaxOutputTokens: 262144,
		InputPrice:      0.16, // 均值: 1.10 CNY / 6.77
		OutputPrice:     0.97, // 均值: 6.60 CNY / 6.77
		CachedPrice:     0.03, // 均值: 0.22 CNY / 6.77
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},

	{
		DisplayName:     "Doubao Seed 2.0 Mini",
		Name:            "bytedance/doubao-seed-2-0-mini",
		Context:         262144,
		MaxOutputTokens: 262144,
		InputPrice:      0.07,
		OutputPrice:     0.69,
		CachedPrice:     0.01,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage, Output: llm.CapabilityText},
	},
}

var TencentModels = []llm.Model{
	{
		DisplayName:     "Hunyuan 3",
		Name:            "tencent/hy3",
		Context:         262144,
		MaxOutputTokens: 262144,
		InputPrice:      0.2,
		OutputPrice:     0.8,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
}
