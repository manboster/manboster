package openrouter

import (
	"context"

	"github.com/manboster/manboster/core/llm"
	"github.com/sashabaranov/go-openai"
)

// Chat allows you to chat with your model
func (s *Service) Chat(ctx context.Context, messages []llm.Message) (*llm.Message, error) {
	apiMsgs := make([]openai.ChatCompletionMessage, 0, len(messages))
	for _, msg := range messages {
		apiMsgs = append(apiMsgs, openai.ChatCompletionMessage{
			Role:    string(msg.Role),
			Content: msg.Text,
		})
	}

	req := openai.ChatCompletionRequest{
		Model:       s.cfg.Model,
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
	return &llm.Message{
		Text: resp.Choices[0].Message.Content,
		Type: llm.MessageTypeText,
		Role: llm.RoleTypeAssistant,
	}, nil
}
