package config

import (
	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/spec/schema"
)

// Args returns args needed in configuration
type Args struct {
	Nodes []ArgsNode
	Index map[string]*ArgsNode
}

type ArgsNode struct {
	IsSecret bool
	Default  any
	Arg      *schema.Args
	Children []ArgsNode
}

func (args *Args) ToHuhGroup() *huh.Group {
	return &huh.Group{}
}

// ArgsFromStruct TODO:
func ArgsFromStruct(s interface{}) (*Args, error) {
	return &Args{}, nil
}
