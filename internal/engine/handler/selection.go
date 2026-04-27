package handler

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

// HandleSelectionCallback handles selection and callback.
func (h *Handler) HandleSelectionCallback(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	if msg.SelectionCallback != nil {
		ch := h.selectionSessionManager.GetSelectionChan(msg.SelectionCallback.SelectionSessionId)
		if ch == nil {
			respMsg := msg.Clone()
			respMsg.MessageType = chat.MessageText
			respMsg.Text = &chat.TextPayload{
				Text: "This selection's session doesn't exist or timed out, please try again later.",
			}
			return instance.SendMessage(ctx, respMsg)
		}
		ch <- msg
		return nil
	}
	return ErrInvalidMessageType
}
