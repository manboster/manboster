package search

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Timeout  int    `json:"timeout" description:"Maximum execution time in seconds. Default is 120." validate:"required" example:"120"`
	Provider string `json:"provider" description:"Search provider, use provider to search. If you want a provider list, use 'list', if no provider specified, use 'auto'." example:"auto"`
	Content  string `json:"content" description:"Content to search" example:"Manboster"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
