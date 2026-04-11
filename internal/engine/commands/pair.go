package commands

import (
	"context"
	"strconv"
	"sync"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
)

// Pair executes pair command
func Pair(ctx context.Context, instance chat.Provider, msg *chat.Message, repo repository.Repository, lock *sync.Mutex, code *int64, retry *int, count *int) error {
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
	lock.Lock()
	if num == *code {
		text = "Successfully paired!"
		err := repo.CreateUser(ctx, types.User{
			ID:       0,
			UserID:   msg.UserID,
			Platform: instance.Name(),
			Type:     types.UserRoot,
		})
		if err != nil {
			text += " But we failed to create the user! Error: " + err.Error()
		} else {
			text += "\nEnjoy using your personal Lobster!"
			*code = 0
			*retry = 0
			*count += 1
		}
	} else {
		text = "Pair failed, invalid pair code, please check your code!"
		*retry++
	}
	lock.Unlock()

	msg.Text = &chat.TextPayload{
		Text: text,
	}
	return instance.SendMessage(ctx, msg)
}
