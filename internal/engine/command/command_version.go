package command

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/spec/chat"
)

// cmdVersion when user execute version commands, it will run.
func (h *Handler) cmdVersion(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{
		Text: fmt.Sprintf("Manboster: Your Personal Manbo Lobster!\nManboster version `%s %s@%s`, build at `%s`\nCheckout our latest releases here:\nhttps://github.com/manboster/manboster", config.Version, config.CurrentVersion, config.BuildCommit, config.BuildTime),
	}
	return instance.SendMessage(ctx, msg)
}
