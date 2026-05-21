package command

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/spec/chat"
)

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
