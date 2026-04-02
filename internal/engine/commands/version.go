package commands

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
)

// Version when user execute version commands, it will run.
func Version(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	msg.Text = fmt.Sprintf("Manboster: Your Personal Manbo Lobster!\nManboster Version %s\nCheckout our latest releases here:\nhttps://github.com/manboster/manboster", config.Version)
	return instance.SendMessage(ctx, msg)
}
