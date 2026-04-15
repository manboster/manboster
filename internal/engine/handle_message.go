package engine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/internal/session"
	"github.com/manboster/manboster/internal/util"
)

func (e *Engine) HandleMessage(ctx context.Context, instance chat.Provider, msg *chat.Message) {
	color.Blue("[Manboster Engine] Handling message")
	color.Blue(fmt.Sprintf("[Manboster Engine] Got an message from %s by %s(%s), Type: %d", instance.Name(), msg.Username, msg.UserID, msg.MessageType))

	if msg.MessageType == chat.MessageCommand && chat.IsPublicCommand(msg.Command.CommandType) {
		err := e.HandleCommand(ctx, instance, msg, "")
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling commands message via %s, error: %q", instance.Name(), err))
		}
		return
	}

	if e.userCount == 0 {
		if msg.ChatType == chat.ChatsPersonal {
			err := e.HandleStart(ctx, instance, msg)
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling start onboard via %s, error: %q", instance.Name(), err))
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
		err := e.HandleReject(ctx, instance, msg)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling reject guardrail via %s, error: %q", instance.Name(), err))
		}
		return
	}

	if msg.MessageType == chat.MessageCommand && !chat.IsSessionRequiredCommand(msg.Command.CommandType) {
		err := e.HandleCommand(ctx, instance, msg, "")
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling command via %s, error: %q", instance.Name(), err))
			return
		}
		return
	}

	// get message types
	// sessionId := e.sessionManager.ID(instance.Name(), msg.ChatID)
	var sessionId string
	chatInfo, err := e.repo.GetChat(ctx, instance.Name(), msg.ChatID)
	if err == nil {
		sessionId = chatInfo.SessionID
	} else if errors.Is(err, repository.ErrNotFound) {
		sessionId = util.RandomString(32)
		err := e.repo.CreateSession(ctx, types.Session{
			SessionID:        sessionId,
			LLMProvider:      e.config.App.DefaultLLMProvider,
			LLMProviderModel: e.config.App.DefaultLLMModel,
		})
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while creating session to repository, error: %q", err))
			return
		}
		// set a new session
		e.sessionManager.SetSession(sessionId, session.Session{
			Model:    e.config.App.DefaultLLMModel,
			Provider: e.config.App.DefaultLLMProvider,
			Messages: []llm.Message{},
			Active:   false,
			Cancel:   nil,
		})

		err = e.repo.CreateChat(ctx, types.Chat{
			ChatID:         msg.ChatID,
			SessionID:      sessionId,
			ChatProvider:   instance.Name(),
			ChatPermission: 1, // TODO: add chat permission but there is no need, so occupy?
		})
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while creating chat to repository, error: %q", err))
			return
		}
	} else {
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while reading user information to repository, error: %q", err))
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

	// then we begin to read latest messages database storages
	chatDataInfo, err := e.repo.GetChatData(ctx, sessionId)
	if err == nil {

	} else if errors.Is(err, repository.ErrNotFound) {
		textPayload := &llm.MessageTextPayload{
			Text: config.InitialSystemPrompt, // TODO: prompt engineering
		}
		e.sessionManager.AppendMessage(sessionId, llm.Message{
			Role: llm.RoleSystem,
			Text: textPayload,
			Type: llm.MessageText,
		})
		jsonPayload, err := json.Marshal(textPayload)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while jsontifying payload, error: %q", err))
			return
		}
		// create a chat data
		err = e.repo.CreateChatData(cancelCtx, types.ChatData{
			SessionID:        sessionId,
			Role:             llm.RoleSystem,
			MessageType:      llm.MessageText,
			MessagePayload:   string(jsonPayload),
			PromptTokens:     0,
			CompletionTokens: 0,
			TotalTokens:      0,
		})
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while appending chat data to repository, error: %q", err))
			return
		}
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
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling message type(%d) message via %s, error: %q", msg.MessageType, instance.Name(), err))
	}
}
