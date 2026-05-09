package runner

import (
	"context"
	"fmt"

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
			fmt.Println(data)
			// TODO: process data
			return nil
		}
	}
}
