package engine

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/repository/types"
	"gorm.io/gorm"
)

func (e *Engine) HandleMessage(ctx context.Context, instance chat.Provider, msg *chat.Message) {
	color.Blue("[Manboster Engine] Handling message")
	color.Blue(fmt.Sprintf("[Manboster Engine] Got an message from %s by %s(%s), Type: %d", instance.Name(), msg.Username, msg.UserID, msg.MessageType))

	if msg.MessageType == chat.MessageCommand && (msg.Command.CommandType == chat.CommandId || msg.Command.CommandType == chat.CommandPair || msg.Command.CommandType == chat.CommandVersion) {
		err := e.HandleCommand(ctx, instance, msg)
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
		if !errors.Is(err, gorm.ErrRecordNotFound) {
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

	err = nil
	// get message types
	switch msg.MessageType {
	case chat.MessageText:
		err = e.HandleText(ctx, instance, msg)
	case chat.MessageCommand:
		err = e.HandleCommand(ctx, instance, msg)
	default:
		color.Yellow("[Manboster Engine] Ignoring message from unknown type.")
	}

	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] We encountered an error while handling message type(%d) message via %s, error: %q", msg.MessageType, instance.Name(), err))
	}
}
