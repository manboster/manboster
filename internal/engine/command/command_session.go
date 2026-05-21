package command

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/spec/chat"
)

// cmdSession return and modify session if args is empty, it would display the list of sessions. if args is not empty, it would change session to given session id by modifying database
func (h *Handler) cmdSession(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	var respString strings.Builder

	// return session ids
	if len(msg.Command.CommandArgs) == 0 {
		sessionData, err := h.repo.GetSessions(ctx)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Command Handler] we encountered an error when handling session data: %q", err))
			respString.WriteString("An error was occurred when handling session data!")
			respMessage.Text = &chat.TextPayload{
				Text: respString.String(),
			}
			return instance.SendMessage(ctx, respMessage)
		}
		respString.WriteString(fmt.Sprintf("Session List(for short, we only list 20 latest sessions, if you want to get current session id, please run `/status`.):\n"))
		for _, data := range sessionData {
			respString.WriteString(fmt.Sprintf("Session ID: `%s`(Create Time: `%s`, Provider: `%s`, Model: `%s`) Run `/session %s` to change.\n", data.SessionID, data.CreatedAt.Format("2006-01-02T15:04:05 -07"), data.LLMProvider, data.LLMProviderModel, data.SessionID))
		}
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	// checkout session
	sid := msg.Command.CommandArgs[0]
	_, err := h.repo.GetSession(ctx, sid)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			color.Yellow(fmt.Sprintf("[Manboster Command Handler] we could not found any session id"))
			respString.WriteString("The session id you entered does not exist!")
		} else {
			color.Red(fmt.Sprintf("[Manboster Command Handler] we encountered an error when getting session: %s", err))
			respString.WriteString("An error was occurred when getting session id you entered!")
		}
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	err = h.repo.UpdateChat(ctx, msg.ChatID, instance.Name(), sid)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Command Handler] we encountered an error when handling updating chat's session data: %q", err))
		respString.WriteString("An error was occurred when changing session id for this chat!")
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	respString.WriteString(fmt.Sprintf("Successfully changed session to `%s`!", sid))
	respMessage.Text = &chat.TextPayload{
		Text: respString.String(),
	}
	return instance.SendMessage(ctx, respMessage)
}
