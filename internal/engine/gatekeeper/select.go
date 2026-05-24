package gatekeeper

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

func (s *Service) Select(ctx context.Context, instance chat.Provider, msg *chat.Message, selection []chat.Selection, prompt string, callback func(msg *chat.Message) (bool, error)) (bool, error) {
	selectMsg := msg.Clone()
	selectMsg.MessageType = chat.MessageSelection | chat.MessageText
	selectMsg.Selection = &chat.SelectionPayload{
		Selection:   selection,
		SelectionId: "",
	}
	selectMsg.Text = &chat.TextPayload{
		Text: prompt,
	}
	resp, err := s.gatewayService.SendSelect(ctx, instance, selectMsg)
	if err != nil {
		color.Yellow("[Manboster Gatekeeper] Failed to get select result")
		return false, fmt.Errorf("failed to get select result: %v", err)
	}
	if resp.SelectionCallback == nil {
		return false, fmt.Errorf("failed to get selection callback")
	}

	return callback(resp)
}
