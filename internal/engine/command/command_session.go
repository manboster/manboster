package command

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/spec/chat"
)

// cmdSession return and modify session if args is empty, it would display the list of sessions. if args is not empty, it would change session to given session id by modifying database
func (h *Handler) cmdSession(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	var respString strings.Builder

	if len(msg.Command.CommandArgs) == 0 {
		sessionData, err := h.repo.GetSessions(ctx)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Command Handler] we encountered an error when handling session data: %q", err))
			respString.WriteString(i18n.T(keys.CmdSessionDataError))
			respMessage.Text = &chat.TextPayload{Text: respString.String()}
			return instance.SendMessage(ctx, respMessage)
		}
		respString.WriteString(i18n.T(keys.CmdSessionList))

		for i, data := range sessionData {
			respString.WriteString(i18n.T(keys.CmdSessionInfo, map[string]any{
				"ID":      i + 1,
				"Created": data.CreatedAt.Format("2006-01-02T15:04:05-07"),
				"Name":    data.LLMProvider,
				"Model":   data.LLMProviderModel,
				"Session": data.SessionID,
			}))
		}
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}

		return instance.SendMessage(ctx, respMessage)
	}

	sid := msg.Command.CommandArgs[0]
	_, err := h.repo.GetSession(ctx, sid)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			color.Yellow(fmt.Sprintf("[Manboster Command Handler] we could not found any session id"))
			respString.WriteString(i18n.T(keys.CmdSessionNotFound))
		} else {
			color.Red(fmt.Sprintf("[Manboster Command Handler] we encountered an error when getting session: %s", err))
			respString.WriteString(i18n.T(keys.CmdSessionGetError))
		}

		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}

		return instance.SendMessage(ctx, respMessage)
	}

	err = h.repo.UpdateChat(ctx, msg.ChatID, instance.Name(), sid)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Command Handler] we encountered an error when handling updating chat's session data: %q", err))
		respString.WriteString(i18n.T(keys.CmdSessionUpdateError))
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}

		return instance.SendMessage(ctx, respMessage)
	}

	respString.WriteString(i18n.Te(keys.CmdSessionSuccess, sid, nil))
	respMessage.Text = &chat.TextPayload{Text: respString.String()}
	return instance.SendMessage(ctx, respMessage)
}
