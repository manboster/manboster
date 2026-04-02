package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
)

// RunChat is a separate goroutine running for polling chats.
func (e *Engine) RunChat(ctx context.Context, instance chat.Provider, conf any) {
	err := instance.Start(ctx, conf, func(msg *chat.Message) {
		color.Blue(fmt.Sprintf("Got an message from %s by %s(%s)", instance.Name(), msg.Username, msg.UserID))

		// we get commands first, then others, in order to avoid errors
		// and command handler should authenticate!
		if msg.MessageType == chat.MessageCommand {
			err := e.HandleCommand(ctx, instance, msg)
			if err != nil {
				color.Red(err.Error())
				return
			}
			return
		}

		// TODO: before receiving messages, we should check users' identity.
		// get message types
		switch msg.MessageType {
		case chat.MessageText:
			err := e.HandleText(ctx, instance, msg)
			if err != nil {
				color.Red(err.Error())
				return
			}
			// TODO: Add more types...
		}

	})
	if err != nil {
		color.Red(err.Error())
		return
	}
}
