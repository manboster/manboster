package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
)

// HandleCommand handles commands occurs
func (e *Engine) HandleCommand(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	// TODO: check user
	if msg.Command == nil {
		return ErrInvalidParams
	}
	color.Blue(fmt.Sprintf("[Manboster Engine] Handling command... Received command: %s, args: %s", msg.Command.CommandType, msg.Command.CommandArgs))

	switch msg.Command.CommandType {
	case chat.CommandVersion:
		return e.cmdVersion(ctx, instance, msg)
	case chat.CommandId:
		return e.cmdId(ctx, instance, msg)
	case chat.CommandHelp:
		return e.cmdHelp(ctx, instance, msg)
	case chat.CommandOp:
		return e.cmdOp(ctx, instance, msg)
	case chat.CommandDeOp:
		return e.cmdDeOp(ctx, instance, msg)
	case chat.CommandStatus:
		return e.cmdStatus(ctx, instance, msg, sessionId)
	case chat.CommandSave:
		return e.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, e.cmdSave)
	case chat.CommandNew:
		return e.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, e.cmdNew)
	case chat.CommandCompact:
		return e.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, e.HandleCompact)
	case chat.CommandModel:
		return e.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, e.cmdModel)
	case chat.CommandModels:
		// TODO: interactive select
	case chat.CommandProvider:
		return e.handleAdminCommandWithSessionID(ctx, instance, msg, sessionId, e.cmdProvider)
	case chat.CommandProviders:
		// TODO: interactive select
	case chat.CommandSession:
		return e.handleAdminCommand(ctx, instance, msg, e.cmdSession)
	case chat.CommandSessions:
		// TODO: interactive select
	case chat.CommandStart:
		return e.cmdStart(ctx, instance, msg)
	case chat.CommandPair:
		return e.cmdPair(ctx, instance, msg)
	case chat.CommandCancel:
		return e.cmdCancel(ctx, instance, msg, sessionId)
	default:
		if msg.ChatType == chat.ChatsPersonal {
			return e.cmdDefault(ctx, instance, msg)
		}
	}
	return nil
}

func (e *Engine) handleAdminCommand(ctx context.Context, instance chat.Provider, msg *chat.Message, call func(ctx context.Context, instance chat.Provider, msg *chat.Message) error) error {
	isAdmin := e.safeguardService.IsAdmin(e.safeguardService.UserType(ctx, instance.Name(), msg.UserID))
	if isAdmin {
		return call(ctx, instance, msg)
	}
	return e.HandleReject(ctx, instance, msg)
}

func (e *Engine) handleAdminCommandWithSessionID(ctx context.Context, instance chat.Provider, msg *chat.Message, sid string, call func(ctx context.Context, instance chat.Provider, msg *chat.Message, sid string) error) error {
	isAdmin := e.safeguardService.IsAdmin(e.safeguardService.UserType(ctx, instance.Name(), msg.UserID))
	if isAdmin {
		return call(ctx, instance, msg, sid)
	}
	return e.HandleReject(ctx, instance, msg)
}
