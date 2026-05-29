package gatekeeper

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/internal/session/ignorance"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

// Guard is core component of Manboster gatekeeper service.
func (s *Service) Guard(ctx context.Context, instance chat.Provider, msg *chat.Message, toolProvider tool.Provider, req llm.MessageToolCallRequestPayload, sid string) (bool, error) {
	executeGroup := toolProvider.CacheGroup(fmt.Sprintf("%s", req.ToolArgs))

	overallUd := BuildSessionId(instance.Name(), msg.ChatID, sid)
	ud := buildToolId(overallUd, toolProvider.Name(), executeGroup)
	err := s.CheckSession(ud)
	if err != nil {
		return false, err
	}

	requireUserType := types.UserTypeFromString(toolProvider.MetaData().MinUserType)
	actualUserType := s.safeguardService.UserType(ctx, instance.Name(), msg.UserID)

	m, mT := s.ignoranceSessionManager.GetMark(overallUd)
	if m && mT == ignorance.MarkCancelAll {
		return false, fmt.Errorf("user manually canceled your all tool calls in next 10 minutes")
	}
	if m && (mT == ignorance.MarkHachimiAll || mT == ignorance.MarkHachimiAllSuspicious) {
		// run hachimi here...
		return s.HachimiHandler(ctx, instance, msg, toolProvider, req, ud, sid)
	}
	if m && mT == ignorance.MarkContinueAll {
		return true, nil
	}

	mark, markType := s.ignoranceSessionManager.GetMark(ud)
	if (mark && markType == ignorance.MarkHachimi && requireUserType <= actualUserType) || msg.MessageType&chat.MessageFromCron != 0 {
		return s.HachimiHandler(ctx, instance, msg, toolProvider, req, ud, sid)
	}
	if mark && markType == ignorance.MarkIgnore && requireUserType <= actualUserType || msg.MessageType&chat.MessageFromCronIgnore != 0 {
		return true, nil
	}

	var selection []chat.Selection
	if s.hachimiConfig.Enabled {
		selection = buildSelectionWithHachimi()
	} else {
		selection = buildSelectionNoHachimi()
	}

	return s.Select(ctx, instance, msg, selection, util.DescribeToHuman(req, toolProvider), func(msg *chat.Message) (bool, error) {
		overallId := BuildSessionId(instance.Name(), msg.ChatID, sid)
		id := buildToolId(overallId, toolProvider.Name(), executeGroup)

		err = s.CheckSession(id)
		if err != nil {
			return false, err
		}

		// get by session id
		m, mT := s.ignoranceSessionManager.GetMark(overallId)
		if m && mT == ignorance.MarkCancelAll {
			return false, fmt.Errorf("user manually canceled your all tool calls in next 10 minutes")
		}
		if m && (mT == ignorance.MarkHachimiAll || mT == ignorance.MarkHachimiAllSuspicious) {
			// run hachimi here...
			return s.HachimiHandler(ctx, instance, msg, toolProvider, req, id, sid)
		}
		if m && mT == ignorance.MarkContinueAll {
			return true, nil
		}

		// get by tool id and group
		mark, markType := s.ignoranceSessionManager.GetMark(id)
		if mark && markType == ignorance.MarkHachimi {
			// run hachimi here...
			return s.HachimiHandler(ctx, instance, msg, toolProvider, req, id, sid)
		}
		if mark && markType == ignorance.MarkIgnore {
			color.Yellow("[Manboster] Ignored and updated data.")
			return true, nil
		}

		// get tool's min permission and compare it with current user's
		cb := msg.SelectionCallback
		minPermission := types.UserTypeFromString(toolProvider.MetaData().MinUserType)
		uPermission := s.safeguardService.UserType(ctx, instance.Name(), cb.SelectionBy)
		if uPermission < minPermission {
			return false, fmt.Errorf("the permission user who performs the action is too low, please contact the owner")
		}

		ttl := 0
		// set TTL based on tools required user permission
		switch minPermission {
		case types.UserUnknown:
			ttl = 60 * 240 // 4 hours
		case types.UserAdmin:
			ttl = 60 * 120 // 2 hours
		case types.UserRoot:
			ttl = 60 * 30 // 30 minutes
		default:
		}

		respMsg := msg.Clone()
		respMsg.MessageType = chat.MessageText
		respMsg.Text = &chat.TextPayload{}
		respMsg.Reply = nil

		// get resp based on
		switch guardSelectType(cb.SelectionValue) {
		case guardSelectHachimi:
			s.ignoranceSessionManager.SetMark(id, true, ttl, ignorance.MarkHachimi)
			return true, errors.New(i18n.T(keys.GatekeeperHachimiActivated, map[string]any{
				"Next": ttl / 60,
			}))
		case guardSelectHachimiAll:
			s.ignoranceSessionManager.SetMark(overallId, true, 60*60, ignorance.MarkHachimiAll)
			return true, errors.New(i18n.T(keys.GatekeeperHachimiActivated, map[string]any{
				"Next": 60,
			}))
		case guardSelectIgnore:
			s.ignoranceSessionManager.SetMark(id, true, ttl, ignorance.MarkIgnore)
			return true, errors.New(i18n.T(keys.GatekeeperShutUpMsg, map[string]any{
				"Next": ttl / 60,
				"Name": toolProvider.DisplayName(),
			}))
		case guardSelectContinue:
			return true, nil
		case guardSelectContinueAll:
			s.ignoranceSessionManager.SetMark(overallId, true, ttl, ignorance.MarkContinueAll)
			return true, errors.New(i18n.T(keys.GatekeeperContinueAllMsg))
		case guardSelectCancel:
			return false, fmt.Errorf("user manually canceled your request")
		case guardSelectCancelIgnore:
			s.ignoranceSessionManager.SetMark(id, true, 15*60, ignorance.MarkCancel)
			return false, fmt.Errorf("user manually canceled your request for this tool call for 15 minutes")
		case guardSelectCancelAll:
			respMsg.Text.Text = i18n.T(keys.GatekeeperCancelAllMsg)
			err := s.gatewayService.SendMessage(ctx, instance, respMsg)
			if err != nil {
				color.Yellow("[Manboster Gatekeeper] Failed to send ignore prompt message")
			}
			s.ignoranceSessionManager.SetMark(overallId, true, ttl, ignorance.MarkCancelAll)
			go s.gatewayService.Recall(ctx, instance, respMsg)
			return false, fmt.Errorf("user manually canceled your all tool calls in next 10 minutes")
		default:
			return false, fmt.Errorf("invalid selection value: %v", cb.SelectionValue)
		}
	})
}
