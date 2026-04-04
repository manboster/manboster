package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/chat"
)

func Id(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	respText := strings.Builder{}
	respText.WriteString(fmt.Sprintf("Message ID: %s\n", msg.MessageID))
	respText.WriteString(fmt.Sprintf("Message User ID: %s\n", msg.UserID))
	respText.WriteString(fmt.Sprintf("Message Chat ID: %s\n", msg.ChatID))
	if msg.Reply != nil {
		respText.WriteString(fmt.Sprintf("Message Replying ID: %s\n", msg.Reply.MessageID))
		respText.WriteString(fmt.Sprintf("Message Replying Chat ID: %s\n", msg.Reply.ChatID))
		respText.WriteString(fmt.Sprintf("Message Replying User ID: %s\n", msg.Reply.UserID))
	}
	msg.Text = &chat.TextPayload{
		Text: respText.String(),
	}

	return instance.SendMessage(ctx, msg)
}
