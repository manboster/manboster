package request

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Shell   string `json:"shell" description:"Shell command to execute." validate:"required" example:"ls -al"`
	Timeout int    `json:"timeout" description:"Maximum execution time in seconds. Default is 120." validate:"required" example:"120"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
