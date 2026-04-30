package oai_compat

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/llm"
	"github.com/sashabaranov/go-openai"
)

// Chat allows you to chat with your model
func (s *Service) Chat(ctx context.Context, model string, tools []tool.Provider, messages []llm.Message) (*llm.Event, error) {
	apiMsgs := make([]openai.ChatCompletionMessage, 0, len(messages))

	//j, _ := json.MarshalIndent(messages, "", "  ")
	//fmt.Println(string(j))

	for _, msg := range messages {
		if msg.Type&(llm.MessageToolCallResponse) != 0 {
			// going to check tool call resp and get it back to llm!
			for _, resp := range msg.ToolCallResponse {
				apiMsgs = append(apiMsgs, openai.ChatCompletionMessage{
					Role:       openai.ChatMessageRoleTool,
					Content:    resp.Result,
					ToolCallID: resp.ID,
					Name:       resp.ToolName,
				})
			}
		} else {
			ccMsg := openai.ChatCompletionMessage{
				Role: string(msg.Role),
			}

			if msg.Role == llm.RoleSystem {
				ccMsg.Content = msg.Parts[0].Text.Text
				apiMsgs = append(apiMsgs, ccMsg)
				continue
			}

			// get there is a request or not
			if msg.Type&(llm.MessageToolCallRequest) != 0 {
				// fmt.Printf("run!")
				for _, req := range msg.ToolCallRequest {
					ccMsg.ToolCalls = append(ccMsg.ToolCalls, openai.ToolCall{
						ID:   req.ID,
						Type: openai.ToolTypeFunction,
						Function: openai.FunctionCall{
							Name:      req.ToolName,
							Arguments: req.ToolArgs.(string),
						},
					})
				}
			}

			// get there is a reasoning or not
			if msg.Type&(llm.MessageThinking) != 0 && msg.Thinking != nil {
				ccMsg.ReasoningContent = msg.Thinking.Thinking
			}

			if msg.Type&(llm.MessageText|llm.MessageFile|llm.MessageImage) != 0 {
				for _, part := range msg.Parts {
					switch part.PartsType {
					case llm.MessagePartsText:
						ccMsg.Content += part.Text.Text + "\n"
					}
				}
			}
			apiMsgs = append(apiMsgs, ccMsg)
		}

	}

	//jsonify, _ := json.MarshalIndent(apiMsgs, "", " ")
	//fmt.Println(string(jsonify))

	req := openai.ChatCompletionRequest{
		Model:       model,
		Messages:    apiMsgs,
		Temperature: 0.7,
		Tools:       s.ConvertTools(tools),
	}

	// call it
	resp, err := s.cli.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return nil, ErrNoResponse
	}
	m := resp.Choices[0].Message

	jsonify, _ := json.Marshal(m)
	fmt.Println(string(jsonify))

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
		Model:     model,
		Provider:  s.Name(),
		Message:   msg,
		Usage: &llm.Usage{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	}, nil
}

// ChatStream is the next generation WIP TODO:
func (s *Service) ChatStream(ctx context.Context, model string, tools []tool.Provider, messages []llm.Message) (<-chan *llm.Event, error) {
	msgChan := make(chan *llm.Event)
	return msgChan, nil
}
