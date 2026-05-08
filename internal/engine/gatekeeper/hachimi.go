package gatekeeper

import (
	"context"
	"fmt"
	"strings"

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
		var txt strings.Builder
		txt.WriteString(fmt.Sprintf("**Hachimi thinks this tool call is unsafe! Please look at it carefully and decide!**\n"))
		txt.WriteString(util.DescribeToHuman(req, toolProvider))
		txt.WriteString(fmt.Sprintf("\nHachimi reports reason: `%s`\n", resp.Reason))
		return s.Select(ctx, instance, msg, selectionHachimi, txt.String(), func(cb *chat.SelectionCallbackPayload) (bool, error) {
			switch cb.SelectionValue {
			case "allow":
				return true, nil
			case "deny":
				return false, fmt.Errorf("hachimi thinks it's unsafe and user denied it")
			}
			return false, fmt.Errorf("invalid selection value: %s", cb.SelectionValue)
		})
	case hachimi.ResponseStatusInspect:
		var txt strings.Builder
		txt.WriteString(fmt.Sprintf("**Hachimi thinks this tool call is suspicious! Please look at it carefully and decide!**\n"))
		txt.WriteString(util.DescribeToHuman(req, toolProvider))
		txt.WriteString(fmt.Sprintf("\nHachimi reports reason: %s\n", resp.Reason))
		return s.Select(ctx, instance, msg, selectionHachimi, txt.String(), func(cb *chat.SelectionCallbackPayload) (bool, error) {
			switch cb.SelectionValue {
			case "allow":
				return true, nil
			case "deny":
				return false, fmt.Errorf("hachimi thinks it's suspicious and user denied it")
			}
			return false, fmt.Errorf("invalid selection value: %s", cb.SelectionValue)
		})
	case hachimi.ResponseStatusSafe:
		color.Blue("[Manboster Gatekeeper] Hachimi thinks it's safe to go!")
		return true, nil
	default:
		return false, fmt.Errorf("unexpected response type: %v", resp.Type)
	}
}
