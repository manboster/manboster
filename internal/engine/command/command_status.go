package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
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
	respString.WriteString(i18n.Te(keys.CmdStatusHeader, sessionId, nil))
	respString.WriteString(i18n.T(keys.CmdStatusSummary, map[string]any{
		"Count":    len(sessData.Events),
		"LLMCount": llmCallTimes,
		"Name":     provider.DisplayName(),
		"Model":    model.DisplayName,
	}))

	if isFull {
		respString.WriteString(i18n.T(keys.CmdStatusSouls, map[string]any{
			"Count": len(sessData.Souls),
			"List":  fmt.Sprintf("%s", sessData.Souls),
		}))
	}

	respString.WriteString(i18n.T(keys.CmdStatusTokens, map[string]any{
		"Count":    usage.TotalTokens,
		"Input":    usage.PromptTokens,
		"Output":   usage.CompletionTokens,
		"Thinking": usage.TotalTokens - usage.PromptTokens - usage.CompletionTokens,
	}))

	if usage.TotalCost > 0 {
		respString.WriteString(i18n.T(keys.CmdStatusCost, map[string]any{
			"Cost": fmt.Sprintf("%.4f", usage.TotalCost),
		}))
	}
	respString.WriteString("\n")

	if isFull {
		if usage.InputCost > 0 {
			respString.WriteString(i18n.T(keys.CmdStatusInputPrice, map[string]any{
				"Cost":  fmt.Sprintf("%.4f", usage.InputCost),
				"Input": fmt.Sprintf("%.4f", model.InputPrice),
			}))
		}

		if usage.OutputCost > 0 {
			respString.WriteString(i18n.T(keys.CmdStatusOutputPrice, map[string]any{
				"Cost":   fmt.Sprintf("%.4f", usage.OutputCost),
				"Output": fmt.Sprintf("%.4f", model.OutputPrice),
			}))
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
			respString.WriteString(i18n.T(keys.CmdStatusModelUsage))
			for modelStr, count := range modelMaps {
				respString.WriteString(fmt.Sprintf("`%s`:%d times\n", modelStr, count))
			}
			respString.WriteString("\n")
		}
	}

	respString.WriteString(i18n.T(keys.CmdStatusContext, map[string]any{
		"Current": totTokens,
		"Full":    usage.TotalTokens,
		"Percent": fmt.Sprintf("%.2f", float64(totTokens*100)/float64(model.Context)),
	}))

	if !isFull {
		respString.WriteString(i18n.T(keys.CmdStatusFullHint))
	}

	respMessage.Text = &chat.TextPayload{
		Text: respString.String(),
	}
	return instance.SendMessage(ctx, respMessage)
}
