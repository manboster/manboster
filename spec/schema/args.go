package schema

// Args defines a plugin's call args tree.
type Args struct {
	Name        string   `json:"name" yaml:"name"`               // Required. The arg's name
	Type        ArgsType `json:"type" yaml:"type"`               // Required. The args's type, like int32, etc...
	Description string   `json:"description" yaml:"description"` // Required. The arg's description, you need to write descriptive so that there are something ambiguous
	Required    bool     `json:"required" yaml:"required"`       // Required. it means this is required or not.

	Example    any     `json:"example" yaml:"example"`       // Optional. This is the example that displays what the value would be like.
	Properties []*Args `json:"properties" yaml:"properties"` // Optional. Required when ArgsType == ArgsTypeObject
	Items      *Args   `json:"items" yaml:"items"`           // Optional. Required When ArgsType == ArgsTypeArray
	Enum       []any   `json:"enum" yaml:"enum"`             // Optional. Required when ArgsType == ArgsTypeEnum
}

type ArgsType int16

const (
	ArgsTypeUnknown ArgsType = iota
	ArgsTypeString
	ArgsTypeInt32
	ArgsTypeInt64
	ArgsTypeBool
	ArgsTypeFloat
	ArgsTypeEnum
	ArgsTypeArray
	ArgsTypeObject
)
