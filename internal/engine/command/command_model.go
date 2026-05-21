package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
)

// cmdModel if args is empty, it would display the list of models. if args is not empty, is would change models to given provider id by modifying database
func (h *Handler) cmdModel(ctx context.Context, instance chat.Provider, msg *chat.Message, sid string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	var respString strings.Builder

	s, _ := h.sessionService.Manager.ChatSession.GetSession(sid)

	p, _ := util.GetModelWithFallback(ctx, h.llmProviders, s.Provider, s.Model)
	// fmt.Printf("%s %s %s %s", s.Model, s.Provider, p.DisplayName(), sid)
	if len(msg.Command.CommandArgs) == 0 {
		respString.Reset()
		respString.WriteString("Available Models(If you want to see current model, please run `/status`, if you want to change model, please run `/model [id]`.):\n")
		for i, m := range p.Models() {
			respString.WriteString(fmt.Sprintf("ID:`%d`) `%s`, context: `%d`, max output tokens: `%d` input: `$%.4f`/mtokens, output: `$%.4f`/mtokens. Run `/model %s` to change.\n", i+1, m.DisplayName, m.Context, m.MaxOutputTokens, m.InputPrice, m.OutputPrice, m.Name))
		}
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	id := msg.Command.CommandArgs[0]
	if id == "" {
		respString.WriteString("Invalid input data!\n")
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		color.Red(fmt.Sprintf("[Manboster Command Handler] An error was occurred when parsing data"))
		return instance.SendMessage(ctx, respMessage)
	}

	flag := false
	for _, m := range p.Models() {
		if m.Name == s.Model {
			flag = true
			break
		}
	}
	if !flag {
		respString.WriteString("Could not find model named `" + s.Model + "`")
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	s.Model = id
	h.sessionService.Manager.ChatSession.SetSession(sid, s)
	err := h.repo.UpdateSession(ctx, sid, map[string]interface{}{
		"llm_provider_model": s.Model,
	})
	if err != nil {
		respString.WriteString("An error was occurred when updating model name for this session!")
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		color.Red(fmt.Sprintf("[Manboster Command Handler] An error was occurred when updating model name for this session: %q", err))
		return instance.SendMessage(ctx, respMessage)
	}

	respString.WriteString(fmt.Sprintf("Successfully changed this session's model to `%s`.", s.Model))
	respMessage.Text = &chat.TextPayload{
		Text: respString.String(),
	}
	return instance.SendMessage(ctx, respMessage)
}
