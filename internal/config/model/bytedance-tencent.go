package model

import "github.com/manboster/manboster/spec/llm"

var ByteDanceModels = []llm.Model{
	// current missing Seed Models
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
