package engine

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/llm"
)

func (e *Engine) HandleToolCall(ctx context.Context, instance chat.Provider, msg *chat.Message, event llm.Event) (llm.Event, bool, error) {
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
		resp, err := e.HandleToolExec(ctx, safeName, fmt.Sprintf("%v", req.ToolArgs))
		if err != nil {
			resp = err.Error()
		} else {
			successExecution = true
		}

		color.Blue(fmt.Sprintf("[Manboster Engine] Tool call %s responded with response: %q", safeName, resp))

		callMsg := msg.Clone()
		callMsg.MessageType = chat.MessageText
		callMsg.Reply = nil

		toolProvider, av := e.toolMaps[safeName]
		if !av {
			callMsg.Text = &chat.TextPayload{
				Text: fmt.Sprintf("Model called tool `%s` but not found.", safeName),
			}
		} else {
			if successExecution {
				callMsg.Text = &chat.TextPayload{
					Text: fmt.Sprintf("Model called tool `%s`(`%s`).", toolProvider.DisplayName(), safeName),
				}
			} else {
				callMsg.Text = &chat.TextPayload{
					Text: fmt.Sprintf("Model called tool `%s`(`%s`) but returned error: %q.", toolProvider.DisplayName(), safeName, err),
				}
			}
		}

		err = e.SendMessage(ctx, instance, callMsg)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Engine] Error sending message: %s", err))
		}

		respEvent.Message.ToolCallResponse = append(respEvent.Message.ToolCallResponse, llm.MessageToolCallResponsePayload{
			ID:       req.ID,
			ToolName: req.ToolName,
			Result:   resp,
		})
	}

	return respEvent, successExecution, nil
}
