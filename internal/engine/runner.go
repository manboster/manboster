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
	sessChan := e.sessionService.Manager.ChatSession.GetChan(sessionId)
	displayName := instance.DisplayName()

	timer := time.NewTimer(time.Minute * 30)
	defer timer.Stop()

	for {
		select {
		case msg := <-sessChan:
			color.Blue("[Manboster Engine] Runner received message from engine")
			timer.Reset(time.Minute * 30)

			isCompacted, newSessionId, err := e.handler.CheckCompact(ctx, instance, msg, sessionId)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to compact session: %s", err))
			}
			if isCompacted {
				e.BuildMessageRunner(instance, newSessionId)
				ch := e.sessionService.Manager.ChatSession.GetChan(newSessionId)
				ch <- msg
				e.sessionService.Manager.ChatSession.DeleteSession(sessionId)
				return nil
			}

			cancelCtx, cancel := context.WithCancel(ctx)
			e.sessionService.Manager.ChatSession.Activate(sessionId, cancel)
			err = e.MessageHandler(cancelCtx, instance, msg, sessionId)
			e.sessionService.Manager.ChatSession.Deactivate(sessionId)

			if err != nil {
				err := instance.Notify(ctx, msg, chat.ActionError)
				if err != nil {
					color.Yellow(fmt.Sprintf("[Manboster Engine] We encountered an error while notifying chat provider %s, error: %q", displayName, err))
				}
				color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling message type(%d) message via %s, error: %q", msg.MessageType, displayName, err))
			}
		case <-timer.C:
			color.Yellow("[Manboster Engine Runner] Timed out for receiving message, bye!")
			e.sessionService.Manager.ChatSession.DeleteSession(sessionId)
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	}

}

func (e *Engine) BuildMessageRunner(instance chat.Provider, sessionId string) {
	color.Blue("[Manboster Engine] This session is not available in memory storage, now loading from database")
	cancelCtx, cancelFunc := context.WithCancel(context.Background())
	e.sessionService.Manager.ChatSession.SetSessionCancel(sessionId, cancelFunc)
	e.sessionService.Manager.ChatSession.CreateChan(sessionId, make(chan *chat.Message, 10))
	go func() {
		err := e.MessageRunner(cancelCtx, instance, sessionId)
		if err != nil {
			color.Yellow("[Manboster Engine] We encountered an error while handling runner via %s, error: %q", instance.DisplayName(), err)
		}
	}()
}
