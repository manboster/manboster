package gatekeeper

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/repository/types"
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
	if s.ignoranceSessionManager.GetIgnoreMark(ud) && requireUserType <= actualUserType {
		return s.HachimiHandler(ctx, instance, msg, toolProvider, req, sid)
	}

	txt := fmt.Sprintf("Model wants to call tool `%s`(`%s`) ", toolProvider.DisplayName(), req.ToolName)
	var result map[string]interface{}
	err = json.Unmarshal([]byte(fmt.Sprintf("%v", req.ToolArgs)), &result)
	if err != nil {
		color.Yellow("[Manboster Handler] Failed to unmarshal tool call result")
	}
	params := util.JSONParse(result)
	if params != "" {
		txt += fmt.Sprintf("with params: %s", params)
	}
	txt += ", do you want to continue?"

	var selection []chat.Selection
	if s.hachimiConfig.Enabled {
		selection = selectionWithHachimi
	} else {
		selection = selectionNoHachimi
	}

	return s.Select(ctx, instance, msg, selection, txt, func(cb *chat.SelectionCallbackPayload) (bool, error) {
		id := fmt.Sprintf("%s:%s:%s:%s:%s", instance.Name(), cb.SelectionBy, sid, toolProvider.Name(), executeGroup)

		err = s.CheckSession(id)
		if err != nil {
			return false, err
		}
		if s.ignoranceSessionManager.GetIgnoreMark(id) {
			// run hachimi here...
			return s.HachimiHandler(ctx, instance, msg, toolProvider, req, sid)
		}

		// get tool's min permission and compare it with current user's
		minPermission := types.UserTypeFromString(toolProvider.MetaData().MinUserType)
		uPermission := s.safeguardService.UserType(ctx, instance.Name(), cb.SelectionBy)
		if uPermission < minPermission {
			return false, fmt.Errorf("the permission user who performs the action is too low, please contact the owner")
		}

		// get resp based on
		switch cb.SelectionValue {
		case "hachimi":
			ttl := 0
			// set TTL based on tools required user permission
			switch minPermission {
			case types.UserUnknown:
				ttl = 60 * 120 // 2 hours
			case types.UserAdmin:
				ttl = 60 * 60 // 1 hour
			case types.UserRoot:
				ttl = 60 * 30 // 30 minutes
			default:
			}

			if s.hachimiConfig.Enabled {
				respMsg := msg.Clone()
				respMsg.MessageType = chat.MessageText
				respMsg.Text = &chat.TextPayload{
					Text: "You activated hachimi, it will help you handle this tool in next " + strconv.Itoa(ttl/60) + " minutes, enjoy your time!",
				}
				respMsg.Reply = nil
				err := s.gatewayService.SendMessage(ctx, instance, respMsg)
				if err != nil {
					color.Yellow("[Manboster Gatekeeper] Failed to send hachimi prompt message")
				}
				s.ignoranceSessionManager.SetIgnoreMark(id, true, ttl)

				go func(instance chat.Provider, rMsg *chat.Message) {
					err := s.RecallRunner(ctx, instance, rMsg, 5*time.Second)
					if err != nil {
						color.Yellow("[Manboster Gatekeeper] Failed to recall result")
					}
				}(instance, respMsg)
			}
			return true, nil
		case "continue":
			return true, nil
		case "cancel":
			return false, fmt.Errorf("user manually canceled your request")
		case "cAnCel":
			s.ignoranceSessionManager.SetCancelMark(id, true)
			return false, fmt.Errorf("user manually canceled your request")
		default:
			return false, fmt.Errorf("invalid selection value: %v", cb.SelectionValue)
		}
	})
}
