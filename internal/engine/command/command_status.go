package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

func (h *Handler) cmdStatus(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	usage, err := h.repo.CountChatDataTokenBySession(ctx, sessionId)
	if err != nil {
		return err
	}

	totTokens, err := h.repo.GetTotalToken(ctx, sessionId)
	if err != nil {
		return err
	}

	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	sessData, _ := h.sessionService.Manager.ChatSession.GetSession(sessionId)
	p, m := util.GetModelWithFallback(ctx, h.llmProviders, sessData.Provider, sessData.Model)
	provider := p
	model := m

	llmCallTimes := 0
	for _, data := range sessData.Events {
		if data.EventType&llm.EventMessage != 0 {
			if data.Message != nil && data.Message.Role == llm.RoleAssistant {
				llmCallTimes++
			}
		}
	}

	isFull := false
	if len(msg.Command.CommandArgs) > 0 && msg.Command.CommandArgs[0] == "full" {
		isFull = true
	}

	var respString strings.Builder
	respString.WriteString(fmt.Sprintf("Current Status of Session `%s`:\n", sessionId))
	respString.WriteString(fmt.Sprintf("Chat %d times(call LLM %d times), provider: `%s`, model: `%s`\n", len(sessData.Events), llmCallTimes, provider.DisplayName(), model.DisplayName))
	if isFull {
		respString.WriteString(fmt.Sprintf("Using %d souls: `%s`\n", len(sessData.Souls), sessData.Souls))
	}

	respString.WriteString(fmt.Sprintf("Total: %d tokens(input: %d, output: %d, thinking %d)", usage.TotalTokens, usage.PromptTokens, usage.CompletionTokens, usage.TotalTokens-usage.PromptTokens-usage.CompletionTokens))
	if usage.TotalCost > 0 {
		respString.WriteString(fmt.Sprintf("，estimated cost: `$%.6f`", usage.TotalCost))
	}
	respString.WriteString(fmt.Sprintf("\n"))

	if isFull {
		if usage.InputCost > 0 {
			respString.WriteString(fmt.Sprintf("Input price: `$%.6f`/mtokens, estimated cost: `$%.6f`\n", model.InputPrice, usage.InputCost))
		}
		if usage.OutputCost > 0 {
			respString.WriteString(fmt.Sprintf("Output price: `$%.6f`/mtokens, estimated cost: `$%.6f`\n", model.OutputPrice, usage.OutputCost))
		}
	}

	if isFull {
		modelMaps := map[string]int{}
		for _, data := range sessData.Events {
			if data.EventType&llm.EventMessage != 0 {
				if data.Model != "" {
					if data.Provider == "" {
						data.Provider = "unknown"
					}
					_, avail := modelMaps[data.Provider+":"+data.Model]
					if !avail {
						modelMaps[data.Provider+":"+data.Model] = 1
					} else {
						modelMaps[data.Provider+":"+data.Model] += 1
					}
				}
			}
		}
		if len(modelMaps) > 0 {
			respString.WriteString(fmt.Sprintf("This session's Model Usage Data:\n"))
			for modelStr, count := range modelMaps {
				respString.WriteString(fmt.Sprintf("`%s`:%d times\n", modelStr, count))
			}
			respString.WriteString(fmt.Sprintf("\n"))
		}
	}

	respString.WriteString(fmt.Sprintf("Current Context length: `%d / %d`(`%.2f%%`)", totTokens, model.Context, float64(totTokens*100)/float64(model.Context)))
	if !isFull {
		respString.WriteString("\nFor full status, please run `/status full`.")
	}
	respMessage.Text = &chat.TextPayload{
		Text: respString.String(),
	}
	return instance.SendMessage(ctx, respMessage)
}
