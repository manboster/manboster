package util

import (
	"github.com/manboster/manboster/spec/llm"
)

func CalculateCost(e *llm.Event, m llm.Model) {
	if e.EventType&llm.EventUsage != 0 && e.Usage != nil {
		inputTokens := e.Usage.PromptTokens
		outputTokens := e.Usage.TotalTokens - inputTokens
		e.Usage.InputCost = float64(inputTokens) * m.InputPrice / 1000000
		e.Usage.OutputCost = float64(outputTokens) * m.OutputPrice / 1000000
		e.Usage.TotalCost = e.Usage.InputCost + e.Usage.OutputCost
	}
}
