package system

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Name   NameType   `json:"name" description:"Operation to run: os_info returns system details; process manages running processes." validate:"required" enum:"os_info,process" example:"os_info"`
	Action ActionType `json:"action" description:"Process action: list all processes, get info for one PID, or kill one PID." enum:"list,info,kill" example:"list"`
	PID    int        `json:"pid" description:"Target process ID. Used when name is process and action is info or kill." example:"12345"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
