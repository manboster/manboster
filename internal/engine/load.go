package engine

import (
	"context"
	"fmt"
)

func (e *Engine) Load(ctx context.Context) error {
	// initialize variables

	fmt.Println("[Manboster Engine] Loading engine...")

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
