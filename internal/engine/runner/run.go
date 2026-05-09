package runner

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
)

func (r *Runner) Run(ctx context.Context) error {
	color.Blue("[Manboster Engine] starting polling runner...")
	for {
		select {
		case <-ctx.Done():
			color.Green("[Manboster Engine] stopping polling runner...")
			return ctx.Err()
		case data := <-r.InputCh:
			if data.ChatMsg != nil {
				return fmt.Errorf("could not read message")
			}
			provider, err := chat.GetProvider(data.ChatMsg.Provider)
			if err != nil {
				return err
			}
			switch data.Type {
			case MsgPrompt:
				return r.engine.HandleMessage(ctx, provider, data.ChatMsg)
			case MsgText:
				return r.gateway.SendMessage(ctx, provider, data.ChatMsg)
			}
		}
	}
}
