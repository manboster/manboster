package config

import (
	"reflect"
	"strings"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/spec/schema"
)

// Args returns args needed in configuration
type Args struct {
	Nodes []ArgsNode
	Index map[string]*ArgsNode
}

type ArgsNode struct {
	IsSecret            bool
	Default             any
	SingleOrMultiSelect bool // true is multi, false is single, only valid in array mode.
	DisplayName         string
	Arg                 *schema.Args
	Children            []ArgsNode
}

// ArgsFromStruct builds args from a struct
func ArgsFromStruct(s interface{}) *Args {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	t := v.Type()
	var nodes []ArgsNode

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}

		tag := parseTag(field.Tag.Get("manboconfig"))
		name := jsonName(field)
		if tag["skip"] == "true" {
			continue
		}

		var enums []any
		isEnum := false
		enum, avail := field.Tag.Lookup("enum")
		if avail {
			enumList := strings.Split(enum, ",")
			if len(enumList) != 0 {
				isEnum = true
				for _, en := range enumList {
					enums = append(enums, en)
				}
			}
		}

		desc := tag["desc"]
		if tag["id"] != "" {
			desc = i18n.T(tag["id"] + ".desc")
		}
		dpName := tag["name"]
		if tag["id"] != "" {
			dpName = i18n.T(tag["id"])
		}

		arg := &schema.Args{
			Name:        name,
			Type:        goKindToArgsType(field.Type.Kind()),
			Description: desc,
			Required:    tag["required"] == "true",
			Enum:        enums,
			IsEnum:      isEnum,
		}

		validate, avail := field.Tag.Lookup("validation")
		if avail {
			arg.Validate = validate
		}

		node := ArgsNode{
			Arg:         arg,
			IsSecret:    tag["secret"] == "true",
			DisplayName: dpName,
		}

		if d := tag["default"]; d != "" {
			node.Default = d
		}

		if field.Type.Kind() == reflect.Struct {
			child := ArgsFromStruct(v.Field(i).Interface())
			if child != nil {
				node.Children = child.Nodes
			}
		}

		nodes = append(nodes, node)
	}

	return &Args{Nodes: nodes, Index: buildIndex(nodes)}
}

func parseTag(tag string) map[string]string {
	result := make(map[string]string)
	for _, part := range strings.Split(tag, ";") {
		if kv := strings.SplitN(part, ":", 2); len(kv) == 2 {
			result[kv[0]] = kv[1]
		} else {
			result[part] = "true"
		}
	}
	return result
}

func jsonName(field reflect.StructField) string {
	name := field.Tag.Get("json")
	if idx := strings.Index(name, ","); idx != -1 {
		name = name[:idx]
	}
	if name == "" {
		name = field.Name
	}
	return name
}

func goKindToArgsType(k reflect.Kind) schema.ArgsType {
	switch k {
	case reflect.String:
		return schema.ArgsTypeString
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return schema.ArgsTypeInt64
	case reflect.Bool:
		return schema.ArgsTypeBool
	case reflect.Slice, reflect.Array:
		return schema.ArgsTypeArray
	case reflect.Map, reflect.Struct:
		return schema.ArgsTypeObject
	default:
		return schema.ArgsTypeUnknown
	}
}

func buildIndex(nodes []ArgsNode) map[string]*ArgsNode {
	index := make(map[string]*ArgsNode)
	for i := range nodes {
		node := &nodes[i]
		if node.Arg != nil && node.Arg.Name != "" {
			index[node.Arg.Name] = node
		}
		for k, v := range buildIndex(node.Children) {
			index[k] = v
		}
	}
	return index
}
