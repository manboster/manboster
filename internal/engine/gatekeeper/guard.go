package gatekeeper

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

// HachimiGuard is core component of Manboster gatekeeper service.
func (s *Service) HachimiGuard(ctx context.Context, instance chat.Provider, msg *chat.Message, toolProvider tool.Provider, req llm.MessageToolCallRequestPayload) (bool, error) {
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
			Name:  "Don't disturb me in this session",
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
		switch resp.SelectionCallback.SelectionValue {
		case "hachimi":
		// TODO: hachimi automatically score
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
