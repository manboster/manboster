package gatekeeper

import (
	"context"
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
	overallUd := fmt.Sprintf("%s:%s:%s", instance.Name(), msg.ChatID, sid)
	ud := overallUd + fmt.Sprintf(":%s:%s", toolProvider.Name(), executeGroup)
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
	if m && mT == ignorance.MarkHachimiAll {
		// run hachimi here...
		return s.HachimiHandler(ctx, instance, msg, toolProvider, req, ud)
	}
	if m && mT == ignorance.MarkContinueAll {
		return true, nil
	}

	mark, markType := s.ignoranceSessionManager.GetMark(ud)
	if (mark && markType == ignorance.MarkHachimi && requireUserType <= actualUserType) || msg.MessageType&chat.MessageFromCron != 0 {
		return s.HachimiHandler(ctx, instance, msg, toolProvider, req, ud)
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
		cb := msg.SelectionCallback
		overallId := fmt.Sprintf("%s:%s:%s", instance.Name(), msg.ChatID, sid)
		id := overallId + fmt.Sprintf(":%s:%s", toolProvider.Name(), executeGroup)

		err = s.CheckSession(id)
		if err != nil {
			return false, err
		}

		m, mT := s.ignoranceSessionManager.GetMark(overallId)
		if m && mT == ignorance.MarkCancelAll {
			return false, fmt.Errorf("user manually canceled your all tool calls in next 10 minutes")
		}
		if m && mT == ignorance.MarkHachimiAll {
			// run hachimi here...
			return s.HachimiHandler(ctx, instance, msg, toolProvider, req, id)
		}
		if m && mT == ignorance.MarkContinueAll {
			return true, nil
		}

		mark, markType := s.ignoranceSessionManager.GetMark(id)
		if mark && markType == ignorance.MarkHachimi {
			// run hachimi here...
			return s.HachimiHandler(ctx, instance, msg, toolProvider, req, id)
		}
		if mark && markType == ignorance.MarkIgnore {
			color.Yellow("[Manboster] Ignored and updated data.")
			return true, nil
		}
		// get tool's min permission and compare it with current user's
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
			respMsg.Text.Text = fmt.Sprintf(i18n.T(keys.GatekeeperHachimiActivated), ttl/60)
			err := s.gatewayService.SendMessage(ctx, instance, respMsg)
			if err != nil {
				color.Yellow("[Manboster Gatekeeper] Failed to send hachimi prompt message")
			}
			s.ignoranceSessionManager.SetMark(id, true, ttl, ignorance.MarkHachimi)
			go s.Recall(ctx, instance, respMsg)
			return true, nil
		case guardSelectHachimiAll:
			respMsg.Text.Text = fmt.Sprintf(i18n.T(keys.GatekeeperHachimiActivated), ttl/60)
			err := s.gatewayService.SendMessage(ctx, instance, respMsg)
			if err != nil {
				color.Yellow("[Manboster Gatekeeper] Failed to send hachimi prompt message")
			}
			s.ignoranceSessionManager.SetMark(overallId, true, 60*60, ignorance.MarkHachimiAll)
			go s.Recall(ctx, instance, respMsg)
			return true, nil
		case guardSelectIgnore:
			respMsg.Text.Text = fmt.Sprintf(i18n.T(keys.GatekeeperShutUpMsg), ttl/60)
			err := s.gatewayService.SendMessage(ctx, instance, respMsg)
			if err != nil {
				color.Yellow("[Manboster Gatekeeper] Failed to send ignore prompt message")
			}
			s.ignoranceSessionManager.SetMark(id, true, ttl, ignorance.MarkIgnore)
			go s.Recall(ctx, instance, respMsg)
			return true, nil
		case guardSelectContinue:
			return true, nil
		case guardSelectContinueAll:
			respMsg.Text.Text = i18n.T(keys.GatekeeperContinueAllMsg)
			err := s.gatewayService.SendMessage(ctx, instance, respMsg)
			if err != nil {
				color.Yellow("[Manboster Gatekeeper] Failed to send ignore prompt message")
			}
			s.ignoranceSessionManager.SetMark(overallId, true, ttl, ignorance.MarkContinueAll)
			go s.Recall(ctx, instance, respMsg)
			return true, nil
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
			go s.Recall(ctx, instance, respMsg)
			return false, fmt.Errorf("user manually canceled your all tool calls in next 10 minutes")
		default:
			return false, fmt.Errorf("invalid selection value: %v", cb.SelectionValue)
		}
	})
}
