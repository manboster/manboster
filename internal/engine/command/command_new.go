package command

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/spec/chat"
)

// cmdNew deletes old session and creates a new session
func (h *Handler) cmdNew(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	if sessionId == "" {
		respMessage.Text = &chat.TextPayload{
			Text: "Session is not active, there is nothing to do!",
		}
		return instance.SendMessage(ctx, respMessage)
	}

	_, avail := h.sessionService.Manager.ChatSession.GetSession(sessionId)
	if !avail {
		respMessage.Text = &chat.TextPayload{
			Text: "Session is not active, there is nothing to do!",
		}
		return instance.SendMessage(ctx, respMessage)
	}

	h.sessionService.Manager.ChatSession.DeleteSession(sessionId)
	err := h.repo.DeleteChat(ctx, msg.ChatID, instance.Name())
	if err != nil {
		return err
	}

	// delete chat data and session data
	err = h.repo.DeleteChatData(ctx, sessionId)
	if err != nil {
		return err
	}
	err = h.repo.DeleteSession(ctx, sessionId)
	if err != nil {
		return err
	}

	sid, err := h.sessionService.NewChatSession(ctx, instance.Name(), msg)
	if err != nil {
		return err
	}

	// auto migration
	err = h.repo.ReplaceChatSessions(ctx, sessionId, sid)
	if err != nil {
		return err
	}

	respMessage.Text = &chat.TextPayload{
		Text: fmt.Sprintf("Deleted session `%s` and created session: `%s` with default provider `%s` and model `%s`.\nIf you want to change provider or model, please use `/provider` or `/model`.\nIf you want to save and create a new session, please use `/save` command.", sessionId, sid, h.config.App.DefaultLLMProvider, h.config.App.DefaultLLMModel),
	}
	return instance.SendMessage(ctx, respMessage)
}
