package model

import (
	"github.com/manboster/manboster/spec/llm"
)

// GlmModels defines Models from bigmodel.cn / z.ai
var GlmModels = []llm.Model{
	{
		DisplayName:     "GLM 5.2",
		Name:            "z-ai/glm-5.2",
		Context:         1048576,
		MaxOutputTokens: 131072,
		InputPrice:      0.97,
		OutputPrice:     3.04,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GLM 5.1",
		Name:            "z-ai/glm-5.1",
		Context:         202752,
		MaxOutputTokens: 66000,
		InputPrice:      1.26,
		OutputPrice:     3.96,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GLM 5V Turbo",
		Name:            "z-ai/glm-5v-turbo",
		Context:         202752,
		MaxOutputTokens: 131000,
		InputPrice:      1.2,
		OutputPrice:     4,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GLM 5 Turbo",
		Name:            "z-ai/glm-5-turbo",
		Context:         202752,
		MaxOutputTokens: 131000,
		InputPrice:      1.2,
		OutputPrice:     4,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GLM 5",
		Name:            "z-ai/glm-5",
		Context:         204800,
		MaxOutputTokens: 131072,
		InputPrice:      0.72,
		OutputPrice:     2.3,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GLM 4.7",
		Name:            "z-ai/glm-4.7",
		Context:         202752,
		MaxOutputTokens: 131072,
		InputPrice:      0.4,
		OutputPrice:     1.75,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GLM 4.7-flash",
		Name:            "z-ai/glm-4.7",
		Context:         200000,
		MaxOutputTokens: 131072,
		InputPrice:      0.06,
		OutputPrice:     0.40,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
}
