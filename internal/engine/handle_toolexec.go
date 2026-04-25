package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

func (e *Engine) HandleToolExec(ctx context.Context, name string, args string) (string, error) {
	toolProvider, avail := e.toolMaps[name]
	if !avail {
		return "", fmt.Errorf("there is no tool named %s", name)
	}
	color.Yellow(fmt.Sprintf("[Manboster] Model called tool %q", toolProvider.DisplayName()))
	return toolProvider.Run(ctx, args)
}
