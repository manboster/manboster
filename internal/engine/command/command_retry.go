package command

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

// cmdRetry retries last error chat
func (h *Handler) cmdRetry(ctx context.Context, instance chat.Provider, sessionId string) error {
	ch, ok := h.sessionManager.ChatSession.LoadOrCreateChan(sessionId)
	if ok {
		h.engine.BuildMessageRunner(instance, sessionId)
	}

	msg := h.sessionManager.ChatSession.GetInputMsg(sessionId)
	if msg == nil {
		return nil
	}

	forkedMsg := msg.Fork()
	forkedMsg.MessageType |= chat.MessageFromRetry
	ch <- forkedMsg

	return nil
}
