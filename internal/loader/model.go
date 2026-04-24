package loader

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/util"
)

func (l *Loader) loadDefaultModel(ctx context.Context) {
	p, m := util.GetModelWithFallback(ctx, l.llmProviders, l.cfg.App.DefaultLLMProvider, l.cfg.App.DefaultLLMModel)
	l.cfg.App.DefaultLLMProvider = p.Name()
	l.cfg.App.DefaultLLMModel = m.Name
	color.Blue(fmt.Sprintf("[Manboster Engine] Default LLM Model loaded. Provider: %s, Model: %s", p.DisplayName(), m.DisplayName))
}
