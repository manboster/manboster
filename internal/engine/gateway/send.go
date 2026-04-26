package gateway

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

// SendMessage builtin retries and wrapped in a single function.
func (s *Service) SendMessage(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	times := 5
	tries := 1
	var err error = nil
	for tries <= times {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		err = instance.SendMessage(timeoutCtx, msg)
		cancel()

		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Tried %d times sending via %q, got error: %q", tries, instance.Name(), err))
			time.Sleep(time.Second * time.Duration(tries+1))
			tries++
			continue
		} else {
			color.Green(fmt.Sprintf("[Manboster Engine] Tried %d times sending via %q, success.", tries, instance.Name()))
			err := instance.Notify(ctx, msg, chat.ActionSuccess)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to notify provider, error: %q", err))
			}
			return nil
		}
	}
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] Failed to send message to provider %q, error: %q", instance.DisplayName(), err))
		err := instance.Notify(ctx, msg, chat.ActionError)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to notify provider, error: %q", err))
		}
	}
	return err
}
