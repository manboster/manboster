package engine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/spec/chat"
)

func (e *Engine) LLMChat(ctx context.Context, p llm.Provider, m llm.Model, msgList []llm.Message) (*llm.Event, error) {
	var err error = nil
	var event = &llm.Event{}
	// try 3 times
	times := 3
	// tries def
	tries := 1

	currentProvider := p
	currentModel := m

	for tries <= times {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		color.Blue(fmt.Sprintf("[Manboster Engine] Fetching message response from LLMProvider %q, model %q, try %d times", currentProvider.DisplayName(), currentModel.DisplayName, tries))
		// we make timeout requests.
		timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)

		event, err = currentProvider.Chat(timeoutCtx, currentModel.Name, e.toolProviders, msgList)

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

func (e *Engine) HandleLLMChatError(ctx context.Context, instance chat.Provider, msg *chat.Message, pName string, mName string, err error) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	if err == nil {
		return nil
	}
	llmProviderDisplayName := pName
	llmModelDisplayName := mName

	color.Red(fmt.Sprintf("[Manboster Engine] Failed to get message from LLMProvider %q, model %q, get error: %q", llmProviderDisplayName, llmModelDisplayName, err))
	// now we have to wrap this into friendly prompt
	tips := fmt.Sprintf("%+v", err)
	text := fmt.Sprintf("[Manboster] Failed to get message from LLMProvider %q, model %q, get error: %s\nYou can resend your message or check the API's availability.", llmProviderDisplayName, llmModelDisplayName, err)
	if strings.Contains(tips, "429") {
		text = fmt.Sprintf("[Manboster] Provider %s, Model %s has been suffering a very high traffic and triggered rate limit, please try again later or change provider's models.", llmProviderDisplayName, llmModelDisplayName)
	} else if strings.Contains(tips, "500") || strings.Contains(tips, "502") || strings.Contains(tips, "503") || strings.Contains(tips, "501") {
		text = fmt.Sprintf("[Manboster] Provider %s has been down, please check your provider's status page, or change providers and try again later.", llmProviderDisplayName)
	} else if strings.Contains(tips, "context deadline exceeded") {
		text = fmt.Sprintf("[Manboster] It seems that there is a connection issue between you and provider %s, please check your internet connection and try again.", llmProviderDisplayName)
	} else if strings.Contains(tips, "403") || strings.Contains(tips, "401") {
		text = fmt.Sprintf("[Manboster] Access denied or unauthorized in provider %s, please check your API key or other credentials is valid.", llmProviderDisplayName)
	} else if strings.Contains(tips, "cancel") {
		text = fmt.Sprintf("[Manboster] You cancelled provider %s's request.", llmProviderDisplayName)
	}
	respMessage.Text = &chat.TextPayload{
		Text: text,
	}
	return e.SendMessage(ctx, instance, respMessage)
}
