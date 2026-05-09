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
			color.Blue("[Manboster Engine] received a cronjob and processing it...")
			switch data.Type {
			case MsgPrompt:
				err := r.engine.HandleMessage(ctx, provider, data.ChatMsg)
				if err != nil {
					color.Yellow("[Manboster Engine] could not handle message")
				}
			case MsgText:
				err := r.gateway.SendMessage(ctx, provider, data.ChatMsg)
				if err != nil {
					color.Yellow("[Manboster Engine] could not send message")
				}
			}
		}
	}
}
