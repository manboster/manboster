package model

import "github.com/manboster/manboster/internal/llm"

// XiaomiModels defines MiMo models
var XiaomiModels = []llm.Model{
	//	"xiaomi/mimo-v2-pro",
	//	"xiaomi/mimo-v2-flash",
	{
		DisplayName:     "Xiaomi MiMo v2.5 Pro",
		Name:            "xiaomi/mimo-v2.5-pro",
		Context:         1048576,
		MaxOutputTokens: 131000,
		InputPrice:      1,
		OutputPrice:     3,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Xiaomi MiMo v2.5",
		Name:            "xiaomi/mimo-v2.5",
		Context:         1048576,
		MaxOutputTokens: 131000,
		InputPrice:      0.4,
		OutputPrice:     2,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityAudio | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Xiaomi MiMo v2 Pro",
		Name:            "xiaomi/mimo-v2-pro",
		Context:         1048576,
		MaxOutputTokens: 103000,
		InputPrice:      1,
		OutputPrice:     3,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Xiaomi MiMo v2 Omni",
		Name:            "xiaomi/mimo-v2-omni",
		Context:         262144,
		MaxOutputTokens: 66000,
		InputPrice:      0.4,
		OutputPrice:     2,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityAudio | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "Xiaomi MiMo v2 Flash",
		Name:            "xiaomi/mimo-v2-flash",
		Context:         262144,
		MaxOutputTokens: 66000,
		InputPrice:      0.09,
		OutputPrice:     0.29,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
}
