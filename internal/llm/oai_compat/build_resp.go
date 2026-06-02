package oai_compat

import (
	"github.com/manboster/manboster/spec/llm"
	"github.com/sashabaranov/go-openai"
)

func (s *Service) buildResponse(resp openai.ChatCompletionResponse, model llm.Model) (*llm.Event, error) {
	if len(resp.Choices) == 0 {
		return nil, ErrNoResponse
	}
	m := resp.Choices[0].Message

	msg := &llm.Message{
		Role: llm.RoleAssistant,
	}

	// handle tool call request
	if len(m.ToolCalls) != 0 {
		msg.Type |= llm.MessageToolCallRequest
		for _, call := range m.ToolCalls {
			msg.ToolCallRequest = append(msg.ToolCallRequest, llm.MessageToolCallRequestPayload{
				ID:       call.ID,
				ToolName: call.Function.Name,
				ToolArgs: call.Function.Arguments,
			})
		}
	}

	// handle message text
	if m.Content != "" {
		msg.Type |= llm.MessageText
		msg.Parts = append(msg.Parts, llm.MessageParts{
			PartsType: llm.MessagePartsText,
			Text: &llm.MessageTextPayload{
				Text: m.Content,
			},
		})
	}

	if m.ReasoningContent != "" {
		msg.Type |= llm.MessageThinking
		msg.Thinking = &llm.MessageThinkingPayload{
			Thinking: m.ReasoningContent,
		}
	}

	// then return its message
	return &llm.Event{
		EventType: llm.EventMessage | llm.EventUsage,
		Model:     model.Name,
		Provider:  s.Name(),
		Message:   msg,
		Usage: &llm.Usage{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	}, nil
}
