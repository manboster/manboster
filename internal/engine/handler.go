package engine

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

func (e *Engine) HandleMessage(ctx context.Context, instance chat.Provider, msg *chat.Message) {
	displayName := instance.DisplayName()
	color.Blue("[Manboster Engine] Handling message")
	color.Blue(fmt.Sprintf("[Manboster Engine] Got an message from %s by %s(%s), Type: %d", displayName, msg.Username, msg.UserID, msg.MessageType))

	if msg.MessageType == chat.MessageSelectionCallback {
		err := e.handler.HandleSelectionCallback(ctx, instance, msg)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling selection callback message via %s, error: %q", displayName, err))
		}
		return
	}

	if msg.MessageType == chat.MessageCommand && chat.IsPublicCommand(msg.Command.CommandType) {
		err := e.commandHandler.Handle(ctx, instance, msg, "")
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling commands message via %s, error: %q", displayName, err))
		}
		return
	}

	if e.onboard != nil {
		if msg.ChatType == chat.ChatsPersonal {
			err := e.handler.HandleOnBoard(ctx, instance, msg)
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling onboard via %s, error: %q", displayName, err))
			}
		}
		return
	}

	//  before receiving messages, we should check users' identity.
	// get user information
	uType := e.safeguardService.UserType(ctx, instance.Name(), msg.UserID)

	if !e.safeguardService.IsAdmin(uType) && msg.ChatType == chat.ChatsPersonal {
		color.Yellow(fmt.Sprintf("[Manboster Engine] We detected an unknown user wants to talk with your lobster in person!"))
		err := e.handler.HandleReject(ctx, instance, msg)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling reject guardrail via %s, error: %q", displayName, err))
		}
		return
	}

	if msg.MessageType == chat.MessageCommand && !chat.IsSessionRequiredCommand(msg.Command.CommandType) {
		err := e.commandHandler.Handle(ctx, instance, msg, "")
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling command via %s, error: %q", displayName, err))
			return
		}
		return
	}

	// get message types
	// sessionId := e.sessionManager.ID(displayName, msg.ChatID)
	sessionId, err := e.sessionService.LoadChatSession(ctx, instance, msg, e.safeguardService.IsAdmin(uType))
	// if you're not an administrator, you can not create a new session
	if errors.Is(err, ErrAccessDenied) {
		color.Yellow(fmt.Sprintf("[Manboster Engine] We detected an unknown user wants to start a new chat!"))
		err := e.handler.HandleReject(ctx, instance, msg)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling reject guardrail via %s, error: %q", displayName, err))
		}
		return
	}
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while loading sessionId, error: %q", err))
		return
	}

	// cancel command, passthrough session lock
	if msg.MessageType == chat.MessageCommand && msg.Command.CommandType == chat.CommandCancel {
		color.Blue(fmt.Sprintf("[Manboster Engine] Handling cancel command via %s, sessionId: %s", displayName, sessionId))
		err := e.commandHandler.Handle(ctx, instance, msg, sessionId)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling command via %s, error: %q", displayName, err))
			return
		}
		return
	}

	// if there is no valid thing, we will handle it via creating channel.
	if !e.sessionService.Manager.ChatSession.AvailChan(sessionId) {
		color.Blue("[Manboster Engine] This session is not available in memory storage, now loading from database")
		cancelCtx, cancelFunc := context.WithCancel(context.Background())
		e.sessionService.Manager.ChatSession.SetSessionCancel(sessionId, cancelFunc)
		e.sessionService.Manager.ChatSession.CreateChan(sessionId, make(chan *chat.Message, 10))
		go func() {
			err := e.MessageRunner(cancelCtx, instance, sessionId)
			if err != nil {
				color.Yellow("[Manboster Engine] We encountered an error while handling runner via %s, error: %q", displayName, err)
			}
		}()
	}

	if msg.MessageType == chat.MessageCommand {
		err := e.commandHandler.Handle(ctx, instance, msg, sessionId)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Engine] We encountered an error while handling command: %q", err))
		}
		return
	}

	color.Blue("[Manboster Engine] Getting channel...")
	ch := e.sessionService.Manager.ChatSession.GetChan(sessionId)
	ch <- msg
}
