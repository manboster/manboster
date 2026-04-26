package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

func (h *Handler) HandleToolCall(ctx context.Context, instance chat.Provider, msg *chat.Message, event llm.Event) (llm.Event, bool, error) {
	successExecution := false

	var respEvent llm.Event
	respEvent.EventType = llm.EventMessage
	respEvent.Message = &llm.Message{
		Role: llm.RoleToolCall,
		Type: llm.MessageToolCallResponse,
	}

	for _, req := range event.Message.ToolCallRequest {
		resp := ""
		safeName := strings.ReplaceAll(req.ToolName, "_", ".")
		resp, err := h.HandleToolExec(ctx, safeName, fmt.Sprintf("%v", req.ToolArgs))
		if err != nil {
			resp = err.Error()
		} else {
			successExecution = true
		}

		color.Blue(fmt.Sprintf("[Manboster Handler] Tool call %s responded with response: %q", safeName, resp))

		callMsg := msg.Clone()
		callMsg.MessageType = chat.MessageText
		callMsg.Reply = nil

		toolProvider, av := h.toolMaps[safeName]
		if !av {
			callMsg.Text = &chat.TextPayload{
				Text: fmt.Sprintf("Model called tool `%s` but not found.", safeName),
			}
		} else {
			if successExecution {
				txt := fmt.Sprintf("Model called tool `%s`(`%s`) ", toolProvider.DisplayName(), safeName)

				var result map[string]interface{}
				err := json.Unmarshal([]byte(fmt.Sprintf("%v", req.ToolArgs)), &result)
				if err != nil {
					color.Yellow("[Manboster Handler] Failed to unmarshal tool call result")
				}
				params := util.JSONParse(result)
				if params != "" {
					txt += fmt.Sprintf("with params: %s", params)
				}
				txt += "."

				callMsg.Text = &chat.TextPayload{
					Text: txt,
				}
			} else {
				callMsg.Text = &chat.TextPayload{
					Text: fmt.Sprintf("Model called tool `%s`(`%s`) but returned error: %q.", toolProvider.DisplayName(), safeName, err),
				}
			}
		}

		err = h.gateway.SendMessage(ctx, instance, callMsg)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Handler] Error sending message: %s", err))
		}

		respEvent.Message.ToolCallResponse = append(respEvent.Message.ToolCallResponse, llm.MessageToolCallResponsePayload{
			ID:       req.ID,
			ToolName: req.ToolName,
			Result:   resp,
		})
	}

	return respEvent, successExecution, nil
}
