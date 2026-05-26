package gateway

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

// SendMessage builtin retries and wrapped in a single function.
func (s *Service) SendMessage(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	name := "chat_send_" + instance.Name() + "_" + strconv.FormatInt(time.Now().Unix(), 10)
	err := withRetry(ctx, name, 5, func(ctx context.Context) error {
		timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		return instance.SendMessage(timeoutCtx, msg)
	}, nil)

	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] Failed to send message to provider %q, error: %q", instance.DisplayName(), err))
		err := instance.Notify(ctx, msg, chat.ActionError)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to notify provider, error: %q", err))
		}
	} else {
		color.Green(fmt.Sprintf("[Manboster Gateway] Successfully sent message."))
		err := instance.Notify(ctx, msg, chat.ActionSuccess)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to notify provider, error: %q", err))
		}
	}
	return err
}
