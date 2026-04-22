package config

import (
	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/spec/schema"
)

// Args returns args needed in configuration
type Args struct {
	Args     []schema.Args
	Default  any
	IsSecret bool
}

func (args *Args) ToHuhGroup() *huh.Group {
	return &huh.Group{}
}

func ArgsFromStruct(s interface{}) (*Args, error) {
	return &Args{}, nil
}
