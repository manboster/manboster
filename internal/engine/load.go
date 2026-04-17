package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/engine/onboard"
)

func (e *Engine) Load(ctx context.Context) error {
	// initialize variables
	color.Blue("[Manboster Engine] Loading engine...")

	// First, we get user counts using cache
	count, err := e.repo.UserCounts(ctx)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while getting user counts from repository, error: %s", err))
	}
	if err != nil || count == 0 {
		e.onboard = onboard.New()
	}

	return nil
}
