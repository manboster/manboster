package config

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/manboster/manboster/spec/cli"
	"github.com/manboster/manboster/spec/schema"
)

// ToCliProvider binds p to a new CliForm. Call CliForm.Build to run the
// interaction and collect values.
func (args *Args) ToCliProvider(p cli.Provider) *CliForm {
	return &CliForm{
		values: make(map[string]any),
		args:   args,
		p:      p,
	}
}

func collectProviderValues(
	nodes []ArgsNode,
	p cli.Provider,
	out map[string]any,
	prefix string,
	initial map[string]any,
) error {
	for _, node := range nodes {
		if node.Arg == nil {
			continue
		}

		key := node.Arg.Name
		if prefix != "" {
			key = prefix + "." + node.Arg.Name
		}

		// Recurse into nested objects
		if node.Arg.Type == schema.ArgsTypeObject && len(node.Children) > 0 {
			if err := collectProviderValues(node.Children, p, out, key, initial); err != nil {
				return err
			}
			continue
		}

		displayName := node.DisplayName
		if displayName == "" {
			displayName = node.Arg.Name
		}
		desc := node.Arg.Description

		getInitial := func() any {
			if initial != nil {
				if v := getNestedValue(initial, key); v != nil {
					return v
				}
			}
			return node.Default
		}

		val, err := askCliProvider(node, displayName, desc, getInitial, p)
		if err != nil {
			return fmt.Errorf("field %q: %w", key, err)
		}
		setNested(out, key, val)
	}
	return nil
}

// getNestedValue looks up a dotted key (e.g. "db.host") in a nested map.
func getNestedValue(m map[string]any, key string) any {
	parts := strings.SplitN(key, ".", 2)
	v, ok := m[parts[0]]
	if !ok {
		return nil
	}
	if len(parts) == 1 {
		return v
	}
	nested, ok := v.(map[string]any)
	if !ok {
		return nil
	}
	return getNestedValue(nested, parts[1])
}

func askCliProvider(
	node ArgsNode,
	displayName, desc string,
	getInitial func() any,
	p cli.Provider,
) (any, error) {
	required := node.Arg.Required

	switch node.Arg.Type {

	case schema.ArgsTypeString:
		if node.Arg.IsEnum && len(node.Arg.Enum) > 0 {
			opts := enumToCliOptions(node.Arg.Enum, getInitial())
			chosen, err := p.Select(displayName, desc, opts, func(o cli.Option) error {
				if required && o.Value == "" {
					return fmt.Errorf("%s is required", displayName)
				}
				return nil
			})
			if err != nil {
				return nil, err
			}
			return chosen.Value, nil
		}
		raw, err := p.Input(displayName, desc, func(input string) error {
			if required && strings.TrimSpace(input) == "" {
				return fmt.Errorf("%s is required", displayName)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		return fmt.Sprintf("%v", raw), nil

	case schema.ArgsTypeInt32, schema.ArgsTypeUInt32,
		schema.ArgsTypeInt64, schema.ArgsTypeUInt64:
		raw, err := p.Input(displayName, desc, func(input string) error {
			if required && strings.TrimSpace(input) == "" {
				return fmt.Errorf("%s is required", displayName)
			}
			if input != "" {
				if _, err := strconv.ParseInt(input, 10, 64); err != nil {
					return fmt.Errorf("%s must be an integer", displayName)
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		return fmt.Sprintf("%v", raw), nil

	case schema.ArgsTypeFloat:
		raw, err := p.Input(displayName, desc, func(input string) error {
			if required && strings.TrimSpace(input) == "" {
				return fmt.Errorf("%s is required", displayName)
			}
			if input != "" {
				if _, err := strconv.ParseFloat(input, 64); err != nil {
					return fmt.Errorf("%s must be a number", displayName)
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		return fmt.Sprintf("%v", raw), nil

	case schema.ArgsTypeBool:
		defVal := false
		if b, ok := getInitial().(bool); ok {
			defVal = b
		}
		title := fmt.Sprintf("%s (current: %v)", displayName, defVal)
		confirmed, err := p.Prompt(desc, title, "true", "false")
		if err != nil {
			return nil, err
		}
		return confirmed, nil

	case schema.ArgsTypeArray:
		if node.Arg.IsEnum && len(node.Arg.Enum) > 0 {
			opts := enumToCliOptions(node.Arg.Enum, nil)
			if node.SingleOrMultiSelect {
				if iv := getInitial(); iv != nil {
					preSelected := toStringSlice(iv)
					for i := range opts {
						for _, s := range preSelected {
							if opts[i].Value == s {
								opts[i].Selected = true
							}
						}
					}
				}
				chosen, err := p.MultiSelect(displayName, desc, opts, func(options []cli.Option) error {
					if required && len(options) == 0 {
						return fmt.Errorf("%s requires at least one selection", displayName)
					}
					return nil
				})
				if err != nil {
					return nil, err
				}
				result := make([]string, len(chosen))
				for i, c := range chosen {
					result[i] = c.Value
				}
				return result, nil
			}
			markCliSelected(opts, getInitial())
			chosen, err := p.Select(displayName, desc, opts, func(o cli.Option) error {
				if required && o.Value == "" {
					return fmt.Errorf("%s is required", displayName)
				}
				return nil
			})
			if err != nil {
				return nil, err
			}
			return chosen.Value, nil
		}
		raw, err := p.Input(displayName, desc, func(input string) error {
			if required && strings.TrimSpace(input) == "" {
				return fmt.Errorf("%s is required", displayName)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		s := strings.TrimSpace(fmt.Sprintf("%v", raw))
		if s == "" {
			return []string{}, nil
		}
		parts := strings.Split(s, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		return parts, nil

	default:
		return nil, nil
	}
}

func enumToCliOptions(enums []any, initial any) []cli.Option {
	initStr := fmt.Sprintf("%v", initial)
	opts := make([]cli.Option, 0, len(enums))
	for _, e := range enums {
		v := fmt.Sprintf("%v", e)
		opts = append(opts, cli.Option{
			Key:      v,
			Value:    v,
			Selected: v == initStr,
		})
	}
	return opts
}

func markCliSelected(opts []cli.Option, initial any) {
	initStr := fmt.Sprintf("%v", initial)
	for i := range opts {
		opts[i].Selected = opts[i].Value == initStr
	}
}

func toStringSlice(v any) []string {
	switch vv := v.(type) {
	case []string:
		return vv
	case []any:
		result := make([]string, 0, len(vv))
		for _, item := range vv {
			result = append(result, fmt.Sprintf("%v", item))
		}
		return result
	}
	return nil
}
