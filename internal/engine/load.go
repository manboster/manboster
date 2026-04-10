package engine

import (
	"context"

	"github.com/fatih/color"
)

func (e *Engine) Load(ctx context.Context) error {
	// initialize variables

	color.Blue("[Manboster Engine] Loading engine...")

	// TODO: get model data from SQLite(Repository)
	// First, we activate LLMs.
	llmProviders, err := loadLLM(ctx, e.config.LLMs)
	if err != nil {
		return err
	}
	e.llmProviders = llmProviders

	// Then, we activate chats.
	err = e.loadChats(ctx)
	if err != nil {
		return err
	}
	return nil
}
