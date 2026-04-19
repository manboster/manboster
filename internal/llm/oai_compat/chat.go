package oai_compat

import (
	"context"

	"github.com/manboster/manboster/internal/llm"
	"github.com/sashabaranov/go-openai"
)

// Chat allows you to chat with your model
func (s *Service) Chat(ctx context.Context, model string, messages []llm.Message) (*llm.Event, error) {
	apiMsgs := make([]openai.ChatCompletionMessage, 0, len(messages))
	for _, msg := range messages {
		if msg.Type&(llm.MessageText) != 0 {
			for _, part := range msg.Parts {
				switch part.PartsType {
				case llm.MessagePartsText:
					apiMsgs = append(apiMsgs, openai.ChatCompletionMessage{
						Role:    string(msg.Role),
						Content: part.Text.Text,
					})
				}
			}
		}
	}

	req := openai.ChatCompletionRequest{
		Model:       model,
		Messages:    apiMsgs,
		Temperature: 0.7,
	}

	// call it
	resp, err := s.cli.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, err
	}

	if len(resp.Choices) == 0 {
		return nil, ErrNoResponse
	}

	// fmt.Println(resp.Choices[0].Message)

	// then return its message
	return &llm.Event{
		EventType: llm.EventMessage | llm.EventUsage,
		Message: &llm.Message{
			Parts: []llm.MessageParts{
				{
					PartsType: llm.MessagePartsText,
					Text: &llm.MessageTextPayload{
						Text: resp.Choices[0].Message.Content,
					},
				},
			},
			Type: llm.MessageText,
			Role: llm.RoleAssistant,
		},
		Usage: &llm.Usage{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	}, nil
}

// ChatStream is the next generation WIP TODO:
func (s *Service) ChatStream(ctx context.Context, model string, messages []llm.Message) (<-chan *llm.Event, error) {
	msgChan := make(chan *llm.Event)
	return msgChan, nil
}
