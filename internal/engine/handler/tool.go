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

func (h *Handler) HandleToolCall(ctx context.Context, instance chat.Provider, msg *chat.Message, event llm.Event, sid string) (llm.Event, bool, error) {
	successExecution := false

	var respEvent llm.Event
	respEvent.EventType = llm.EventMessage
	respEvent.Message = &llm.Message{
		Role: llm.RoleToolCall,
		Type: llm.MessageToolCallResponse,
	}

	toolCallReqCount := 0
	toolCallMsgId := ""
	var txt strings.Builder

	for _, req := range event.Message.ToolCallRequest {
		originalName := req.ToolName
		safeName := strings.ReplaceAll(req.ToolName, "_", ".")
		req.ToolName = safeName

		callMsg := msg.Clone()
		callMsg.MessageType = chat.MessageText
		callMsg.Reply = nil

		toolProvider, avail := h.toolMaps[req.ToolName]
		if !avail {
			color.Red(fmt.Sprintf("[Manboster Engine] There is no tool named %q", req.ToolName))
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
		isOK, err := h.gatekeeperService.Guard(ctx, instance, msg, toolProvider, req, sid)
		if !isOK {
			color.Red(fmt.Sprintf("[Manboster Handler] Gatekeeper Rejected the tool call `%s`: %q", req.ToolName, err))
			callMsg.Text = &chat.TextPayload{
				Text: fmt.Sprintf("Gatekeeper Rejected the tool call `%s`: %q", req.ToolName, err.Error()),
			}
			e := h.gateway.SendMessage(ctx, instance, callMsg)
			if e != nil {
				color.Yellow(fmt.Sprintf("[Manboster Handler] Error sending message: %s", err))
			}

			respEvent.Message.ToolCallResponse = append(respEvent.Message.ToolCallResponse, llm.MessageToolCallResponsePayload{
				ID:       req.ID,
				ToolName: req.ToolName,
				Result:   fmt.Sprintf("Gatekeeper Rejected the tool call `%s`: %q", req.ToolName, err),
			})
			continue
		}

		resp := ""
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

		if toolCallReqCount%5 == 0 {
			err = h.gateway.SendMessage(ctx, instance, callMsg)
			toolCallMsgId = callMsg.MessageID
			if toolCallReqCount != 0 {
				txt.Reset()
			}
		} else {
			cMsg := callMsg.Clone()
			cMsg.MessageID = toolCallMsgId
			cMsg.MessageType = callMsg.MessageType | chat.MessageUnknown
			cMsg.Text = &chat.TextPayload{
				Text: txt.String(),
			}
			err = h.gateway.EditMessage(ctx, instance, cMsg)
		}

		toolCallReqCount++

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
