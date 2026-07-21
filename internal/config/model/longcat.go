package model

import "github.com/manboster/manboster/spec/llm"

var LongCatModels = []llm.Model{
	{
		DisplayName:     "Longcat 2.0",
		Name:            "meituan/longcat-2.0",
		Context:         1048576,
		MaxOutputTokens: 131000,
		InputPrice:      0.3,
		OutputPrice:     1.2,
		Reasoning:       true,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
}
