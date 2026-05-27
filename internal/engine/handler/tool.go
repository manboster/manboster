package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
	"github.com/manboster/manboster/spec/schema"
)

func (h *Handler) HandleToolCall(ctx context.Context, instance chat.Provider, msg *chat.Message, event llm.Event, sid string, count *int, msgId *string, toolCallMsg *string) (llm.Event, bool, error) {
	successExecution := false

	var respEvent llm.Event
	respEvent.EventType = llm.EventMessage
	respEvent.Message = &llm.Message{
		Role: llm.RoleToolCall,
		Type: llm.MessageToolCallResponse,
	}

	for _, req := range event.Message.ToolCallRequest {
		originalName := req.ToolName
		safeName := strings.ReplaceAll(req.ToolName, "_", ".")
		req.ToolName = safeName

		callMsg := msg.Clone()
		callMsg.MessageType = chat.MessageText
		callMsg.Reply = nil

		toolProvider, avail := h.toolMaps[req.ToolName]

		if !avail {
			color.Red(fmt.Sprintf("[Manboster Handler] There is no tool named %q", req.ToolName))
			callMsg.Text = &chat.TextPayload{
				Text: fmt.Sprintf("Model called tool `%s` but not found.", safeName),
			}
			err := h.gateway.SendMessage(ctx, instance, callMsg)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Handler] Error sending message: %s", err))
			}

			respEvent.Message.ToolCallResponse = append(respEvent.Message.ToolCallResponse, llm.MessageToolCallResponsePayload{
				ID:       req.ID,
				ToolName: req.ToolName,
				Result:   fmt.Sprintf("there is no tool named %q", originalName),
			})
			continue
		}

		if toolProvider.Args() != nil {
			intf, err := util.FromPayloadToInterface(req)
			if err != nil {
				color.Yellow("[Manboster Handler] Failed to convert payload to interface for %s: %s", req.ToolName, err)
			}
			err = schema.Validate(intf, *toolProvider.Args())
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Handler] Validate `%s` failed: %q", req.ToolName, err))
				result := i18n.Te(keys.GatekeeperValidateRejectMsg, req.ToolName, err)
				callMsg.Text = &chat.TextPayload{
					Text: result,
				}
				e := h.gateway.SendMessage(ctx, instance, callMsg)
				if e != nil {
					color.Yellow(fmt.Sprintf("[Manboster Handler] Error sending message: %s", err))
				}

				respEvent.Message.ToolCallResponse = append(respEvent.Message.ToolCallResponse, llm.MessageToolCallResponsePayload{
					ID:       req.ID,
					ToolName: req.ToolName,
					Result:   result,
				})
				continue
			}
		}

		isOK, err := h.gatekeeperService.Guard(ctx, instance, msg, toolProvider, req, sid)
		if !isOK {
			color.Red(fmt.Sprintf("[Manboster Handler] Gatekeeper Rejected the tool call `%s`: %q", req.ToolName, err))
			result := i18n.Te(keys.GatekeeperRejectMsg, req.ToolName, err)

			callMsg.Text = &chat.TextPayload{
				Text: result,
			}
			e := h.gateway.SendMessage(ctx, instance, callMsg)
			if e != nil {
				color.Yellow(fmt.Sprintf("[Manboster Handler] Error sending message: %s", err))
			}

			respEvent.Message.ToolCallResponse = append(respEvent.Message.ToolCallResponse, llm.MessageToolCallResponsePayload{
				ID:       req.ID,
				ToolName: req.ToolName,
				Result:   result,
			})
			continue
		}

		resp := ""
		// bring things passthrough to tools
		resp, err = h.HandleToolExec(h.PassthroughContextValues(ctx, instance, msg, sid), toolProvider, fmt.Sprintf("%v", req.ToolArgs))
		if err != nil {
			resp = err.Error()
		} else {
			successExecution = true
		}
		color.Blue(fmt.Sprintf("[Manboster Handler] Tool call %s responded successfully.", safeName)) // with response resp.

		err = h.DistributeFeedbackMsg(ctx, instance, msg, sid, toolProvider, req, err)
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
