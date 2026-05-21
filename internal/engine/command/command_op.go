package command

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/spec/chat"
)

// cmdOp Command gives Operator to replied users or given user ids.
func (h *Handler) cmdOp(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	// first we check whether there is any uid or not.
	args := msg.Command.CommandArgs

	msg.MessageType = chat.MessageText

	grantUserID, err := getTargetUserID(msg, args)
	if err != nil {
		msg.Text = &chat.TextPayload{
			Text: "The one you want to grant permission is not found in database.",
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

	// check user availability
	_, err = h.repo.UserInfo(ctx, instance.Name(), grantUserID)
	if err == nil {
		msg.Text = &chat.TextPayload{
			Text: "The one you want to grant permission is already an Administrator.",
		}
		return instance.SendMessage(ctx, msg)
	}

	// create user and make it as an Administrator
	err = h.repo.CreateUser(ctx, types.User{
		UserID:   grantUserID,
		Platform: instance.Name(),
		Type:     types.UserAdmin,
		ID:       0,
	})
	if err != nil {
		msg.Text = &chat.TextPayload{
			Text: "We encountered an error while trying to create the user.",
		}
		return instance.SendMessage(ctx, msg)
	}

	color.Red(fmt.Sprintf("[Manboster Engine] Successfully make %s operator. If this is not your action, please user /deop %s to revert it.", grantUserID, grantUserID))
	msg.Text = &chat.TextPayload{
		Text: fmt.Sprintf("Successfully created make (%s) as an Administrator.", grantUserID),
	}
	return instance.SendMessage(ctx, msg)
}
