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
	color.Yellow(fmt.Sprintf("[Manboster Engine] Model called tool %q", toolProvider.DisplayName()))
	respData, err := toolProvider.Run(ctx, args)
	if err != nil {
		return "", err
	}
	if !respData.Hangup {
		return respData.Response, nil
	}
	// hangup process...
	return "", nil
}
