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
	"github.com/manboster/manboster/internal/session/gatekeeper"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

func (s *Service) HachimiHandler(ctx context.Context, instance chat.Provider, mes *chat.Message, toolProvider tool.Provider, req llm.MessageToolCallRequestPayload, id string, sid string) (bool, error) {
	msg := mes.Clone()
	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{}

	sessId := BuildSessionId(instance.Name(), msg.ChatID, sid)
	_, markType := s.sessionService.Manager.Ignorance.GetMark(sessId)

	if !*s.hachimiLoaded || s.hachimiProvider == nil {
		color.Yellow("[Manboster Gatekeeper] Hachimi is not loaded!")
		return false, fmt.Errorf("hachimi is not loaded")
	}

	desc := util.DescribeToHachimi(req, toolProvider)
	u, avail := s.sessionService.Manager.Ignorance.GetHachimiCache(desc)
	if avail {
		if !u {
			return false, ErrHachimiDenied
		}
		return true, ErrHachimiSafe
	}

	resp, err := s.hachimiProvider.Chat(ctx, desc)
	if err != nil {
		return false, err
	}
	s.sessionService.Manager.Ignorance.UpdateMark(id)

	if resp == nil {
		resp = &hachimi.Response{
			Type:   hachimi.ResponseStatusUnsafe,
			Reason: "input message is too long so hachimi could not make decision for you",
		}
	}

	provider, model, _ := s.sessionService.Manager.ChatSession.GetModel(sid)
	p, m := util.GetModelWithFallback(ctx, s.llmProviders, provider, model)

	var descMsg strings.Builder
	if mes.Text != nil {
		descMsg.WriteString(fmt.Sprintf("user message: %s\n", *mes.Text))
	}
	descMsg.WriteString(util.DescribeToHachimi(req, toolProvider) + "\n")

	switch resp.Type {
	case hachimi.ResponseStatusUnsafe:
		var txt strings.Builder
		txt.WriteString(i18n.T(keys.GatekeeperHachimiUnsafe))
		txt.WriteString(util.DescribeToHuman(req, toolProvider))
		txt.WriteString(i18n.Te(keys.GatekeeperHachimiRemoteReason, "", errors.New(resp.Reason)))

		descMsg.WriteString(fmt.Sprintf("Hachimi thinks it's %s", resp.Type))
		res, chatErr := s.gatewayService.LLMQuickChat(ctx, p, m, buildTemplate(resp.Type), descMsg.String())
		if chatErr != nil {
			color.Yellow("[Manboster Gatekeeper] Failed to quick chat")
		}
		if res != nil && res.Message != nil && len(res.Message.Parts) != 0 && res.Message.Parts[0].Text != nil {
			txt.WriteString(i18n.Te(keys.GatekeeperHachimiRemoteReason, "", errors.New(res.Message.Parts[0].Text.Text)))
		}

		return s.Select(ctx, instance, msg, buildSelectionHachimiUnsafe(), txt.String(), func(msg *chat.Message) (bool, error) {
			return s.hachimiSelectionHandler(msg, desc, sessId)
		})

	case hachimi.ResponseStatusInspect:
		if markType == gatekeeper.MarkHachimiAllSuspicious {
			return true, ErrHachimiSuspicious
		}

		var txt strings.Builder
		txt.WriteString(i18n.T(keys.GatekeeperHachimiSuspicious))
		txt.WriteString(util.DescribeToHuman(req, toolProvider))
		txt.WriteString(i18n.Te(keys.GatekeeperHachimiReason, "", errors.New(resp.Reason)))

		descMsg.WriteString(fmt.Sprintf("Hachimi thinks it's %s", resp.Type))

		res, chatErr := s.gatewayService.LLMQuickChat(ctx, p, m, buildTemplate(resp.Type), descMsg.String())
		if chatErr != nil {
			color.Yellow("[Manboster Gatekeeper] Failed to quick chat")
		}
		if res != nil && res.Message != nil && len(res.Message.Parts) != 0 && res.Message.Parts[0].Text != nil {
			txt.WriteString(i18n.Te(keys.GatekeeperHachimiReason, "", errors.New(res.Message.Parts[0].Text.Text)))
		}

		return s.Select(ctx, instance, msg, buildSelectionHachimiSuspicious(), txt.String(), func(msg *chat.Message) (bool, error) {
			return s.hachimiSelectionHandler(msg, desc, sessId)
		})
	case hachimi.ResponseStatusSafe:
		s.sessionService.Manager.Ignorance.SetHachimiCache(desc, true)

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
		s.sessionService.Manager.Ignorance.SetHachimiCache(desc, true)
		return true, ErrHachimiSafe
	case "deny":
		s.sessionService.Manager.Ignorance.SetHachimiCache(desc, false)
		return false, ErrHachimiDenied
	case "allow-suspicious":
		s.sessionService.Manager.Ignorance.SetHachimiCache(desc, false)
		s.sessionService.Manager.Ignorance.SetMark(sid, true, 60*60, gatekeeper.MarkHachimiAllSuspicious)
		return true, errors.New(i18n.T(keys.GateKeeperHachimiSuspiciousMsg))
	}
	return false, fmt.Errorf("invalid selection value: %s", cb.SelectionValue)
}
