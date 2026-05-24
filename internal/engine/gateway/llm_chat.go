package gateway

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
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
	text := fmt.Sprintf(i18n.T(keys.GatewayLLMErrorDefault), llmProviderDisplayName, llmModelDisplayName, err)
	if strings.Contains(tips, "429") {
		text = fmt.Sprintf(i18n.T(keys.GatewayLLMErrorRateLimit), llmProviderDisplayName, llmModelDisplayName)
	} else if strings.Contains(tips, "500") || strings.Contains(tips, "502") || strings.Contains(tips, "503") || strings.Contains(tips, "501") {
		text = fmt.Sprintf(i18n.T(keys.GatewayLLMErrorDown), llmProviderDisplayName)
	} else if strings.Contains(tips, "context deadline exceeded") {
		text = fmt.Sprintf(i18n.T(keys.GatewayLLMErrorTimeout), llmProviderDisplayName)
	} else if strings.Contains(tips, "403") || strings.Contains(tips, "401") {
		text = fmt.Sprintf(i18n.T(keys.GatewayLLMErrorAuth), llmProviderDisplayName)
	} else if strings.Contains(tips, "cancel") {
		text = fmt.Sprintf(i18n.T(keys.GatewayLLMErrorCancelled), llmProviderDisplayName)
	}
	respMessage.Text = &chat.TextPayload{
		Text: text,
	}
	return s.SendMessage(ctx, instance, respMessage)
}
