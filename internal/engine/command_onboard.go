package engine

import (
	"context"
	"strconv"

	"github.com/manboster/manboster/internal/chat"
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

	if e.onboard != nil {
		err = e.onboard.Pair(ctx, instance, msg, e.repo, num)
		if err != nil {
			text = err.Error()
		} else {
			text = "Successfully paired!"
			e.onboard = nil
		}
	} else {
		text = "There is no need to pair!"
	}

	msg.Text = &chat.TextPayload{
		Text: text,
	}
	return instance.SendMessage(ctx, msg)
}
