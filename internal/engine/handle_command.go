package engine

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/engine/commands"
)

// HandleCommand handles commands occurs
func (e *Engine) HandleCommand(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	// TODO: check user
	switch msg.CommandType {
	case chat.CommandVersion:
		return commands.Version(ctx, instance, msg)
	case chat.CommandId:
		return commands.Id(ctx, instance, msg)
	case chat.CommandHelp:
		// return commands.Help(ctx, instance, msg)
	case chat.CommandOp:
		// return commands.Op(ctx, instance, msg)
	case chat.CommandDeOp:
		// return commands.DeOp(ctx, instance, msg)
	case chat.CommandStatus:
		// return commands.Status(ctx, instance, msg)
	case chat.CommandSave:
		// return commands.Save(ctx, instance, msg)
	case chat.CommandNew:
	case chat.CommandSummary:
	case chat.CommandModels:
	case chat.CommandStart:
	}
	return nil
}
