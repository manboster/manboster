package handler

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/tool"
)

func (h *Handler) HandleToolExec(ctx context.Context, tool tool.Provider, args string) (string, error) {
	color.Yellow(fmt.Sprintf("[Manboster Handler] Model called tool %q", tool.DisplayName()))
	respData, err := tool.Run(ctx, args)
	if err != nil {
		return "", err
	}
	if !respData.Hangup {
		return respData.Response, nil
	}
	// hangup process...
	return "", nil
}
