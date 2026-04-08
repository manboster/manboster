package model

import "github.com/manboster/manboster/internal/llm"

var StepFunModels = []llm.Model{
	//	"stepfun/step-3.5-flash:free",
	//	"stepfun/step-3.5-flash",
	{
		DisplayName:     "StepFun Step 3.5 Flash",
		Name:            "stepfun/step-3.5-flash",
		Context:         262144,
		MaxOutputTokens: 66000,
		InputPrice:      0.10,
		OutputPrice:     0.30,
		Capabilities:    llm.Capabilities{Input: llm.CapabilityText, Output: llm.CapabilityText},
	},
}
