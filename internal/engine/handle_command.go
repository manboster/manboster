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
	}
	return nil
}
