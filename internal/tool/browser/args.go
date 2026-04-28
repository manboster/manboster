package browser

import "github.com/manboster/manboster/spec/schema"

type Arg struct {
}

func (s *Service) Args() *schema.Args {
	return schema.ArgsFromStruct(&Arg{})
}
