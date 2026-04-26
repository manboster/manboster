package config

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/charmbracelet/huh"
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
	Arg                 *schema.Args
	Children            []ArgsNode
}

// ToHuhGroup is written by Manboster, powered by DeepSeek V4 Pro.
func (args *Args) ToHuhGroup() []*huh.Group {
	groups := make([]*huh.Group, 0)
	if args == nil {
		return groups
	}
	collectGroups(args.Nodes, &groups)
	return groups
}

func collectGroups(nodes []ArgsNode, groups *[]*huh.Group) {
	fields := make([]huh.Field, 0)
	for _, node := range nodes {
		if node.Arg == nil {
			continue
		}
		if node.Arg.Type == schema.ArgsTypeObject && len(node.Children) > 0 {
			collectGroups(node.Children, groups)
			continue
		}
		if f := toField(node); f != nil {
			fields = append(fields, f)
		}
	}
	if len(fields) > 0 {
		*groups = append(*groups, huh.NewGroup(fields...))
	}
}

func toField(node ArgsNode) huh.Field {
	name := node.Arg.Name
	desc := node.Arg.Description

	switch node.Arg.Type {
	case schema.ArgsTypeString:
		var val string
		if s, ok := node.Default.(string); ok {
			val = s
		}
		inp := huh.NewInput().Title(name).Description(desc).Value(&val)
		if node.IsSecret {
			inp.EchoMode(huh.EchoModePassword)
		}
		return inp

	case schema.ArgsTypeInt32, schema.ArgsTypeUInt32,
		schema.ArgsTypeInt64, schema.ArgsTypeUInt64,
		schema.ArgsTypeFloat:
		var val string
		if node.Default != nil {
			val = fmt.Sprintf("%v", node.Default)
		}
		return huh.NewInput().Title(name).Description(desc).Value(&val)

	case schema.ArgsTypeBool:
		var val bool
		if b, ok := node.Default.(bool); ok {
			val = b
		}
		return huh.NewConfirm().Title(name).Description(desc).Value(&val)

	case schema.ArgsTypeArray:
		if node.Arg.IsEnum && len(node.Arg.Enum) > 0 {
			opts := make([]huh.Option[string], len(node.Arg.Enum))
			for i, v := range node.Arg.Enum {
				opts[i] = huh.NewOption(fmt.Sprintf("%v", v), fmt.Sprintf("%v", v))
			}
			if node.SingleOrMultiSelect {
				var vals []string
				return huh.NewMultiSelect[string]().
					Title(name).Description(desc).Options(opts...).Value(&vals)
			}
			var val string
			return huh.NewSelect[string]().
				Title(name).Description(desc).Options(opts...).Value(&val)
		}
		var val string
		if node.Default != nil {
			val = fmt.Sprintf("%v", node.Default)
		}
		return huh.NewInput().Title(name).Description(desc).Value(&val)

	default:
		return nil
	}
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

		arg := &schema.Args{
			Name:        name,
			Type:        goKindToArgsType(field.Type.Kind()),
			Description: tag["desc"],
			Required:    tag["required"] == "true",
		}

		node := ArgsNode{
			Arg:      arg,
			IsSecret: tag["secret"] == "true",
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
	for _, part := range strings.Split(tag, ",") {
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
