package command

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
	"gorm.io/gorm"
)

// cmdDeOp Command revokes an administrator.
func (h *Handler) cmdDeOp(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	args := msg.Command.CommandArgs
	msg.MessageType = chat.MessageText

	targetUserID, err := getTargetUserID(msg, args)
	if err != nil {
		msg.Text = &chat.TextPayload{
			Text: err.Error(),
		}
		return instance.SendMessage(ctx, msg)
	}

	err = validateUser(ctx, instance, msg, h.repo)
	if err != nil {
		msg.Text = &chat.TextPayload{
			Text: err.Error(),
		}
		return instance.SendMessage(ctx, msg)
	}

	_, err = h.repo.UserInfo(ctx, instance.Name(), targetUserID)
	if err == nil {
		err = h.repo.DeleteUser(ctx, instance.Name(), targetUserID)
		if err != nil {
			msg.Text = &chat.TextPayload{
				Text: "Failed to delete user.",
			}
			return instance.SendMessage(ctx, msg)
		}
		msg.Text = &chat.TextPayload{
			Text: fmt.Sprintf("Successfully degraded permission %s.", targetUserID),
		}
		color.Blue(fmt.Sprintf("[Manboster Engine] Successfully degraded user %s's permission.", targetUserID))
		return instance.SendMessage(ctx, msg)
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		msg.Text = &chat.TextPayload{
			Text: "The one you want to degrade permission is not found in database.",
		}
		return instance.SendMessage(ctx, msg)
	}

	msg.Text = &chat.TextPayload{
		Text: "We entountered an error while trying to find the user.",
	}
	return instance.SendMessage(ctx, msg)

}
