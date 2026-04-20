package engine

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
)

// HandleCompact compacts data
func (e *Engine) HandleCompact(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	return e.chatDataService.Compact(ctx, instance, sessionId)
}
