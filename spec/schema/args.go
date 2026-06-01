package schema

import "fmt"

// Args defines a plugin's call args tree.
type Args struct {
	Name        string   `json:"name" yaml:"name"`               // Required. The arg's name
	Type        ArgsType `json:"type" yaml:"type"`               // Required. The args's type, like int32, etc...
	Description string   `json:"description" yaml:"description"` // Required. The arg's description, you need to write descriptive so that there are something ambiguous
	Required    bool     `json:"required" yaml:"required"`       // Required. it means this is required or not.
	IsEnum      bool     `json:"isEnum" yaml:"isEnum"`           // Required. it means this is an enum type or not

	Example     any     `json:"example,omitempty" yaml:"example"`       // Optional. This is the example that displays what the value would be like.
	Properties  []*Args `json:"properties,omitempty" yaml:"properties"` // Optional. Required when ArgsType == ArgsTypeObject
	Items       *Args   `json:"items,omitempty" yaml:"items"`           // Optional. Required When ArgsType == ArgsTypeArray
	Enum        []any   `json:"enum,omitempty" yaml:"enum"`             // Optional. Required when IsEnum == true
	DisplayName string  `json:"displayName" yaml:"display_name"`        // Optional. Display name for the tool call
}

type ArgsType int16

const (
	ArgsTypeUnknown ArgsType = iota
	ArgsTypeString
	ArgsTypeInt32
	ArgsTypeUInt32
	ArgsTypeInt64
	ArgsTypeUInt64
	ArgsTypeBool
	ArgsTypeFloat
	ArgsTypeArray
	ArgsTypeObject
)

func (t ArgsType) MarshalJSON() ([]byte, error) {
	var s = "string"
	switch t {
	case ArgsTypeInt64, ArgsTypeInt32, ArgsTypeUInt32, ArgsTypeUInt64:
		s = "integer"
	case ArgsTypeFloat:
		s = "number"
	case ArgsTypeBool:
		s = "boolean"
	case ArgsTypeObject:
		s = "object"
	case ArgsTypeArray:
		s = "array"
	default:
	}
	return []byte(fmt.Sprintf("%q", s)), nil
}
