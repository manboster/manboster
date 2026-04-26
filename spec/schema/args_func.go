package schema

import (
	"reflect"
	"strings"
)

// ArgsFromStruct is a useful function help you convert struct to args
func ArgsFromStruct(s interface{}) *Args {
	r := reflect.ValueOf(s)
	// if it's a pointer, set the value of its pointer assets
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	if r.Kind() != reflect.Struct {
		return nil
	}
	t := r.Type()

	return &Args{
		Type:       ArgsTypeObject,
		Name:       t.Name(),
		Properties: parseFields(t),
	}
}

func parseFields(t reflect.Type) []*Args {
	var args []*Args
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		typeField := t.Field(i)
		if !f.IsExported() {
			continue
		}

		descTag := typeField.Tag.Get("description")

		name := typeField.Name
		nameJSONTag, avail := typeField.Tag.Lookup("json")
		if avail {
			name = strings.Split(nameJSONTag, ",")[0]
		}

		arg := &Args{
			Name:        name,
			Description: descTag,
			Required:    strings.Contains(f.Tag.Get("validate"), "required"),
		}

		if enumTag, ok := f.Tag.Lookup("enum"); ok {
			arg.IsEnum = true
			parts := strings.Split(enumTag, ",")
			for _, p := range parts {
				arg.Enum = append(arg.Enum, strings.TrimSpace(p))
			}
		}

		if exampleTag, ok := f.Tag.Lookup("example"); ok {
			arg.Example = exampleTag
		}

		handleKind(arg, f.Type)
		args = append(args, arg)
	}
	return args
}

func handleKind(arg *Args, rt reflect.Type) {
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	switch rt.Kind() {
	case reflect.Struct:
		arg.Type = ArgsTypeObject
		arg.Properties = parseFields(rt)
	case reflect.Map:
		arg.Type = ArgsTypeObject
		if rt.Key().Kind() == reflect.String {
			itemArg := &Args{
				Name: rt.Key().Name(),
			}
			handleKind(itemArg, rt.Elem()) // recursively get map data
			arg.Items = itemArg
		}
	case reflect.Slice, reflect.Array:
		arg.Type = ArgsTypeArray
		itemArg := &Args{Name: "item"}
		handleKind(itemArg, rt.Elem())
		arg.Items = itemArg
	case reflect.Int, reflect.Int32, reflect.Int16, reflect.Int8:
		arg.Type = ArgsTypeInt32
	case reflect.Uint, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		arg.Type = ArgsTypeUInt32
	case reflect.Int64:
		arg.Type = ArgsTypeInt64
	case reflect.Uint64:
		arg.Type = ArgsTypeUInt64
	case reflect.Float32, reflect.Float64:
		arg.Type = ArgsTypeFloat
	case reflect.String:
		arg.Type = ArgsTypeString
	case reflect.Bool:
		arg.Type = ArgsTypeBool
	default:
		arg.Type = ArgsTypeUnknown
	}
}
