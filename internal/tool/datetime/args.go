package datetime

import "github.com/manboster/manboster/spec/schema"

type RunArgs struct {
	Name string `json:"name" description:"The name you want call, it would be enum, only 2 values: date and time, first returns current date like 2026-03-21, second returns current time in this machine like 01:14:51." validate:"required" enum:"date,time"`
}

func (s *Service) Args() []*schema.Args {
	return schema.ArgsFromStruct(RunArgs{})
}
