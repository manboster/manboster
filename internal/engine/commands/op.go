package commands

import (
	"context"

	"github.com/manboster/manboster/internal/chat"
)

// Op Command gives Operator to replied users or given user ids.
func Op(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	// first we check whether there is any uid or not.

	return nil
}

// DeOp Command TODO：revokes an administrator.
func DeOp(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	return nil
}
