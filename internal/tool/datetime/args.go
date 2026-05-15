package datetime

import "github.com/manboster/manboster/spec/schema"

type RunArgs struct {
	Name NameType `json:"name" description:"Operation to run: date returns the current date; time returns the current local time." validate:"required" enum:"date,time" example:"date"`
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
