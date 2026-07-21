package command

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/release"
	"github.com/manboster/manboster/spec/chat"
)

// cmdVersion when user execute version commands, it will run.
func (h *Handler) cmdVersion(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{
		Text: fmt.Sprintf("Manboster: Your Personal Manbo Lobster!\nManboster version `%s %s@%s`, build at `%s`\nCheckout our latest releases here:\nhttps://github.com/manboster/manboster", release.Version, release.CurrentChannel, release.BuildCommit, release.BuildTime),
	}
	return instance.SendMessage(ctx, msg)
}
