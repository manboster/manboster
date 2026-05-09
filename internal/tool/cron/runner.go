package cron

import (
	"time"

	"github.com/manboster/manboster/internal/engine/runner"
)

// DelayRunner is a really simple runner, there is no need to use context.
func (s *Service) DelayRunner(delay time.Duration, msgData runner.MsgData) error {
	time.Sleep(delay)
	runner.Instance.InputCh <- msgData
	return nil
}
