package schema

type Args struct {
	Name        string
	Type        ArgsType
	Description string
	Required    bool

	Properties map[string]*Args // When Type == Object
	Items      *Args            // When Type == Items
	Enum       []any            //
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
