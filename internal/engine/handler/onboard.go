package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/manboster/manboster/spec/chat"
)

// HandleOnBoard helps user tackle onboarding problems
func (h *Handler) HandleOnBoard(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	if h.onboard != nil {
		h.onboard.HandleOnBoard()
	}
	respMessage := msg.Clone()

	var text strings.Builder
	text.WriteString(fmt.Sprintf("Welcome to use Manboster!\n"))
	text.WriteString(fmt.Sprintf("Before you start chatting, you need to adopt this Lobster in order to prevent frauds and improper adoptions. Here is what you need to do:\n"))
	text.WriteString(fmt.Sprintf("1. Open the Manboster window and fine something like `[Manboster Engine] !!! Your Pair Code is xxxxxx! You can enter '/pair xxxxxx' in your dialog window and adopt to this Lobster! !!!`.\n"))
	text.WriteString(fmt.Sprintf("(If you can't find that window, or you are running Manboster in Daemon mode, please open your terminal run 'manboster log' to watch live logs.)\n"))
	text.WriteString(fmt.Sprintf("2. Get that 6-digit number code showed above and send it with `/pair`, just like `/pair [your 6-digit code]`.\n"))
	text.WriteString(fmt.Sprintf("3. Let Lobster to validate it, if it is ok, you can chat with your Lobster with ease!\n"))
	text.WriteString(fmt.Sprintf("Wish you a wonderful journey with your Lobster!"))

	respMessage.MessageType = chat.MessageText
	respMessage.Text = &chat.TextPayload{
		Text: text.String(),
	}
	return instance.SendMessage(ctx, respMessage)
}
