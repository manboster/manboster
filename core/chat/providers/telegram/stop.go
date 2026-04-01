package telegram

import (
	"context"

	"github.com/fatih/color"
)

// Stop stops Telegram bot
func (s *Service) Stop(ctx context.Context) error {
	<-ctx.Done()
	color.Blue("Stopping the telegram bot...")
	s.tgInstance.Stop()
	return nil
}
