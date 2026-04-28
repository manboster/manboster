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
	ud := fmt.Sprintf("%s:%s:%s", instance.Name(), msg.UserID, sid)
	if s.ignoranceSessionManager.GetIgnoreMark(ud) && types.UserTypeFromString(toolProvider.MetaData().MinUserType) <= s.safeguardService.UserType(ctx, instance.Name(), msg.UserID) {
		// run hachimi here...
		return true, nil
	}
	txt := fmt.Sprintf("Model want to call tool `%s`(`%s`) ", toolProvider.DisplayName(), req.ToolName)
	var result map[string]interface{}
	err := json.Unmarshal([]byte(fmt.Sprintf("%v", req.ToolArgs)), &result)
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
			Name:  "Cancel",
			Value: "cancel",
		},
		{
			Name:  "Continue and don't disturb me in this session",
			Value: "hachimi",
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
		id := fmt.Sprintf("%s:%s:%s", instance.Name(), resp.SelectionCallback.SelectionBy, sid)
		// fmt.Println(id)
		// fmt.Println(s.ignoranceSessionManager.GetIgnoreMark(id))
		if s.ignoranceSessionManager.GetIgnoreMark(id) {
			// run hachimi here...
			return true, nil
		}

		minPermission := types.UserTypeFromString(toolProvider.MetaData().MinUserType)
		uPermission := s.safeguardService.UserType(ctx, instance.Name(), resp.SelectionCallback.SelectionBy)
		// fmt.Printf("%s %s", minPermission, uPermission)
		if uPermission < minPermission {
			return false, fmt.Errorf("user access denied")
		}

		switch resp.SelectionCallback.SelectionValue {
		case "hachimi":
			// TODO: hachimi automatically score
			s.ignoranceSessionManager.SetIgnoreMark(id, true)
			return true, nil
		case "continue":
			return true, nil
		case "cancel":
			return false, fmt.Errorf("user select cancel")
		default:
			return false, fmt.Errorf("invalid selection value: %v", resp.SelectionCallback.SelectionValue)
		}
	}
	return false, fmt.Errorf("response do not contain available message params")
}
