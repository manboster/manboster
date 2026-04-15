package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

func (e *Engine) loadDefaultModel(ctx context.Context) {
	pIndex, mIndex := e.modelIndexWithFallback(ctx, e.config.App.DefaultLLMProvider, e.config.App.DefaultLLMModel)
	e.config.App.DefaultLLMProvider = e.llmProviders[pIndex].Name()
	e.config.App.DefaultLLMModel = e.llmProviders[pIndex].Models()[mIndex].Name
	color.Blue(fmt.Sprintf("[Manboster Engine] Default LLM Model loaded. Provider: %s, Model: %s", e.llmProviders[pIndex].DisplayName(), e.llmProviders[pIndex].Models()[mIndex].DisplayName))
}
