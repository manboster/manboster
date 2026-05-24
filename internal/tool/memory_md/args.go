package memory_md

import "github.com/manboster/manboster/spec/schema"

type SetArgs struct {
	Value string `json:"value" description:"Markdown content to store." example:"# Notes\ncontent" validate:"required"`
}

func (s *Service) Args() *schema.Args {
	return nil
}
