package handler

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

func (h *Handler) PassthroughContextValues(ctx context.Context, instance chat.Provider, msg *chat.Message, sid string) context.Context {
	valueCtx := context.WithValue(ctx, "chat_id", msg.ChatID)
	valueCtx = context.WithValue(valueCtx, "user_id", msg.UserID)
	valueCtx = context.WithValue(valueCtx, "chat_provider", instance.Name())
	valueCtx = context.WithValue(valueCtx, "session_id", sid)
	valueCtx = context.WithValue(valueCtx, "user_type", h.safeguardService.UserType(ctx, instance.Name(), msg.UserID).String())
	return valueCtx
}
