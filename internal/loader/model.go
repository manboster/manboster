package loader

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/util"
)

func (l *Loader) loadDefaultModel(ctx context.Context) {
	pIndex, mIndex := util.GetModelIndexWithFallback(ctx, l.llmProviders, l.cfg.App.DefaultLLMProvider, l.cfg.App.DefaultLLMModel)
	l.cfg.App.DefaultLLMProvider = l.llmProviders[pIndex].Name()
	l.cfg.App.DefaultLLMModel = l.llmProviders[pIndex].Models()[mIndex].Name
	color.Blue(fmt.Sprintf("[Manboster Engine] Default LLM Model loaded. Provider: %s, Model: %s", l.llmProviders[pIndex].DisplayName(), l.llmProviders[pIndex].Models()[mIndex].DisplayName))
}
