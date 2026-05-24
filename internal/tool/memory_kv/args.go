package memory_kv

import "github.com/manboster/manboster/spec/schema"

type GetArgs struct {
	Key string `json:"key" description:"Memory key to retrieve." example:"EXAMPLE_KEY" validate:"required"`
}

type SetArgs struct {
	Key   string `json:"key" description:"Memory key to store." example:"EXAMPLE_KEY" validate:"required"`
	Value string `json:"value" description:"Value to store." example:"value" validate:"required"`
}

type DeleteArgs struct {
	Key string `json:"key" description:"Memory key to delete." example:"EXAMPLE_KEY" validate:"required"`
}

func (s *Service) Args() *schema.Args {
	return nil
}
