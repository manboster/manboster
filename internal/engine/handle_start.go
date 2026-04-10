package engine

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/util"
)

// HandleStart helps user tackle onboarding problems
func (e *Engine) HandleStart(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	e.onboardLock.Lock()
	if e.pairKey == 0 || e.retry > 5 {
		e.retry = 0
		e.pairKey = util.RandomNumber(100000, 999999)
		color.HiCyan(fmt.Sprintf("[Manboster Engine] !!! Your Pair Code is %d! You can enter '/pair %d' in your dialog window and adopt to this Lobster! !!!", e.pairKey, e.pairKey))
	}
	e.onboardLock.Unlock()

	var text strings.Builder
	text.WriteString(fmt.Sprintf("Welcome to use Manboster!\n"))
	text.WriteString(fmt.Sprintf("Before you start chatting, you need to adopt this Lobster in order to prevent frauds and improper adoptions. Here is what you need to do:\n"))
	text.WriteString(fmt.Sprintf("1. Open the Manboster window and fine something like '[Manboster Engine] !!! Your Pair Code is xxxxxx! You can enter '/pair xxxxxx' in your dialog window and adopt to this Lobster! !!!'.\n"))
	text.WriteString(fmt.Sprintf("(If you can't find that window, or you are running Manboster in Daemon mode, please run 'manboster log' to watch live logs.)\n"))
	text.WriteString(fmt.Sprintf("2. Get that 6-digit number code showed above and send it with '/pair', just like '/pair [your 6-digit code]'.\n"))
	text.WriteString(fmt.Sprintf("3. Let Lobster to validate it, if it is ok, you can chat with your Lobster with ease!\n"))
	text.WriteString(fmt.Sprintf("Wish you a wonderful journey with your Lobster!"))
	color.HiCyan(fmt.Sprintf("[Manboster Engine] !!! Your Pair Code is %d! You can enter '/pair %d' in your dialog window and adopt to this Lobster! !!!", e.pairKey, e.pairKey))

	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{
		Text: text.String(),
	}
	return instance.SendMessage(ctx, msg)
}
