package command

import (
	"context"
	"strings"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/chat"
)

func (h *Handler) cmdStart(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	var txt strings.Builder
	txt.WriteString(i18n.T(keys.CmdStartWelcome))
	txt.WriteString(i18n.T(keys.CmdStartFirstUse))
	txt.WriteString(i18n.T(keys.CmdStartNotYours))
	txt.WriteString(i18n.T(keys.CmdStartCommands))
	txt.WriteString("/help - show the whole help command\n")
	txt.WriteString("/id - get current information\n")
	txt.WriteString("/cancel - cancel the request.\n")
	txt.WriteString("/status - get current status of this conversation.\n")
	msg.Text = &chat.TextPayload{
		Text: txt.String(),
	}
	return instance.SendMessage(ctx, msg)
}
