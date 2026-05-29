package engine

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/chat"
)

// MessageRunner is a single goroutine running for checking messages and handle it
func (e *Engine) MessageRunner(ctx context.Context, instance chat.Provider, sessionId string, ch chan *chat.Message) error {
	displayName := instance.DisplayName()

	timer := time.NewTimer(time.Minute * 30)
	defer timer.Stop()

	for {
		select {
		case msg := <-ch:
			e.sessionService.Manager.ChatSession.SetMsg(sessionId, msg)

			color.Blue("[Manboster Engine] Runner received message from engine")
			timer.Reset(time.Minute * 30)

			isCompacted, newSessionId, err := e.handler.CheckCompact(ctx, instance, msg, sessionId)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to compact session: %s", err))
			}

			if isCompacted {
				newCh, created := e.sessionService.Manager.ChatSession.LoadOrCreateChan(newSessionId)
				if created {
					e.BuildMessageRunner(instance, newSessionId)
				}
				newCh <- msg
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
	ch, _ := e.sessionService.Manager.ChatSession.LoadOrCreateChan(sessionId)
	go func() {
		err := e.MessageRunner(cancelCtx, instance, sessionId, ch)
		if err != nil {
			color.Yellow("[Manboster Engine] We encountered an error while handling runner via %s, error: %q", instance.DisplayName(), err)
		}
	}()
}

func (e *Engine) MessageNotifyRunner(ctx context.Context, ch chan int, instance chat.Provider, msg *chat.Message) {
	for {
		select {
		case times, ok := <-ch:
			if !ok {
				return
			}
			respMsg := msg.Clone()
			respMsg.Reply = nil
			respMsg.MessageType = chat.MessageText
			respMsg.Text = &chat.TextPayload{
				Text: i18n.T(keys.GatewayLLMTryTimes, times),
			}

			err := e.gateway.SendMessage(ctx, instance, respMsg)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] we encountered an error while sending message: %q", err))
				return
			}
		case <-ctx.Done():
			return
		}
	}
}
