package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
)

// cmdModel if args is empty, it would display the list of models. if args is not empty, is would change models to given provider id by modifying database
func (h *Handler) cmdModel(ctx context.Context, instance chat.Provider, msg *chat.Message, sid string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	var respString strings.Builder

	provider, model, _ := h.sessionService.Manager.ChatSession.GetModel(sid)

	p, _ := util.GetModelWithFallback(ctx, h.llmProviders, provider, model)
	if len(msg.Command.CommandArgs) == 0 {
		respString.Reset()
		respString.WriteString(i18n.T(keys.CmdModelList))
		for i, m := range p.Models() {
			respString.WriteString(i18n.T(keys.CmdModelInfo, map[string]any{
				"ID":          i + 1,
				"DisplayName": m.DisplayName,
				"Context":     m.Context,
				"MaxOutput":   m.MaxOutputTokens,
				"Input":       fmt.Sprintf("%.4f", m.InputPrice),
				"Output":      fmt.Sprintf("%.4f", m.OutputPrice),
				"Name":        m.Name,
			}) + "\n")
		}
		respMessage.Text = &chat.TextPayload{Text: respString.String()}
		return instance.SendMessage(ctx, respMessage)
	}

	id := msg.Command.CommandArgs[0]
	if id == "" {
		respString.WriteString(i18n.T(keys.CmdModelInvalid))
		respMessage.Text = &chat.TextPayload{Text: respString.String()}
		color.Red(fmt.Sprintf("[Manboster Command Handler] An error was occurred when parsing data"))
		return instance.SendMessage(ctx, respMessage)
	}

	flag := false
	for _, m := range p.Models() {
		if m.Name == id {
			flag = true
			break
		}
	}
	if !flag {
		respString.WriteString(i18n.Te(keys.CmdModelNotFound, id, nil))
		respMessage.Text = &chat.TextPayload{Text: respString.String()}
		return instance.SendMessage(ctx, respMessage)
	}

	model = id
	h.sessionService.Manager.ChatSession.SetModel(sid, provider, id)
	err := h.repo.UpdateSession(ctx, sid, map[string]interface{}{
		"llm_provider_model": model,
	})
	if err != nil {
		respString.WriteString(i18n.T(keys.CmdModelUpdateError))
		respMessage.Text = &chat.TextPayload{Text: respString.String()}
		color.Red(fmt.Sprintf("[Manboster Command Handler] An error was occurred when updating model name for this session: %q", err))
		return instance.SendMessage(ctx, respMessage)
	}

	respString.WriteString(i18n.Te(keys.CmdModelSuccess, model, nil))
	respMessage.Text = &chat.TextPayload{Text: respString.String()}
	return instance.SendMessage(ctx, respMessage)
}
