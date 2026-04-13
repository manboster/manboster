package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
)

// RunChat is a separate goroutine running for polling chats.
func (e *Engine) RunChat(ctx context.Context, instance chat.Provider, conf any) {
	tries := 1
	for tries <= 3 {
		color.Blue(fmt.Sprintf("[Manboster Engine] Try %d times, now activating chat provider %s...", tries, instance.Name()))

		err := instance.Start(ctx, conf, func(msg *chat.Message) {
			e.HandleMessage(ctx, instance, msg)
		})

		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Failed to start a chat provider on %s, get error: %q", instance.Name(), err))
			tries++
			continue
		} else {
			// color.Green(fmt.Sprintf("[Manboster Engine] Successfully started a message provider on %s", instance.Name()))
			return
		}
	}
	if tries > 3 {
		color.Red(fmt.Sprintf("[Manboster Engine] Failed to start the chat instance: %s", instance.Name()))
		return
	}

	<-ctx.Done()
	color.Yellow(fmt.Sprintf("[Manboster Engine] Stopping chat provider: %s", instance.Name()))
	if stopErr := instance.Stop(ctx); stopErr != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] Error stopping chat provider %s: %v", instance.Name(), stopErr))
	}
}
