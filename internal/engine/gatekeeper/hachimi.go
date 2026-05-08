package gatekeeper

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/hachimi"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

func (s *Service) HachimiHandler(ctx context.Context, instance chat.Provider, msg *chat.Message, toolProvider tool.Provider, req llm.MessageToolCallRequestPayload, sid string) (bool, error) {
	if !*s.hachimiLoaded || s.hachimiProvider == nil {
		color.Yellow("[Manboster Gatekeeper] Hachimi is not loaded!")
		return true, nil
	}
	resp, err := s.hachimiProvider.Chat(ctx, util.DescribeToHachimi(req, toolProvider))
	if err != nil {
		return false, err
	}

	switch resp.Type {
	case hachimi.ResponseStatusUnsafe:
		// unsafe prompt
	case hachimi.ResponseStatusInspect:
		// inspect prompt
	case hachimi.ResponseStatusSafe:
		color.Blue("[Manboster Gatekeeper] Hachimi thinks it's safe to go!")
		return true, nil
	default:
		return false, fmt.Errorf("unexpected response type: %v", resp.Type)
	}
	return false, nil
}
