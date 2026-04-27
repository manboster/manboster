package handler

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

// HandleSelectionCallback handles selection and callback.
func (h *Handler) HandleSelectionCallback(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	if msg.SelectionCallback != nil {
		ch := h.selectionSessionManager.GetSelectionChan(msg.SelectionCallback.SelectionSessionId)
		ch <- msg
		return nil
	}
	return ErrInvalidMessageType
}
