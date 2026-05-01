package gateway

import (
	"context"
	"strconv"
	"time"

	"github.com/manboster/manboster/spec/chat"
)

// EditMessage builtin retries and wrapped in a single function.
func (s *Service) EditMessage(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	name := "chat_edit_" + instance.Name() + "_" + strconv.FormatInt(time.Now().Unix(), 10)
	err := withRetry(ctx, name, 5, func(ctx context.Context) error {
		timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		return instance.EditMessage(timeoutCtx, msg)
	})
	return err
}
