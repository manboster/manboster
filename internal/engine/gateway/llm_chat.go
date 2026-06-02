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

func (s *Service) LLMChat(ctx context.Context, currentProvider llm.Provider, currentModel llm.Model, msgList []llm.Message, ch chan int) (*llm.Event, error) {
	var event = &llm.Event{}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	name := "llmChat_" + currentProvider.Name() + "_" + currentModel.Name + "_" + strconv.FormatInt(time.Now().Unix(), 10)
	err := withRetry(ctx, name, 5, func(ctx context.Context) error {
		color.Blue(fmt.Sprintf("[Manboster Gateway] Fetching message response from LLMProvider %q, model %q", currentProvider.DisplayName(), currentModel.DisplayName))
		// we make timeout requests.
		timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
		defer cancel()

		var errChat error = nil
		event, errChat = currentProvider.Chat(timeoutCtx, currentModel, s.toolProviders, msgList)

		return errChat
	}, ch)

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
	text := fmt.Sprintf(i18n.T(keys.GatewayLLMErrorDefault, map[string]any{
		"Name":  llmModelDisplayName,
		"Model": llmProviderDisplayName,
		"Error": tips,
	}))
	if strings.Contains(tips, "429") {
		text = i18n.T(keys.GatewayLLMErrorRateLimit, map[string]any{
			"Name":  llmProviderDisplayName,
			"Model": llmModelDisplayName,
		})
	} else if strings.Contains(tips, "500") || strings.Contains(tips, "502") || strings.Contains(tips, "503") || strings.Contains(tips, "501") {
		text = i18n.Te(keys.GatewayLLMErrorDown, llmProviderDisplayName, nil)
	} else if strings.Contains(tips, "context deadline exceeded") {
		text = i18n.Te(keys.GatewayLLMErrorTimeout, llmProviderDisplayName, nil)
	} else if strings.Contains(tips, "403") || strings.Contains(tips, "401") {
		text = i18n.Te(keys.GatewayLLMErrorAuth, llmProviderDisplayName, nil)
	} else if strings.Contains(tips, "cancel") {
		text = i18n.Te(keys.GatewayLLMErrorCancelled, llmProviderDisplayName, nil)
	}
	text += "\n" + i18n.T(keys.GatewayRetryPrompt)
	respMessage.Text = &chat.TextPayload{
		Text: text,
	}
	return s.SendMessage(ctx, instance, respMessage)
}
