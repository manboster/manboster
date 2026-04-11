package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/engine/commands"
)

// HandleCommand handles commands occurs
func (e *Engine) HandleCommand(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	// TODO: check user
	if msg.Command == nil {
		return ErrInvalidParams
	}
	color.Blue(fmt.Sprintf("[Manboster Engine] Handling command... Received command: %s, args: %s", msg.Command.CommandType, msg.Command.CommandArgs))

	switch msg.Command.CommandType {
	case chat.CommandVersion:
		return commands.Version(ctx, instance, msg)
	case chat.CommandId:
		return commands.Id(ctx, instance, msg)
	case chat.CommandHelp:
		// return commands.Help(ctx, instance, msg)
	case chat.CommandOp:
		return commands.Op(ctx, instance, msg, e.repo)
	case chat.CommandDeOp:
		return commands.DeOp(ctx, instance, msg, e.repo)
	case chat.CommandStatus:
		// return commands.Status(ctx, instance, msg)
	case chat.CommandSave:
		// return commands.Save(ctx, instance, msg)
	case chat.CommandNew:
	case chat.CommandSummary:
	case chat.CommandModels:
	case chat.CommandProviders:
	case chat.CommandStart:
	case chat.CommandPair:
		return commands.Pair(ctx, instance, msg, e.repo, &e.onboardLock, &e.pairKey, &e.retry, &e.userCount)
	case chat.CommandCancel:
		return commands.Cancel(ctx, instance, msg, e.sessionManager)
	default:
		return commands.Default(ctx, instance, msg)
	}
	return nil
}
