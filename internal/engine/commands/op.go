package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
	"gorm.io/gorm"
)

// Op Command gives Operator to replied users or given user ids.
func Op(ctx context.Context, instance chat.Provider, msg *chat.Message, repo repository.Repository) error {
	// first we check whether there is any uid or not.
	args := msg.Command.CommandArgs

	msg.MessageType = chat.MessageText

	grantUserID, err := getTargetUserID(msg, args)

	err = validateUser(ctx, instance, msg, repo)
	if err != nil {
		msg.Text = &chat.TextPayload{
			Text: err.Error(),
		}
		return instance.SendMessage(ctx, msg)
	}

	// check user availability
	_, err = repo.UserInfo(ctx, instance.Name(), grantUserID)
	if err == nil {
		msg.Text = &chat.TextPayload{
			Text: "The one you want to grant permission is already an Administrator.",
		}
		return instance.SendMessage(ctx, msg)
	}

	// create user and make it as an Administrator
	err = repo.CreateUser(ctx, types.User{
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
		Text: fmt.Sprintf("Successfully created a user with ID(%s).", grantUserID),
	}
	return instance.SendMessage(ctx, msg)
}

// DeOp Command revokes an administrator.
func DeOp(ctx context.Context, instance chat.Provider, msg *chat.Message, repo repository.Repository) error {
	args := msg.Command.CommandArgs
	msg.MessageType = chat.MessageText

	targetUserID, err := getTargetUserID(msg, args)
	if err != nil {
		msg.Text = &chat.TextPayload{
			Text: err.Error(),
		}
		return instance.SendMessage(ctx, msg)
	}

	err = validateUser(ctx, instance, msg, repo)
	if err != nil {
		msg.Text = &chat.TextPayload{
			Text: err.Error(),
		}
		return instance.SendMessage(ctx, msg)
	}

	_, err = repo.UserInfo(ctx, instance.Name(), targetUserID)
	if err == nil {
		err = repo.DeleteUser(ctx, instance.Name(), targetUserID)
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
	} else {
		msg.Text = &chat.TextPayload{
			Text: "We entountered an error while trying to find the user.",
		}
		return instance.SendMessage(ctx, msg)
	}
}

func getTargetUserID(msg *chat.Message, args []string) (string, error) {
	var grantUserID string
	if msg.ChatType == chat.ChatsGroup {
		if msg.Reply != nil {
			grantUserID = msg.Reply.UserID
		} else if len(args) > 0 {
			// just like /op 1145141919
			grantUserID = args[0]
		} else {
			return "", fmt.Errorf("we can not get who you want to grant")
		}
	} else {
		if msg.Reply != nil && len(args) == 0 {
			// while args = 0, we will grant user who is in group chat and reply
			return "", fmt.Errorf("this method is used only in Group Chat mode")
		} else if len(args) != 0 {
			grantUserID = args[0]
		} else {
			return "", fmt.Errorf("we can not get who you want to grant")
		}
	}
	return grantUserID, nil
}

func validateUser(ctx context.Context, instance chat.Provider, msg *chat.Message, repo repository.Repository) error {
	// get this user's info
	uInfo, err := repo.UserInfo(ctx, instance.Name(), msg.UserID)
	// there may be some errors while getting information, like user not found etc.
	if err != nil {
		return fmt.Errorf("access denied. If you are the Administrator and sure that your are not wrong, please check your Manboster's log")
	}

	// if you're not root, you can't grant any user.
	if uInfo.Type < types.UserRoot {
		return fmt.Errorf("access denied. You don't have any permission to make user as an operator. Please use another permitted account")
	}
	return nil
}
