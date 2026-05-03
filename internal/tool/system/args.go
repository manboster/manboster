package system

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Name   NameType   `json:"name" description:"The name you want call, it would be enum, only these values: os_info. 'os_info' returns the information of this machine, including OS and architecture, 'process' can help you manage process running in this machine." validate:"required" enum:"os_info,process" example:"os_info"`
	Action ActionType `json:"action" description:"The action you want to call, it would be enum, only these values: list, info or delete. 'list' lists all running progress, 'info' gets a particular info of the progress and 'kill' kills the progress." enum:"list,info,kill" example:"list,info,kill"`
	PID    int        `json:"pid" description:"The pid you want to get information or delete. Only valid in 'name' = 'process' and ('action' = 'info' or 'action' = 'delete')" example:"12345"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
