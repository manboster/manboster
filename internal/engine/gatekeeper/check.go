package gatekeeper

import (
	"fmt"

	"github.com/manboster/manboster/internal/session/ignorance"
)

// CheckSession checks unified status of this tool call
func (s *Service) CheckSession(id string) error {
	mark, markType := s.ignoranceSessionManager.GetMark(id)
	if mark && markType == ignorance.MarkCancel {
		return fmt.Errorf("this user rejected all calls of this tool and please try again after 15 minutes")
	}
	return nil
}
