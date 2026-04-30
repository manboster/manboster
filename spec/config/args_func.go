package config

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/spec/schema"
)

// ToHuhGroup converts the args tree into a Form with huh groups.
// (Partly written by Manboster, powered by DeepSeek V4 Pro)
func (args *Args) ToHuhGroup() *Form {
	form := &Form{
		Groups: make([]*huh.Group, 0),
		refs:   make([]valueRef, 0),
		args:   args,
	}
	if args == nil {
		return form
	}
	collectGroups(args.Nodes, &form.Groups, &form.refs, "", nil)
	return form
}

func collectGroups(nodes []ArgsNode, groups *[]*huh.Group, refs *[]valueRef, prefix string, initialValues map[string]any) {
	fields := make([]huh.Field, 0)
	for _, node := range nodes {
		if node.Arg == nil {
			continue
		}
		key := node.Arg.Name
		if prefix != "" {
			key = prefix + "." + key
		}

		if node.Arg.Type == schema.ArgsTypeObject && len(node.Children) > 0 {
			collectGroups(node.Children, groups, refs, key, initialValues)
			continue
		}

		f, ref := toField(node, key, initialValues)
		if f != nil {
			fields = append(fields, f)
		}
		if ref != nil {
			*refs = append(*refs, *ref)
		}
	}
	if len(fields) > 0 {
		for _, field := range fields {
			*groups = append(*groups, huh.NewGroup(field))
		}
	}
}

func toField(node ArgsNode, key string, initialValues map[string]any) (huh.Field, *valueRef) {
	name := node.Arg.Name
	displayName := node.DisplayName
	if displayName == "" {
		displayName = name
	}
	desc := node.Arg.Description

	// Read initial value from map, fallback to node.Default
	getInitial := func() any {
		if initialValues != nil {
			if v, ok := initialValues[key]; ok {
				return v
			}
		}
		return node.Default
	}

	switch node.Arg.Type {
	case schema.ArgsTypeString:
		var f huh.Field
		val := ""
		if s, ok := getInitial().(string); ok {
			val = s
		}
		if node.Arg.IsEnum {
			var options []huh.Option[string]
			for _, enum := range node.Arg.Enum {
				en, ok := enum.(string)
				if ok {
					options = append(options, huh.NewOption[string](en, en))
				}
			}
			if len(options) > 0 {
				f = huh.NewSelect[string]().Options(options...).Title(displayName).Description(desc).Value(&val)
				return f, &valueRef{key: key, ptr: &val}
			}
		}
		inp := huh.NewInput().Title(displayName).Description(desc).Value(&val)
		if node.IsSecret {
			inp.EchoMode(huh.EchoModePassword)
		}
		f = inp
		return f, &valueRef{key: key, ptr: &val}

	case schema.ArgsTypeInt32, schema.ArgsTypeUInt32,
		schema.ArgsTypeInt64, schema.ArgsTypeUInt64,
		schema.ArgsTypeFloat:
		val := ""
		if v := getInitial(); v != nil {
			val = fmt.Sprintf("%v", v)
		}
		return huh.NewInput().Title(displayName).Description(desc).Value(&val),
			&valueRef{key: key, ptr: &val}

	case schema.ArgsTypeBool:
		val := false
		if b, ok := getInitial().(bool); ok {
			val = b
		}
		return huh.NewConfirm().Title(displayName).Description(desc).Value(&val).Negative("false").Affirmative("true"),
			&valueRef{key: key, ptr: &val}

	case schema.ArgsTypeArray:
		if node.Arg.IsEnum && len(node.Arg.Enum) > 0 {
			opts := make([]huh.Option[string], len(node.Arg.Enum))
			for i, v := range node.Arg.Enum {
				opts[i] = huh.NewOption(fmt.Sprintf("%v", v), fmt.Sprintf("%v", v))
			}
			if node.SingleOrMultiSelect {
				vals := make([]string, 0)
				// Try to pre-fill multi-select from initial values
				if iv := getInitial(); iv != nil {
					switch vv := iv.(type) {
					case []string:
						vals = vv
					case []any:
						for _, item := range vv {
							if s, ok := item.(string); ok {
								vals = append(vals, s)
							}
						}
					}
				}
				return huh.NewMultiSelect[string]().
						Title(displayName).Description(desc).Options(opts...).Value(&vals),
					&valueRef{key: key, ptr: &vals}
			}
			val := ""
			if s, ok := getInitial().(string); ok {
				val = s
			}
			return huh.NewSelect[string]().
					Title(displayName).Description(desc).Options(opts...).Value(&val),
				&valueRef{key: key, ptr: &val}
		}
		val := ""
		if v := getInitial(); v != nil {
			val = fmt.Sprintf("%v", v)
		}
		return huh.NewInput().Title(displayName).Description(desc).Value(&val),
			&valueRef{key: key, ptr: &val}

	default:
		return nil, nil
	}
}
