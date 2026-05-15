package memory_kv

import "github.com/manboster/manboster/spec/schema"

type RunArgs struct {
	Name  NameType `json:"name" description:"Operation to run: get, set, list, or delete memory keys." validate:"required" enum:"get,set,list,delete" example:"get"`
	Key   string   `json:"key" description:"Memory key for get, set, or delete." example:"EXAMPLE_KEY"`
	Value string   `json:"value" description:"Value to store. Used only when name is set." example:"value"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
