package handler

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

func (h *Handler) HandleToolExec(ctx context.Context, name string, args string) (string, error) {
	toolProvider, avail := h.toolMaps[name]
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
