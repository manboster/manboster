package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

func (e *Engine) Load(ctx context.Context) error {
	// initialize variables

	color.Blue("[Manboster Engine] Loading engine...")

	// First, we get user counts using cache
	count, err := e.repo.UserCounts(ctx)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while getting user counts from repository, error: %s", err))
		e.userCount = 0
	}
	e.userCount = count

	// Then, we activate LLMs.
	llmProviders, err := loadLLM(ctx, e.config.LLMs)
	if err != nil {
		return err
	}

	if len(llmProviders) == 0 {
		color.Red(fmt.Sprintf("[Manboster Engine] All LLM providers have tried to connect but failed."))
		return ErrNoAvailableLLMProvider
	}
	e.llmProviders = llmProviders

	// get model data from configuration
	e.loadDefaultModel(ctx)

	// Then, we activate chats.
	err = e.loadChats(ctx)
	if err != nil {
		return err
	}
	return nil
}
