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

	switch msg.Command.CommandType {
	case chat.CommandVersion:
		return h.cmdVersion(ctx, instance, msg)
	case chat.CommandId:
		return h.cmdId(ctx, instance, msg)
	case chat.CommandHelp:
		return h.cmdHelp(ctx, instance, msg)
	case chat.CommandOp:
		return h.cmdOp(ctx, instance, msg)
	case chat.CommandDeOp:
		return h.cmdDeOp(ctx, instance, msg)
	case chat.CommandStatus:
		return h.cmdStatus(ctx, instance, msg, sessionId)
	case chat.CommandSave:
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdSave)
	case chat.CommandNew:
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdNew)
	case chat.CommandCompact:
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.handler.HandleCompact)
	case chat.CommandModel:
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdModel)
	case chat.CommandModels:
		// TODO: interactive select
	case chat.CommandProvider:
		return h.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, h.cmdProvider)
	case chat.CommandProviders:
		// TODO: interactive select
	case chat.CommandSession:
		return h.handleAdminCommand(ctx, instance, msg, h.cmdSession)
	case chat.CommandSessions:
		// TODO: interactive select
	case chat.CommandStart:
		return h.cmdStart(ctx, instance, msg)
	case chat.CommandPair:
		return h.cmdPair(ctx, instance, msg)
	case chat.CommandCancel:
		return h.cmdCancel(ctx, instance, msg, sessionId)
	default:
		if msg.ChatType == chat.ChatsPersonal {
			return h.cmdDefault(ctx, instance, msg)
		}
	}
	return nil
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
