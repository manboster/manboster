package engine

import (
	"context"
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/util"
)

// cmdCancel enables user to cancel their request
func (e *Engine) cmdCancel(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	msg.MessageType = chat.MessageText
	sessData, avail := e.sessionManager.GetSession(sessionId)

	var text string
	if avail {
		if sessData.Active {
			sessData.Cancel()
			text = "[Manboster] Successfully cancelled the request."
		} else {
			text = "[Manboster] The request in this session is not active."
		}
	} else {
		text = "[Manboster] The request in this session is not active."
	}

	msg.Text = &chat.TextPayload{
		Text: text,
	}

	return instance.SendMessage(ctx, msg)
}

func (e *Engine) cmdNew(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	_, avail := e.sessionManager.GetSession(sessionId)
	if !avail {
		respMessage.Text = &chat.TextPayload{
			Text: "Session is not active, there is nothing to do!",
		}
		return instance.SendMessage(ctx, respMessage)
	}

	e.sessionManager.DeleteSession(sessionId)
	err := e.repo.DeleteChat(ctx, msg.ChatID, instance.Name())
	if err != nil {
		return err
	}

	sid, err := e.loadSession(ctx, instance, msg, true)
	if err != nil {
		return err
	}

	respMessage.Text = &chat.TextPayload{
		Text: fmt.Sprintf("Old session %s reserved. New session: %s", sessionId, sid),
	}
	return instance.SendMessage(ctx, respMessage)
}

func (e *Engine) cmdStatus(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	usage, err := e.repo.CountChatDataTokenBySession(ctx, sessionId)
	if err != nil {
		return err
	}

	totTokens, err := e.repo.GetTotalToken(ctx, sessionId)
	if err != nil {
		return err
	}

	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	sessData, _ := e.sessionManager.GetSession(sessionId)
	pIndex, mIndex := util.GetModelIndexWithFallback(ctx, e.llmProviders, sessData.Provider, sessData.Model)
	provider := e.llmProviders[pIndex]
	model := provider.Models()[mIndex]

	llmCallTimes := 0
	for _, data := range sessData.Events {
		if data.EventType&llm.EventMessage != 0 {
			if data.Message != nil && data.Message.Role == llm.RoleAssistant {
				llmCallTimes++
			}
		}
	}

	var respString strings.Builder
	respString.WriteString("Current status of this session:\n")
	respString.WriteString(fmt.Sprintf("Current usage data: \n"))
	respString.WriteString(fmt.Sprintf("Current session id: `%s`\n", sessionId))
	respString.WriteString(fmt.Sprintf("Current chat times: %d(call LLM API %d times)\n", len(sessData.Events), llmCallTimes))
	respString.WriteString(fmt.Sprintf("Current provider: `%s`, model: `%s`\n", provider.DisplayName(), model.DisplayName))
	respString.WriteString(fmt.Sprintf("Total Tokens Cost: %d tokens\n(input: %d tokens, output: %d tokens, thinking %d tokens)\n", usage.TotalTokens, usage.PromptTokens, usage.CompletionTokens, usage.TotalTokens-usage.PromptTokens-usage.CompletionTokens))

	totPrice := 0.0
	if model.InputPrice != 0 {
		inputPrice := model.InputPrice * float64(usage.PromptTokens) / 1000000
		totPrice += inputPrice
		respString.WriteString(fmt.Sprintf("Input Price: $%.6f/mtokens, Current Estimated Input Price Cost: $%.6f\n", model.InputPrice, inputPrice))
	}
	if model.OutputPrice != 0 {
		outputPrice := model.OutputPrice * float64(usage.TotalTokens-usage.PromptTokens) / 1000000
		totPrice += outputPrice
		respString.WriteString(fmt.Sprintf("Output Price: $%.6f/mtokens, Current Estimated Output Price Cost: $%.6f\n", model.OutputPrice, outputPrice))
	}
	if totPrice > 0 {
		respString.WriteString(fmt.Sprintf("Total Estimated Cost: $%.6f\n", totPrice))
	}

	respString.WriteString(fmt.Sprintf("Current Context length: %d / %d, now occupied %.2f%%", totTokens, model.Context, float64(totTokens*100)/float64(model.Context)))

	respMessage.Text = &chat.TextPayload{
		Text: respString.String(),
	}
	return instance.SendMessage(ctx, respMessage)
}
