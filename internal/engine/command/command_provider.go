package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

// cmdProvider returns current available providers, if args is empty, it would display the list of providers. if args is not empty, it would change providers to given provider id by modifying database
func (h *Handler) cmdProvider(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	var respString strings.Builder

	if len(msg.Command.CommandArgs) == 0 {
		respString.WriteString("Available Providers(If you want to see current provider, please run `/status`, if you want to change provider, please run `/provider [name]`, we will automatically change the first model of the provider for you):\n")
		i := 0
		for _, provider := range h.llmProviders {
			respString.WriteString(fmt.Sprintf("ID:`%d`) `%s`, %d available Models. Run `/provider %s` to change.\n", i+1, provider.DisplayName(), len(provider.Models()), provider.Name()))
			i += 1
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
		return instance.SendMessage(ctx, respMessage)
	}

	if _, avail := h.llmProviders[id]; !avail {
		respString.WriteString("Current provider is not found!")
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
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
		respString.WriteString("An error was occurred when updating provider name for this session!")
		color.Red(fmt.Sprintf("[Manboster Command Handler] An error was occurred when updating provider name for this session: %q", err))
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	respString.WriteString(fmt.Sprintf("Successfully changed this session's provider to `%s`, model `%s`.", providerName, modelName))
	respMessage.Text = &chat.TextPayload{
		Text: respString.String(),
	}
	return instance.SendMessage(ctx, respMessage)
}
