package engine

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/llm"
)

func (e *Engine) LLMChat(ctx context.Context, pIndex int, mIndex int, msgList []llm.Message) (*llm.Event, error) {
	var err error = nil
	var event = &llm.Event{}
	// try 3 times
	times := 3
	// tries def
	tries := 1

	currentProvider := e.llmProviders[pIndex]
	currentModel := currentProvider.Models()[mIndex]

	for tries <= times {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		color.Blue(fmt.Sprintf("[Manboster Engine] Fetching message response from LLMProvider %q, model %q, try %d times", currentProvider.DisplayName(), currentModel.DisplayName, tries))
		// we make timeout requests.
		timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)

		event, err = currentProvider.Chat(timeoutCtx, currentModel.Name, msgList)

		cancel()

		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Failed to get message from LLMProvider %q, model %q after %d tries, get error: %q", currentProvider.DisplayName(), currentModel.DisplayName, tries, err))
			time.Sleep(time.Second * time.Duration(tries+1))
			tries++
		} else {
			color.Blue(fmt.Sprintf("[Manboster Engine] Got message feedback from LLMProvider %q, model %q", currentProvider.DisplayName(), currentModel.DisplayName))
			break
		}
	}
	if err != nil {
		return nil, err
	}
	return event, err
}
