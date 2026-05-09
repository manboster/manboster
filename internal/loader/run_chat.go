package loader

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

// RunChat is a separate goroutine running for polling chats.
func (l *Loader) RunChat(ctx context.Context, instance chat.Provider) {
	displayName := instance.DisplayName()

	defer func(instance chat.Provider) {
		displayName := instance.DisplayName()
		color.Yellow(fmt.Sprintf("[Manboster Loader] Stopping chat provider: %s", displayName))
		if stopErr := instance.Stop(); stopErr != nil {
			color.Red(fmt.Sprintf("[Manboster Loader] Error stopping chat provider %s: %v", displayName, stopErr))
		}
	}(instance)

	for tries := 1; tries <= 3; tries++ {
		color.Blue(fmt.Sprintf("[Manboster Loader] Try %d times, now activating chat provider %s...", tries, displayName))
		err := instance.Start(ctx, func(msg *chat.Message) {
			err := l.engine.HandleMessage(ctx, instance, msg)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Loader] Failed to handle message on %s, get error: %q", displayName, err))
				return
			}
		})

		if err != nil {
			if errors.Is(err, context.Canceled) {
				return
			}

			color.Red(fmt.Sprintf("[Manboster Loader] Failed to start a chat provider on %s, get error: %q", displayName, err))
			continue
		}

		return
	}
	color.Red(fmt.Sprintf("[Manboster Loader] Failed to start the chat instance: %s", displayName))
}
