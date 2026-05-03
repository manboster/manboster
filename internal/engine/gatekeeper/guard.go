package gatekeeper

import (
	"context"
	"encoding/json"
	"fmt"

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
		// run hachimi here...
		return true, nil
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

	selection := []chat.Selection{
		{
			Name:  "Continue",
			Value: "continue",
		},
		{
			Name:  "Continue and shut up, handled by hachimi",
			Value: "hachimi",
		},
		{
			Name:  "Cancel",
			Value: "cancel",
		},
		{
			Name:  "Cancel and silence in 15 minutes",
			Value: "cAnCel",
		},
	}
	selectMsg := msg.Clone()
	selectMsg.MessageType = chat.MessageSelection | chat.MessageText
	selectMsg.Selection = &chat.SelectionPayload{
		Selection:   selection,
		SelectionId: "",
	}
	selectMsg.Text = &chat.TextPayload{
		Text: txt,
	}
	resp, err := s.gatewayService.SendSelect(ctx, instance, selectMsg)
	if err != nil {
		color.Yellow("[Manboster Gatekeeper] Failed to get select result")
		return false, fmt.Errorf("failed to get select result: %v", err)
	}
	if resp.SelectionCallback != nil {
		id := fmt.Sprintf("%s:%s:%s:%s:%s", instance.Name(), resp.SelectionCallback.SelectionBy, sid, toolProvider.Name(), executeGroup)

		err = s.CheckSession(id)
		if err != nil {
			return false, err
		}
		if s.ignoranceSessionManager.GetIgnoreMark(id) {
			// run hachimi here...
			return true, nil
		}

		// get tool's min permission and compare it with current user's
		minPermission := types.UserTypeFromString(toolProvider.MetaData().MinUserType)
		uPermission := s.safeguardService.UserType(ctx, instance.Name(), resp.SelectionCallback.SelectionBy)
		if uPermission < minPermission {
			return false, fmt.Errorf("the permission user who performs the action is too low, please contact the owner")
		}

		// get resp based on
		switch resp.SelectionCallback.SelectionValue {
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
			s.ignoranceSessionManager.SetIgnoreMark(id, true, ttl)
			return true, nil
		case "continue":
			return true, nil
		case "cancel":
			return false, fmt.Errorf("user manually canceled your request")
		case "cAnCel":
			s.ignoranceSessionManager.SetCancelMark(id, true)
			return false, fmt.Errorf("user manually canceled your request")
		default:
			return false, fmt.Errorf("invalid selection value: %v", resp.SelectionCallback.SelectionValue)
		}
	}
	return false, fmt.Errorf("response do not contain available message params")
}
