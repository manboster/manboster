package schema

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
)

// Validate helps to check out whether this any is fit args or not.
func Validate(s any, args Args) error {
	return validateValue(s, &args, "$", args.Required)
}

func validateValue(v any, args *Args, path string, required bool) error {
	if args == nil {
		return nil
	}
	if v == nil {
		if required {
			return fmt.Errorf("%s is required", path)
		}
		return fmt.Errorf("%s must not be null", path)
	}

	if args.IsEnum && len(args.Enum) > 0 && !enumContains(args.Enum, v) {
		return fmt.Errorf("%s must be one of %v, got %v", path, args.Enum, v)
	}

	switch args.Type {
	case ArgsTypeString:
		if _, ok := v.(string); !ok {
			return fmt.Errorf("%s must be string, got %T", path, v)
		}
		complied, err := regexp.Compile(args.Validate)
		if err != nil {
			return err
		}
		if !complied.MatchString(v.(string)) {
			return fmt.Errorf("%s must match %s", path, args.Validate)
		}
	case ArgsTypeBool:
		if _, ok := v.(bool); !ok {
			return fmt.Errorf("%s must be boolean, got %T", path, v)
		}
	case ArgsTypeFloat:
		if _, ok := numberAsFloat64(v); !ok {
			return fmt.Errorf("%s must be number, got %T", path, v)
		}
	case ArgsTypeInt32:
		if n, ok := numberAsInt64(v); !ok || n < math.MinInt32 || n > math.MaxInt32 {
			return fmt.Errorf("%s must be int32, got %v", path, v)
		}
	case ArgsTypeUInt32:
		if n, ok := numberAsUint64(v); !ok || n > math.MaxUint32 {
			return fmt.Errorf("%s must be uint32, got %v", path, v)
		}
	case ArgsTypeInt64:
		if _, ok := numberAsInt64(v); !ok {
			return fmt.Errorf("%s must be int64, got %v", path, v)
		}
	case ArgsTypeUInt64:
		if _, ok := numberAsUint64(v); !ok {
			return fmt.Errorf("%s must be uint64, got %v", path, v)
		}
	case ArgsTypeArray:
		values, ok := asSlice(v)
		if !ok {
			return fmt.Errorf("%s must be array, got %T", path, v)
		}
		if args.Items == nil {
			return nil
		}
		for i, item := range values {
			if err := validateValue(item, args.Items, fmt.Sprintf("%s[%d]", path, i), args.Items.Required); err != nil {
				return err
			}
		}
	case ArgsTypeObject:
		values, ok := asStringMap(v)
		if !ok {
			return fmt.Errorf("%s must be object, got %T", path, v)
		}
		for _, prop := range args.Properties {
			if prop == nil {
				continue
			}
			value, exists := values[prop.Name]
			if !exists {
				if prop.Required {
					return fmt.Errorf("%s.%s is required", path, prop.Name)
				}
				continue
			}
			if err := validateValue(value, prop, fmt.Sprintf("%s.%s", path, prop.Name), prop.Required); err != nil {
				return err
			}
		}
	case ArgsTypeUnknown:
		return fmt.Errorf("%s has unknown schema type", path)
	default:
		return fmt.Errorf("%s has unsupported schema type %d", path, args.Type)
	}

	return nil
}

func asStringMap(v any) (map[string]any, bool) {
	if m, ok := v.(map[string]any); ok {
		return m, true
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Map || rv.Type().Key().Kind() != reflect.String {
		return nil, false
	}

	out := make(map[string]any, rv.Len())
	iter := rv.MapRange()
	for iter.Next() {
		out[iter.Key().String()] = iter.Value().Interface()
	}
	return out, true
}

func asSlice(v any) ([]any, bool) {
	if s, ok := v.([]any); ok {
		return s, true
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, false
	}

	out := make([]any, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		out[i] = rv.Index(i).Interface()
	}
	return out, true
}

func numberAsFloat64(v any) (float64, bool) {
	switch n := v.(type) {
	case json.Number:
		f, err := n.Float64()
		return f, err == nil && !math.IsNaN(f) && !math.IsInf(f, 0)
	case float64:
		return n, !math.IsNaN(n) && !math.IsInf(n, 0)
	case float32:
		f := float64(n)
		return f, !math.IsNaN(f) && !math.IsInf(f, 0)
	case int:
		return float64(n), true
	case int8:
		return float64(n), true
	case int16:
		return float64(n), true
	case int32:
		return float64(n), true
	case int64:
		return float64(n), true
	case uint:
		return float64(n), true
	case uint8:
		return float64(n), true
	case uint16:
		return float64(n), true
	case uint32:
		return float64(n), true
	case uint64:
		return float64(n), true
	default:
		return 0, false
	}
}

func numberAsInt64(v any) (int64, bool) {
	switch n := v.(type) {
	case json.Number:
		if i, err := n.Int64(); err == nil {
			return i, true
		}
		f, err := n.Float64()
		if err != nil || math.Trunc(f) != f || f < math.MinInt64 || f > math.MaxInt64 {
			return 0, false
		}
		return int64(f), true
	case float64:
		if math.Trunc(n) != n || n < math.MinInt64 || n > math.MaxInt64 {
			return 0, false
		}
		return int64(n), true
	case float32:
		f := float64(n)
		if math.Trunc(f) != f || f < math.MinInt64 || f > math.MaxInt64 {
			return 0, false
		}
		return int64(f), true
	case int:
		return int64(n), true
	case int8:
		return int64(n), true
	case int16:
		return int64(n), true
	case int32:
		return int64(n), true
	case int64:
		return n, true
	default:
		return 0, false
	}
}

func numberAsUint64(v any) (uint64, bool) {
	switch n := v.(type) {
	case json.Number:
		if u, err := strconv.ParseUint(n.String(), 10, 64); err == nil {
			return u, true
		}
		f, err := n.Float64()
		if err != nil || math.Trunc(f) != f || f < 0 || f > math.MaxUint64 {
			return 0, false
		}
		return uint64(f), true
	case float64:
		if math.Trunc(n) != n || n < 0 || n > math.MaxUint64 {
			return 0, false
		}
		return uint64(n), true
	case float32:
		f := float64(n)
		if math.Trunc(f) != f || f < 0 || f > math.MaxUint64 {
			return 0, false
		}
		return uint64(f), true
	case uint:
		return uint64(n), true
	case uint8:
		return uint64(n), true
	case uint16:
		return uint64(n), true
	case uint32:
		return uint64(n), true
	case uint64:
		return n, true
	case int:
		if n < 0 {
			return 0, false
		}
		return uint64(n), true
	case int8:
		if n < 0 {
			return 0, false
		}
		return uint64(n), true
	case int16:
		if n < 0 {
			return 0, false
		}
		return uint64(n), true
	case int32:
		if n < 0 {
			return 0, false
		}
		return uint64(n), true
	case int64:
		if n < 0 {
			return 0, false
		}
		return uint64(n), true
	default:
		return 0, false
	}
}

func enumContains(values []any, v any) bool {
	for _, enum := range values {
		if reflect.DeepEqual(enum, v) || fmt.Sprint(enum) == fmt.Sprint(v) {
			return true
		}
	}
	return false
}
