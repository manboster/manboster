package template

import "github.com/manboster/manboster/spec/schema"

// RunArgs is an example of demonstrating what this would be worked.
type RunArgs struct {
	Name  string `json:"name" description:"Operation to run: get, set, list, or delete keys." validate:"required" enum:"get,set,list,delete" example:"get"`
	Key   string `json:"key" description:"Key for get, set, or delete." example:"EXAMPLE_KEY"`
	Value string `json:"value" description:"Value to store. Used only when name is set." example:"value"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
