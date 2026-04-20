package onboard

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/util"
)

// HandleOnBoard handles onboard start
func (s *Service) HandleOnBoard() {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.pairKey == 0 || s.retry > 5 {
		if s.retry > 5 {
			color.Red("[Manboster Engine] Retry limit exceeded so we revoked the old key and created a new one.")
			s.retry = 0
		}
		s.pairKey = util.RandomNumber(100000, 999999)
	}
	for i := 0; i < 5; i++ {
		color.HiCyan(fmt.Sprintf("[Manboster Engine] !!! Your Pair Code is %d! You can enter '/pair %d' in your dialog window and adopt to this Lobster! !!!", s.pairKey, s.pairKey))
	}
}
