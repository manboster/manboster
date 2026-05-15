package memory_md

import "github.com/manboster/manboster/spec/schema"

type RunArgs struct {
	Name  NameType `json:"name" description:"Operation to run: get reads markdown memory; set replaces it." validate:"required" enum:"get,set" example:"get"`
	Value string   `json:"value" description:"Markdown content to store. Used only when name is set." example:"# Notes\ncontent"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
