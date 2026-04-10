package engine

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"gorm.io/gorm"
)

func (e *Engine) HandleMessage(ctx context.Context, instance chat.Provider, msg *chat.Message) {
	color.Blue("[Manboster Engine] Handling message")
	color.Blue(fmt.Sprintf("[Manboster Engine] Got an message from %s by %s(%s), Type:%d", instance.Name(), msg.Username, msg.UserID, msg.MessageType))

	count, err := e.repo.UserCounts(ctx)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while getting user counts from repository, error: %s", err))
		return
	}
	// we get commands first, then others, in order to avoid errors
	// and command handler should authenticate!
	if msg.MessageType == chat.MessageCommand {
		cmdType := msg.Command.CommandType
		// 3 commands available when in anonymous
		if cmdType != chat.CommandId && cmdType != chat.CommandVersion && cmdType != chat.CommandPair {
			if count == 0 {
				err := e.HandleStart(ctx, instance, msg)
				if err != nil {
					color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling start guardrail via %s, error: %q", instance.Name(), err))
				}
				return
			} else {
				_, err := e.repo.UserInfo(ctx, instance.Name(), msg.UserID)
				if err != nil {
					if !errors.Is(err, gorm.ErrRecordNotFound) {
						color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while fetching user data from repository, error: %q", err))
					}
					err := e.HandleReject(ctx, instance, msg)
					if err != nil {
						color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling reject guardrail via %s, error: %q", instance.Name(), err))
					}
					return
				}
			}
		}
		err := e.HandleCommand(ctx, instance, msg)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling commands message via %s, error: %q", instance.Name(), err))
			// panic(err)
			return
		}
		return
	}

	if count == 0 {
		err := e.HandleStart(ctx, instance, msg)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling start guardrail via %s, error: %q", instance.Name(), err))
		}
		return
	}

	if msg.ChatType == chat.ChatsPersonal {
		_, err := e.repo.UserInfo(ctx, instance.Name(), msg.UserID)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while fetching user data from repository, error: %q", err))

			}
			err := e.HandleReject(ctx, instance, msg)
			if err != nil {
				color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling reject guardrail via %s, error: %q", instance.Name(), err))
			}
			return
		}
	}
	// TODO: before receiving messages, we should check users' identity.
	// get message types
	switch msg.MessageType {
	case chat.MessageText:
		err := e.HandleText(ctx, instance, msg)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling text message via %s, error: %q", instance.Name(), err))
			// panic(err)
			return
		}
		// TODO: Add more types...
	}
}
