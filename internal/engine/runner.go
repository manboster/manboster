package engine

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

// MessageRunner is a single goroutine running for checking messages and handle it
func (e *Engine) MessageRunner(ctx context.Context, instance chat.Provider, sessionId string) error {
	sessChan := make(chan *chat.Message, 10)
	e.sessionService.Manager.ChatSession.CreateChan(sessionId, sessChan)
	displayName := instance.DisplayName()

	timer := time.NewTimer(time.Minute * 30)
	defer timer.Stop()

	for {
		select {
		case msg := <-sessChan:
			color.Blue("[Manboster Engine] Runner received message from engine")
			timer.Reset(time.Minute * 30)
			cancelCtx, cancel := context.WithCancel(ctx)
			e.sessionService.Manager.ChatSession.Activate(sessionId, cancel)
			err := e.MessageHandler(cancelCtx, instance, msg, sessionId)
			e.sessionService.Manager.ChatSession.Deactivate(sessionId)
			if err != nil {
				err := instance.Notify(ctx, msg, chat.ActionError)
				if err != nil {
					color.Yellow(fmt.Sprintf("[Manboster Engine] We encountered an error while notifying chat provider %s, error: %q", displayName, err))
				}
				color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling message type(%d) message via %s, error: %q", msg.MessageType, displayName, err))
			}
		case <-timer.C:
			e.sessionService.Manager.ChatSession.DeleteSession(sessionId)
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

}
