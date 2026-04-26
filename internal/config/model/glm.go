package model

import (
	"github.com/manboster/manboster/spec/llm"
)

// GlmModels defines Models from bigmodel.cn / z.ai
var GlmModels = []llm.Model{
	{
		DisplayName:     "GLM 5.1",
		Name:            "z-ai/glm-5.1",
		Context:         202752,
		MaxOutputTokens: 66000,
		InputPrice:      1.26,
		OutputPrice:     3.96,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GLM 5V Turbo",
		Name:            "z-ai/glm-5v-turbo",
		Context:         202752,
		MaxOutputTokens: 131000,
		InputPrice:      1.2,
		OutputPrice:     4,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityTextAndImage | llm.CapabilityVideo, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GLM 5 Turbo",
		Name:            "z-ai/glm-5-turbo",
		Context:         202752,
		MaxOutputTokens: 131000,
		InputPrice:      1.2,
		OutputPrice:     4,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
	{
		DisplayName:     "GLM 5",
		Name:            "z-ai/glm-5",
		Context:         81920,
		MaxOutputTokens: 16384,
		InputPrice:      0.72,
		OutputPrice:     2.3,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
}
