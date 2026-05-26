package gatekeeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/hachimi"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

func (s *Service) HachimiHandler(ctx context.Context, instance chat.Provider, mes *chat.Message, toolProvider tool.Provider, req llm.MessageToolCallRequestPayload, id string) (bool, error) {
	msg := mes.Clone()
	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{}

	if !*s.hachimiLoaded || s.hachimiProvider == nil {
		color.Yellow("[Manboster Gatekeeper] Hachimi is not loaded!")
		return true, nil
	}

	desc := util.DescribeToHachimi(req, toolProvider)
	u, avail := s.ignoranceSessionManager.GetHachimiCache(desc)
	if avail {
		if !u {
			return false, fmt.Errorf("hachimi thinks it's unsafe and user denied it")
		}
		return true, nil
	}

	resp, err := s.hachimiProvider.Chat(ctx, desc)
	if err != nil {
		return false, err
	}
	s.ignoranceSessionManager.UpdateMark(id)

	if resp == nil {
		resp = &hachimi.Response{
			Type:   hachimi.ResponseStatusUnsafe,
			Reason: "input message is too long so hachimi could not make decision for you",
		}
	}

	switch resp.Type {
	case hachimi.ResponseStatusUnsafe:
		var txt strings.Builder
		txt.WriteString(i18n.T(keys.GatekeeperHachimiUnsafe))
		txt.WriteString(util.DescribeToHuman(req, toolProvider))
		txt.WriteString(fmt.Sprintf(i18n.T(keys.GatekeeperHachimiReason), resp.Reason))
		return s.Select(ctx, instance, msg, buildSelectionHachimi(), txt.String(), func(msg *chat.Message) (bool, error) {
			cb := msg.SelectionCallback
			switch cb.SelectionValue {
			case "allow":
				s.ignoranceSessionManager.SetHachimiCache(desc, true)
				return true, nil
			case "deny":
				s.ignoranceSessionManager.SetHachimiCache(desc, false)
				return false, fmt.Errorf("hachimi thinks it's unsafe and user denied it")
			}
			return false, fmt.Errorf("invalid selection value: %s", cb.SelectionValue)
		})
	case hachimi.ResponseStatusInspect:
		var txt strings.Builder
		txt.WriteString(i18n.T(keys.GatekeeperHachimiSuspicious))
		txt.WriteString(util.DescribeToHuman(req, toolProvider))
		txt.WriteString(fmt.Sprintf(i18n.T(keys.GatekeeperHachimiReason), resp.Reason))
		return s.Select(ctx, instance, msg, buildSelectionHachimi(), txt.String(), func(msg *chat.Message) (bool, error) {
			cb := msg.SelectionCallback
			switch cb.SelectionValue {
			case "allow":
				s.ignoranceSessionManager.SetHachimiCache(desc, true)
				return true, nil
			case "deny":
				s.ignoranceSessionManager.SetHachimiCache(desc, false)
				return false, fmt.Errorf("hachimi thinks it's suspicious and user denied it")
			}
			return false, fmt.Errorf("invalid selection value: %s", cb.SelectionValue)
		})
	case hachimi.ResponseStatusSafe:
		s.ignoranceSessionManager.SetHachimiCache(desc, true)
		color.Blue("[Manboster Gatekeeper] Hachimi thinks it's safe to go!")
		msg.Text.Text = i18n.T(keys.GatekeeperHachimiHandled)
		err := s.gatewayService.SendMessage(ctx, instance, msg)
		if err != nil {
			color.Yellow("[Manboster Gatekeeper] Failed to send ignore prompt message")
		}
		go s.Recall(ctx, instance, msg)
		return true, nil
	default:
		return false, fmt.Errorf("unexpected response type: %v", resp.Type)
	}
}
