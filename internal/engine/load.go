package engine

import (
	"context"
)

func (e *Engine) Load(ctx context.Context) error {
	// initialize variables

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
