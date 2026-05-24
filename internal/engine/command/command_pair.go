package command

import (
	"context"
	"strconv"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/chat"
)

// cmdPair executes pair command
func (h *Handler) cmdPair(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	if !(msg.ChatType == chat.ChatsPersonal) {
		return nil // in order to avoid leaking...
	}
	var text string
	msg.MessageType = chat.MessageText
	if len(msg.Command.CommandArgs) == 0 {
		text = i18n.T(keys.CmdPairNoCode)
		msg.Text = &chat.TextPayload{Text: text}
		return instance.SendMessage(ctx, msg)
	}
	num, err := strconv.ParseInt(msg.Command.CommandArgs[0], 10, 64)
	if err != nil {
		text = i18n.T(keys.CmdPairInvalidNum)
		msg.Text = &chat.TextPayload{Text: text}
		return instance.SendMessage(ctx, msg)
	}
	if num < 100000 || num > 999999 {
		text = i18n.T(keys.CmdPairInvalidRange)
		msg.Text = &chat.TextPayload{Text: text}
		return instance.SendMessage(ctx, msg)
	}

	if h.onboard != nil {
		err = h.onboard.Pair(ctx, instance, msg, h.repo, num)
		if err != nil {
			text = err.Error()
		} else {
			text = i18n.T(keys.CmdPairSuccess)
			h.onboard = nil
		}
	} else {
		text = i18n.T(keys.CmdPairNoNeed)
	}

	msg.Text = &chat.TextPayload{Text: text}
	return instance.SendMessage(ctx, msg)
}
