package runner

import (
	"context"

	"github.com/fatih/color"
)

func (r *Runner) Run(ctx context.Context) error {
	color.Blue("[Manboster Engine] starting polling runner...")
	for {
		select {
		case <-ctx.Done():
			color.Green("[Manboster Engine] stopping polling runner...")
			return ctx.Err()
		case data := <-r.InputCh:
			if data.ChatMsg == nil {
				color.Yellow("[Manboster Engine] could not read message")
				continue
			}
			provider, ok := r.chatProviders[data.ChatMsg.Provider]
			if !ok {
				color.Yellow("[Manboster Engine] could not get provider")
				continue
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
