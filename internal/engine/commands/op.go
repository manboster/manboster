package commands

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
)

// Op Command TODO: gives Operator to replied users or given user ids.
func Op(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	return nil
}

// DeOp Command TODO：revokes an administrator.
func DeOp(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	return nil
}
