package telegram

import (
	"github.com/fatih/color"
)

// Stop stops Telegram bot
func (s *Service) Stop() error {
	color.Blue("[Manboster Telegram Provider] Stopping the telegram bot...")
	if s.tgInstance != nil {
		s.tgInstance.Stop()
	}
	return nil
}
