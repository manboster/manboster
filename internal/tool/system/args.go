package system

import "github.com/manboster/manboster/spec/schema"

type ProcessInfoArgs struct {
	PID int `json:"pid" description:"Target process ID." example:"12345" validate:"required"`
}

type ProcessKillArgs struct {
	PID int `json:"pid" description:"Target process ID to kill." example:"12345" validate:"required"`
}

func (s *Service) Args() *schema.Args {
	return nil
}
