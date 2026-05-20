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
		var txt strings.Builder
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

		intf, err := util.FromPayloadToInterface(req)
		if err != nil {
			color.Yellow("[Manboster Handler] Failed to convert payload to interface for %s: %s", req.ToolName, err)
		}
		err = schema.Validate(intf, *toolProvider.Args())
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Handler] Validate `%s` failed: %q", req.ToolName, err))
			result := fmt.Sprintf("Gatekeeper Rejected `%s`, params validate failed: %q", req.ToolName, err.Error())
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

		isOK, err := h.gatekeeperService.Guard(ctx, instance, msg, toolProvider, req, sid)
		if !isOK {
			color.Red(fmt.Sprintf("[Manboster Handler] Gatekeeper Rejected the tool call `%s`: %q", req.ToolName, err))
			result := fmt.Sprintf("Gatekeeper Rejected the tool call `%s`: %q", req.ToolName, err)

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
		valueCtx := context.WithValue(ctx, "chat_id", msg.ChatID)
		valueCtx = context.WithValue(valueCtx, "user_id", msg.UserID)
		valueCtx = context.WithValue(valueCtx, "chat_provider", instance.Name())
		valueCtx = context.WithValue(valueCtx, "session_id", sid)
		valueCtx = context.WithValue(valueCtx, "user_type", h.safeguardService.UserType(ctx, instance.Name(), msg.UserID).String())

		resp, err = h.HandleToolExec(valueCtx, toolProvider, fmt.Sprintf("%v", req.ToolArgs))
		if err != nil {
			resp = err.Error()
		} else {
			successExecution = true
		}
		color.Blue(fmt.Sprintf("[Manboster Handler] Tool call %s responded successfully.", safeName)) // with response resp.

		if successExecution {
			txt.WriteString(fmt.Sprintf("Model called `%s`(`%s`) ", toolProvider.DisplayName(), safeName))

			var result map[string]interface{}
			err := json.Unmarshal([]byte(fmt.Sprintf("%v", req.ToolArgs)), &result)
			if err != nil {
				color.Yellow("[Manboster Handler] Failed to unmarshal tool call result")
			}
			params := util.JSONParse(result)
			if params != "" {
				txt.WriteString(fmt.Sprintf("with params: %s", params))
			}
			txt.WriteString(".\n")

			callMsg.Text = &chat.TextPayload{
				Text: txt.String(),
			}
		} else {
			callMsg.Text = &chat.TextPayload{
				Text: fmt.Sprintf("Model called `%s`(`%s`) but returned error: %q.", toolProvider.DisplayName(), safeName, err),
			}
		}

		if *count%5 == 0 {
			err = h.gateway.SendMessage(ctx, instance, callMsg)
			*msgId = callMsg.MessageID
			if *count != 0 {
				*toolCallMsg = ""
			}
		} else {
			*toolCallMsg = *toolCallMsg + "\n" + txt.String()
			cMsg := callMsg.Clone()
			cMsg.MessageID = *msgId
			cMsg.MessageType = callMsg.MessageType | chat.MessageUnknown
			cMsg.Text = &chat.TextPayload{
				Text: *toolCallMsg,
			}
			err = h.gateway.EditMessage(ctx, instance, cMsg)
		}

		*count++

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
