package engine

import (
	"context"
	"strconv"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/repository/types"
)

// cmdPair executes pair command
func (e *Engine) cmdPair(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	var text string
	msg.MessageType = chat.MessageText
	if len(msg.Command.CommandArgs) == 0 {
		text = "No pair code provided!"
		msg.Text = &chat.TextPayload{
			Text: text,
		}
		return instance.SendMessage(ctx, msg)
	}
	num, err := strconv.ParseInt(msg.Command.CommandArgs[0], 10, 64)
	if err != nil {
		text = "What you've input is not a valid number!"
		msg.Text = &chat.TextPayload{
			Text: text,
		}
		return instance.SendMessage(ctx, msg)
	}
	if num < 100000 || num > 999999 {
		text = "Invalid number range!"
		msg.Text = &chat.TextPayload{
			Text: text,
		}
		return instance.SendMessage(ctx, msg)
	}
	e.onboardLock.Lock()
	if num == e.pairKey {
		text = "Successfully paired!"
		err := e.repo.CreateUser(ctx, types.User{
			ID:       0,
			UserID:   msg.UserID,
			Platform: instance.Name(),
			Type:     types.UserRoot,
		})
		if err != nil {
			text += " But we failed to create the user! Error: " + err.Error()
		} else {
			text += "\nEnjoy using your personal Lobster!"
			e.pairKey = 0
			e.retry = 0
			e.retry += 1
		}
	} else {
		text = "Pair failed, invalid pair code, please check your code!"
		e.retry++
	}
	e.onboardLock.Unlock()

	msg.Text = &chat.TextPayload{
		Text: text,
	}
	return instance.SendMessage(ctx, msg)
}
