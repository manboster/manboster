package engine

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
)

func (e *Engine) HandleMessage(ctx context.Context, instance chat.Provider, msg *chat.Message) {
	displayName := instance.DisplayName()
	color.Blue("[Manboster Engine] Handling message")
	color.Blue(fmt.Sprintf("[Manboster Engine] Got an message from %s by %s(%s), Type: %d", displayName, msg.Username, msg.UserID, msg.MessageType))

	if msg.MessageType == chat.MessageCommand && chat.IsPublicCommand(msg.Command.CommandType) {
		err := e.HandleCommand(ctx, instance, msg, "")
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling commands message via %s, error: %q", displayName, err))
		}
		return
	}

	if e.userCount == 0 {
		if msg.ChatType == chat.ChatsPersonal {
			err := e.HandleStart(ctx, instance, msg)
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling start onboard via %s, error: %q", displayName, err))
			}
		}
		return
	}

	//  before receiving messages, we should check users' identity.
	// get user information
	uInfo, err := e.repo.UserInfo(ctx, instance.Name(), msg.UserID)
	if err != nil {
		// cause error!
		if !errors.Is(err, repository.ErrNotFound) {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while fetching user data from repository, error: %q", err))
		}
		uInfo = types.User{
			Type: types.UserUnknown,
		}
	}

	if uInfo.Type < types.UserAdmin && msg.ChatType == chat.ChatsPersonal {
		color.Yellow(fmt.Sprintf("[Manboster Engine] We detected an unknown user wants to talk with your lobster in person!"))
		err := e.HandleReject(ctx, instance, msg)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling reject guardrail via %s, error: %q", displayName, err))
		}
		return
	}

	if msg.MessageType == chat.MessageCommand && !chat.IsSessionRequiredCommand(msg.Command.CommandType) {
		err := e.HandleCommand(ctx, instance, msg, "")
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling command via %s, error: %q", displayName, err))
			return
		}
		return
	}

	// get message types
	// sessionId := e.sessionManager.ID(displayName, msg.ChatID)
	sessionId, err := e.loadSession(ctx, instance, msg, uInfo.Type >= types.UserAdmin)
	// if you're not an administrator, you can not create a new session
	if errors.Is(err, ErrAccessDenied) {
		color.Yellow(fmt.Sprintf("[Manboster Engine] We detected an unknown user wants to start a new chat!"))
		err := e.HandleReject(ctx, instance, msg)
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
		err := e.HandleCommand(ctx, instance, msg, sessionId)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling command via %s, error: %q", displayName, err))
			return
		}
		return
	}

	// TODO: replace it to channel queue
	lock := e.sessionManager.GetSessionLocks(sessionId)
	lock.Lock()
	defer lock.Unlock()

	// make a cancelable context
	cancelCtx, cancel := context.WithCancel(ctx)
	defer func(sid string) {
		cancel()
		e.sessionManager.Deactivate(sid)
	}(sessionId)
	e.sessionManager.Activate(sessionId, cancel)

	// we need to read model and provider.
	chatInfo, err := e.repo.GetSession(ctx, sessionId)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while getting chat data, error: %q", err))
		return
	}
	e.sessionManager.SetModel(chatInfo.SessionID, chatInfo.LLMProvider, chatInfo.LLMProviderModel)

	// then we begin to read latest messages database storages
	chatDataInfo, err := e.repo.GetChatData(ctx, sessionId)
	if len(chatDataInfo) == 0 {
		err := e.newChatData(ctx, sessionId)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while creating chat data, error: %q", err))
			return
		}
	} else if err == nil {
		err := e.mergeChatData(ctx, chatDataInfo, sessionId)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while getting chat data, error: %q", err))
			return
		}
	} else {
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while getting chat data, error: %q", err))
		return
	}

	err = nil
	switch msg.MessageType {
	case chat.MessageText:
		err = e.HandleText(cancelCtx, instance, msg, sessionId)
	case chat.MessageCommand:
		err = e.HandleCommand(cancelCtx, instance, msg, sessionId)
	default:
		color.Yellow("[Manboster Engine] Ignoring message from unknown type.")
	}

	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling message type(%d) message via %s, error: %q", msg.MessageType, displayName, err))
	}
}
