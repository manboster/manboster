package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

func (e *Engine) Distribute(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	displayName := instance.DisplayName()
	color.Blue("[Manboster Engine] Distributing message")
	color.Blue(fmt.Sprintf("[Manboster Engine] Got a message from %s by %s(%s), Type: %d", displayName, msg.Username, msg.UserID, msg.MessageType))

	switch msg.MessageType {
	case chat.MessageCommand:
		err := e.commandHandler.Handle(ctx, instance, msg, sessionId)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling commands message via %s, error: %q", displayName, err))
		}
		return err
	case chat.MessageSelectionCallback:
		err := e.handler.HandleSelectionCallback(ctx, instance, msg)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling selection callback message via %s, error: %q", displayName, err))
		}
		return err
	case chat.MessageUnknown:
		err := e.handler.HandleReject(ctx, instance, msg)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling reject guardrail via %s, error: %q", displayName, err))
		}
		return err
	case chat.MessageStart:
		err := e.handler.HandleOnBoard(ctx, instance, msg)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling onboard via %s, error: %q", displayName, err))
		}
		return err
	default:
	}

	color.Blue("[Manboster Engine] Getting channel...")
	ch, created := e.sessionService.Manager.ChatSession.LoadOrCreateChan(sessionId)
	if created {
		e.BuildMessageRunner(instance, sessionId)
	}
	ch <- msg
	return nil
}
