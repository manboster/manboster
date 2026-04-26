package gateway

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
)

func withRetry(ctx context.Context, name string, times int, action func(ctx context.Context) error) error {
	var err error
	for tries := 1; tries <= times; tries++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		err = action(ctx)

		if err == nil {
			color.Green(fmt.Sprintf("[Manboster Gateway] %s success on try %d.", name, tries))
			return nil
		}

		color.Red(fmt.Sprintf("[Manboster Gateway] %s failed on try %d, error: %q", name, tries, err))
		if tries < times {
			time.Sleep(time.Second * time.Duration(tries))
		}
	}
	return fmt.Errorf("all %d attempts failed for %s, last error: %w", times, name, err)
}
