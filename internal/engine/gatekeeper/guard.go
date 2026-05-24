package gatekeeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
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
	ud := fmt.Sprintf("%s:%s:%s:%s:%s", instance.Name(), msg.UserID, sid, toolProvider.Name(), executeGroup)
	err := s.CheckSession(ud)
	if err != nil {
		return false, err
	}

	requireUserType := types.UserTypeFromString(toolProvider.MetaData().MinUserType)
	actualUserType := s.safeguardService.UserType(ctx, instance.Name(), msg.UserID)

	mark, markType := s.ignoranceSessionManager.GetMark(ud)
	if (mark && markType == ignorance.MarkHachimi && requireUserType <= actualUserType) || msg.MessageType&chat.MessageFromCron != 0 {
		return s.HachimiHandler(ctx, instance, msg, toolProvider, req, ud)
	}
	if mark && markType == ignorance.MarkIgnore && requireUserType <= actualUserType || msg.MessageType&chat.MessageFromCronIgnore != 0 {
		return true, nil
	}

	var selection []chat.Selection
	if s.hachimiConfig.Enabled {
		selection = selectionWithHachimi
	} else {
		selection = selectionNoHachimi
	}

	return s.Select(ctx, instance, msg, selection, util.DescribeToHuman(req, toolProvider), func(cb *chat.SelectionCallbackPayload) (bool, error) {
		id := fmt.Sprintf("%s:%s:%s:%s:%s", instance.Name(), cb.SelectionBy, sid, toolProvider.Name(), executeGroup)

		err = s.CheckSession(id)
		if err != nil {
			return false, err
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
			ttl = 60 * 60 // 1 hour
		case types.UserAdmin:
			ttl = 60 * 30 // 30 minutes
		case types.UserRoot:
			ttl = 60 * 15 // 15 minutes
		default:
		}

		// get resp based on
		switch guardSelectType(cb.SelectionValue) {
		case guardSelectHachimi:
			respMsg := msg.Clone()
			respMsg.MessageType = chat.MessageText
			respMsg.Text = &chat.TextPayload{
				Text: "🐱 You activated hachimi, it will help you handle this tool in next " + strconv.Itoa(ttl/60) + " minutes, enjoy your time!",
			}
			respMsg.Reply = nil
			err := s.gatewayService.SendMessage(ctx, instance, respMsg)
			if err != nil {
				color.Yellow("[Manboster Gatekeeper] Failed to send hachimi prompt message")
			}
			s.ignoranceSessionManager.SetMark(id, true, ttl, ignorance.MarkHachimi)

			go func(instance chat.Provider, rMsg *chat.Message) {
				err := s.RecallRunner(ctx, instance, rMsg, 5*time.Second)
				if err != nil {
					color.Yellow("[Manboster Gatekeeper] Failed to recall result")
				}
			}(instance, respMsg)
			return true, nil
		case guardSelectIgnore:
			respMsg := msg.Clone()
			respMsg.MessageType = chat.MessageText
			respMsg.Text = &chat.TextPayload{
				Text: "⚠️ You ignored this tool call, it will automatically allow in next " + strconv.Itoa(ttl/60) + " minutes.",
			}
			respMsg.Reply = nil
			err := s.gatewayService.SendMessage(ctx, instance, respMsg)
			if err != nil {
				color.Yellow("[Manboster Gatekeeper] Failed to send ignore prompt message")
			}
			s.ignoranceSessionManager.SetMark(id, true, ttl, ignorance.MarkIgnore)

			go func(instance chat.Provider, rMsg *chat.Message) {
				err := s.RecallRunner(ctx, instance, rMsg, 5*time.Second)
				if err != nil {
					color.Yellow("[Manboster Gatekeeper] Failed to recall result")
				}
			}(instance, respMsg)
			return true, nil
		case guardSelectContinue:
			return true, nil
		case guardSelectCancel:
			return false, fmt.Errorf("user manually canceled your request")
		case guardSelectCancelIgnore:
			s.ignoranceSessionManager.SetMark(id, true, 15*60, ignorance.MarkCancel)
			return false, fmt.Errorf("user manually canceled your request")
		default:
			return false, fmt.Errorf("invalid selection value: %v", cb.SelectionValue)
		}
	})
}
