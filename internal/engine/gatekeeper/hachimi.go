package gatekeeper

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/hachimi"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/session/ignorance"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

func (s *Service) HachimiHandler(ctx context.Context, instance chat.Provider, mes *chat.Message, toolProvider tool.Provider, req llm.MessageToolCallRequestPayload, id string, sid string) (bool, error) {
	msg := mes.Clone()
	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{}

	sessId := buildSessionId(instance.Name(), msg.ChatID, sid)
	_, markType := s.ignoranceSessionManager.GetMark(sessId)

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
		txt.WriteString(i18n.Te(keys.GatekeeperHachimiReason, "", errors.New(resp.Reason)))

		return s.Select(ctx, instance, msg, buildSelectionHachimiUnsafe(), txt.String(), func(msg *chat.Message) (bool, error) {
			return s.hachimiSelectionHandler(msg, desc, sessId)
		})

	case hachimi.ResponseStatusInspect:
		var txt strings.Builder
		txt.WriteString(i18n.T(keys.GatekeeperHachimiSuspicious))
		txt.WriteString(util.DescribeToHuman(req, toolProvider))
		txt.WriteString(i18n.Te(keys.GatekeeperHachimiReason, "", errors.New(resp.Reason)))

		if markType == ignorance.MarkHachimiAllSuspicious {
			return true, ErrHachimiSuspicious
		}
		return s.Select(ctx, instance, msg, buildSelectionHachimiSuspicious(), txt.String(), func(msg *chat.Message) (bool, error) {
			return s.hachimiSelectionHandler(msg, desc, sessId)
		})
	case hachimi.ResponseStatusSafe:
		s.ignoranceSessionManager.SetHachimiCache(desc, true)

		color.Blue("[Manboster Gatekeeper] Hachimi thinks it's safe to go!")

		return true, ErrHachimiSafe

	default:
		return false, fmt.Errorf("unexpected response type: %v", resp.Type)
	}
}

func (s *Service) hachimiSelectionHandler(msg *chat.Message, desc string, sid string) (bool, error) {
	cb := msg.SelectionCallback
	switch cb.SelectionValue {
	case "allow":
		s.ignoranceSessionManager.SetHachimiCache(desc, true)
		return true, ErrHachimiSafe
	case "deny":
		s.ignoranceSessionManager.SetHachimiCache(desc, false)
		return false, ErrHachimiDenied
	case "allow-suspicious":
		s.ignoranceSessionManager.SetHachimiCache(desc, false)
		s.ignoranceSessionManager.SetMark(sid, true, 60*60, ignorance.MarkHachimiAllSuspicious)
		return true, errors.New(i18n.T(keys.GateKeeperHachimiSuspiciousMsg))
	}
	return false, fmt.Errorf("invalid selection value: %s", cb.SelectionValue)
}
