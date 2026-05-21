package command

import (
	"context"
	"strings"

	"github.com/manboster/manboster/spec/chat"
)

func (h *Handler) cmdStart(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	var txt strings.Builder
	txt.WriteString("Welcome to use Manboster!\n")
	txt.WriteString("If this is your first use, please send something and trigger pair process.\n")
	txt.WriteString("If this Lobster is not yours, please contact owner to get access.\n")
	txt.WriteString("You can also use the following commands:\n")
	txt.WriteString("/help - show the whole help command\n")
	txt.WriteString("/id - get current information\n")
	txt.WriteString("/cancel - cancel the request.\n")
	txt.WriteString("/status - get current status of this conversation.\n")
	msg.Text = &chat.TextPayload{
		Text: txt.String(),
	}
	return instance.SendMessage(ctx, msg)
}
