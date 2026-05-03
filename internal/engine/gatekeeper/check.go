package gatekeeper

import (
	"fmt"
)

// CheckSession checks unified status of this tool call
func (s *Service) CheckSession(id string) error {
	if s.ignoranceSessionManager.GetCancelMark(id) {
		return fmt.Errorf("this user rejected all calls of this tool and please try again after 15 minutes")
	}
	return nil
}
