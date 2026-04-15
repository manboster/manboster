package engine

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository/types"
)

func (e *Engine) mergeChatData(ctx context.Context, chatDataInfo []types.ChatData, sid string) error {
	s, avail := e.sessionManager.GetSession(sid)
	if avail && len(s.Events) > 0 {
		return nil
	}

	for _, info := range chatDataInfo {
		var event llm.Event
		if info.TotalTokens != 0 || info.PromptTokens != 0 || info.CompletionTokens != 0 {
			event.EventType |= llm.EventUsage
			event.Usage = &llm.Usage{
				TotalTokens:      info.TotalTokens,
				PromptTokens:     info.PromptTokens,
				CompletionTokens: info.CompletionTokens,
			}
		}

		event.Message = &llm.Message{
			Type: info.MessageType,
			Role: info.Role,
		}

		var msg llm.Message

		if info.MessageType&(llm.MessageText|llm.MessageImage|llm.MessageFile) != 0 {
			err := json.Unmarshal([]byte(info.MessagePayload), &msg)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] We encountered an error while reading chat data payload from repository, error: %q", err))
			}
			event.EventType |= llm.EventMessage
			event.Message.Parts = msg.Parts
		}

		if info.MessageType&llm.MessageToolCallRequest != 0 {
			err := json.Unmarshal([]byte(info.MessagePayload), &msg)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] We encountered an error while reading chat data payload from repository, error: %q", err))
			}
			event.EventType |= llm.EventMessage
			event.Message.ToolRequest = msg.ToolRequest
		}
		if info.MessageType&llm.MessageToolCallResponse != 0 {
			err := json.Unmarshal([]byte(info.MessagePayload), &msg)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] We encountered an error while reading chat data payload from repository, error: %q", err))
			}
			event.EventType |= llm.EventMessage
			event.Message.ToolResponse = msg.ToolResponse
		}

		e.sessionManager.AppendEvent(sid, event)
	}

	return nil
}
