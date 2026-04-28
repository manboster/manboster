package shell

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Shell   string `json:"shell" description:"The shell you want to execute in the system" validate:"required" example:"ls -al"`
	Timeout int    `json:"timeout" description:"The maximum timeout for this execution. Default is 120." validate:"required" example:"120"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
