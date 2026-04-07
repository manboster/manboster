package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
)

// RunChat is a separate goroutine running for polling chats.
func (e *Engine) RunChat(ctx context.Context, instance chat.Provider, conf any) {
	//tries := 1
	//for tries <= 3 {
	err := instance.Start(ctx, conf, func(msg *chat.Message) {
		e.HandleMessage(ctx, instance, msg)
	})

	if err != nil {
		color.Red(err.Error())
		//tries++
		//continue
	} else {
		color.Green(fmt.Sprintf("[Manboster Engine]Successfully started a message provider on %s", instance.Name()))
		return
	}
	//}
	//color.Red(fmt.Sprintf("Failed to start the chat instance: %s", instance.Name()))
}
