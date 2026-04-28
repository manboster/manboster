package memory_md

import "github.com/manboster/manboster/spec/schema"

type RunArgs struct {
	Name  string `json:"name" description:"The name you want call, it would be enum, only 4 values: get, set. 'get' returns the markdown content, 'set' sets value of the markdown to the file." validate:"required" enum:"get,set" example:"get"`
	Value string `json:"value" description:"The markdown content to use to set or get." example:"# MD Content....\ncontent123..."`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
