package gateway

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

func (s *Service) LLMChat(ctx context.Context, currentProvider llm.Provider, currentModel llm.Model, msgList []llm.Message) (*llm.Event, error) {
	var event = &llm.Event{}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	name := "llmChat_" + currentProvider.Name() + "_" + currentModel.Name + "_" + strconv.FormatInt(time.Now().Unix(), 10)
	err := withRetry(ctx, name, 3, func(ctx context.Context) error {
		color.Blue(fmt.Sprintf("[Manboster Gateway] Fetching message response from LLMProvider %q, model %q", currentProvider.DisplayName(), currentModel.DisplayName))
		// we make timeout requests.
		timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
		defer cancel()

		var errChat error = nil
		event, errChat = currentProvider.Chat(timeoutCtx, currentModel.Name, s.toolProviders, msgList)

		return errChat
	})
	if err != nil {
		return nil, err
	}
	return event, err
}

func (s *Service) HandleLLMChatError(ctx context.Context, instance chat.Provider, msg *chat.Message, pName string, mName string, err error) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	if err == nil {
		return nil
	}
	llmProviderDisplayName := pName
	llmModelDisplayName := mName

	color.Red(fmt.Sprintf("[Manboster Gateway] Failed to get message from LLMProvider %q, model %q, get error: %q", llmProviderDisplayName, llmModelDisplayName, err))
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
	return s.SendMessage(ctx, instance, respMessage)
}
