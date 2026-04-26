package chatdata

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/spec/llm"
)

func (s *Service) Merge(ctx context.Context, chatDataInfo []types.ChatData, sid string) error {
	sess, avail := s.sessionManager.GetSession(sid)
	if avail && len(sess.Events) > 0 {
		return nil
	}

	for _, info := range chatDataInfo {
		var event llm.Event

		event.Model = info.Model
		event.Provider = info.Provider

		event.Message = &llm.Message{
			Type: info.MessageType,
			Role: info.Role,
		}

		var msg llm.Message

		if info.TotalTokens > 0 || info.PromptTokens > 0 || info.CompletionTokens > 0 || info.TotalCost > 0 || info.PromptTokens > 0 || info.CompletionTokens > 0 {
			event.EventType |= llm.EventUsage
			event.Usage = &llm.Usage{
				TotalTokens:      info.TotalTokens,
				PromptTokens:     info.PromptTokens,
				CompletionTokens: info.CompletionTokens,
				InputCost:        info.InputCost,
				OutputCost:       info.OutputCost,
				TotalCost:        info.TotalCost,
			}
		}

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
			event.Message.ToolCallRequest = msg.ToolCallRequest
		}
		if info.MessageType&llm.MessageToolCallResponse != 0 {
			err := json.Unmarshal([]byte(info.MessagePayload), &msg)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] We encountered an error while reading chat data payload from repository, error: %q", err))
			}
			event.EventType |= llm.EventMessage
			event.Message.ToolCallResponse = msg.ToolCallResponse
		}
		if info.MessageType&llm.MessageThinking != 0 {
			err := json.Unmarshal([]byte(info.MessagePayload), &msg)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] We encountered an error while reading chat data payload from repository, error: %q", err))
			}
			event.EventType |= llm.EventMessage
			event.Message.Thinking = &llm.MessageThinkingPayload{
				Thinking: msg.Thinking.Thinking,
			}
		}

		s.sessionManager.AppendEvent(sid, event)
	}

	return nil
}
