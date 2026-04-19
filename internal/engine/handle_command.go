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
	case chat.CommandNew:
		return e.cmdNew(ctx, instance, msg, sessionId)
	case chat.CommandCompact:
	case chat.CommandModel:
		return e.cmdModel(ctx, instance, msg)
	case chat.CommandModels:
	case chat.CommandProvider:
		return e.cmdProvider(ctx, instance, msg)
	case chat.CommandProviders:
	case chat.CommandSession:
		return e.cmdSession(ctx, instance, msg)
	case chat.CommandSessions:
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
