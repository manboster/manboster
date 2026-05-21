package command

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

// Handle handles commands occurs
func (h *Handler) Handle(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	// TODO: check user
	if msg.Command == nil {
		return ErrInvalidParams
	}
	color.Blue(fmt.Sprintf("[Manboster Command Handler] Handling command... Received command: %s, args: %s", msg.Command.CommandType, msg.Command.CommandArgs))
	cType := msg.Command.CommandType

	return h.provider.Handle(ctx, cType, instance, msg, sessionId)
}

func (h *Handler) handleAdminCommand(ctx context.Context, instance chat.Provider, msg *chat.Message, call func(ctx context.Context, instance chat.Provider, msg *chat.Message) error) error {
	isAdmin := h.safeguardService.IsAdmin(h.safeguardService.UserType(ctx, instance.Name(), msg.UserID))
	if isAdmin {
		return call(ctx, instance, msg)
	}
	return h.handler.HandleReject(ctx, instance, msg)
}

func (h *Handler) handleAdminCommandWithSessionID(ctx context.Context, instance chat.Provider, msg *chat.Message, sid string, call func(ctx context.Context, instance chat.Provider, msg *chat.Message, sid string) error) error {
	isAdmin := h.safeguardService.IsAdmin(h.safeguardService.UserType(ctx, instance.Name(), msg.UserID))
	if isAdmin {
		return call(ctx, instance, msg, sid)
	}
	return h.handler.HandleReject(ctx, instance, msg)
}
