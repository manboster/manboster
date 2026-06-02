package command

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/schema"
)

// cmdOp Command gives Operator to replied users or given user ids.
func (h *Handler) cmdOp(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	args := msg.Command.CommandArgs
	msg.MessageType = chat.MessageText

	grantUserID, err := getTargetUserID(msg, args)
	if err != nil {
		msg.Text = &chat.TextPayload{Text: i18n.T(keys.CmdOpNotFound)}
		return instance.SendMessage(ctx, msg)
	}

	err = validateUser(ctx, instance, msg, h.repo)
	if err != nil {
		msg.Text = &chat.TextPayload{Text: err.Error()}
		return instance.SendMessage(ctx, msg)
	}

	_, err = h.repo.UserInfo(ctx, instance.Name(), grantUserID)
	if err == nil {
		msg.Text = &chat.TextPayload{Text: i18n.T(keys.CmdOpAlreadyAdmin)}
		return instance.SendMessage(ctx, msg)
	}

	err = h.repo.CreateUser(ctx, types.User{
		UserID:   grantUserID,
		Platform: instance.Name(),
		Type:     schema.UserAdmin,
		ID:       0,
	})
	if err != nil {
		msg.Text = &chat.TextPayload{Text: i18n.T(keys.CmdOpCreateError)}
		return instance.SendMessage(ctx, msg)
	}

	color.Red(fmt.Sprintf("[Manboster Engine] Successfully make %s operator. If this is not your action, please user /deop %s to revert it.", grantUserID, grantUserID))
	msg.Text = &chat.TextPayload{
		Text: i18n.Te(keys.CmdOpSuccess, instance.Name()+":"+grantUserID, nil),
	}
	return instance.SendMessage(ctx, msg)
}
