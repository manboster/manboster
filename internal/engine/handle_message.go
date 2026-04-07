package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
)

func (e *Engine) HandleMessage(ctx context.Context, instance chat.Provider, msg *chat.Message) {
	color.Blue("[Manboster Engine]Handling message")
	color.Blue(fmt.Sprintf("[Manboster Engine]Got an message from %s by %s(%s), Type:%d", instance.Name(), msg.Username, msg.UserID, msg.MessageType))
	// we get commands first, then others, in order to avoid errors
	// and command handler should authenticate!
	if msg.MessageType == chat.MessageCommand {
		err := e.HandleCommand(ctx, instance, msg)
		if err != nil {
			color.Red(err.Error())
			panic(err)
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
			panic(err)
			return
		}
		// TODO: Add more types...
	}
}
