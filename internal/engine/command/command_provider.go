package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/chat"
)

// cmdProvider returns current available providers, if args is empty, it would display the list of providers. if args is not empty, it would change providers to given provider id by modifying database
func (h *Handler) cmdProvider(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	var respString strings.Builder

	if len(msg.Command.CommandArgs) == 0 {
		respString.WriteString(i18n.T(keys.CmdProviderList))
		i := 0
		for _, provider := range h.llmProviders {
			respString.WriteString(i18n.T(keys.CmdProviderInfo, map[string]any{
				"ID":          i + 1,
				"Name":        provider.Name(),
				"DisplayName": provider.DisplayName(),
				"Count":       len(provider.Models()),
			}))
			i += 1
		}
		respMessage.Text = &chat.TextPayload{Text: respString.String()}
		return instance.SendMessage(ctx, respMessage)
	}

	id := msg.Command.CommandArgs[0]
	if id == "" {
		respString.WriteString(i18n.T(keys.CmdProviderInvalid))
		respMessage.Text = &chat.TextPayload{Text: respString.String()}
		return instance.SendMessage(ctx, respMessage)
	}

	if _, avail := h.llmProviders[id]; !avail {
		respString.WriteString(i18n.T(keys.CmdProviderNotFound))
		respMessage.Text = &chat.TextPayload{Text: respString.String()}
		return instance.SendMessage(ctx, respMessage)
	}

	providerName := h.llmProviders[id].Name()
	modelName := h.llmProviders[id].Models()[0].Name

	h.sessionService.Manager.ChatSession.SetModel(sessionId, providerName, modelName)
	err := h.repo.UpdateSession(ctx, sessionId, map[string]interface{}{
		"llm_provider":       providerName,
		"llm_provider_model": modelName,
	})
	if err != nil {
		respString.WriteString(i18n.T(keys.CmdProviderUpdateError))
		color.Red(fmt.Sprintf("[Manboster Command Handler] An error was occurred when updating provider name for this session: %q", err))
		respMessage.Text = &chat.TextPayload{Text: respString.String()}
		return instance.SendMessage(ctx, respMessage)
	}

	respString.WriteString(i18n.T(keys.CmdProviderSuccess, map[string]any{
		"Name":  providerName,
		"Model": modelName,
	}))
	respMessage.Text = &chat.TextPayload{Text: respString.String()}
	return instance.SendMessage(ctx, respMessage)
}
