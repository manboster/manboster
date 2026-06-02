package oai_compat

import (
	"context"
	"strings"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/llm"
	"github.com/sashabaranov/go-openai"
)

// Chat allows you to chat with your model
func (s *Service) Chat(ctx context.Context, model llm.Model, tools []tool.Provider, messages []llm.Message) (*llm.Event, error) {
	apiMsgs := make([]openai.ChatCompletionMessage, 0, len(messages))

	for _, msg := range messages {
		apiMsgs = append(apiMsgs, s.buildRequest(msg, model)...)
	}

	req := openai.ChatCompletionRequest{
		Model:       model.Name,
		Messages:    apiMsgs,
		Temperature: 0.7,
		Tools:       s.ConvertTools(tools),
	}

	// tweaks because [Manboster Gateway] llmChat_kimi_kimi-k2.5_1780391259 failed on try 3, error: "error, status code: 400, status: 400 Bad Request, message: invalid temperature: only 1 is allowed for this model"
	if strings.Contains(strings.ToLower(model.Name), "kimi") {
		req.Temperature = 1
	}

	// call it
	resp, err := s.cli.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, err
	}

	return s.buildResponse(resp, model)
}

// ChatStream is the next generation WIP TODO:
func (s *Service) ChatStream(ctx context.Context, model llm.Model, tools []tool.Provider, messages []llm.Message) (<-chan *llm.Event, error) {
	msgChan := make(chan *llm.Event)
	return msgChan, nil
}
