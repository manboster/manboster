package schema

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestValidateJSONMap(t *testing.T) {
	args := Args{
		Type: ArgsTypeObject,
		Properties: []*Args{
			{Name: "name", Type: ArgsTypeString, Required: true, IsEnum: true, Enum: []any{"read", "write"}},
			{Name: "timeout", Type: ArgsTypeInt64, Required: true},
			{Name: "public", Type: ArgsTypeBool},
			{
				Name: "items",
				Type: ArgsTypeArray,
				Items: &Args{
					Type: ArgsTypeObject,
					Properties: []*Args{
						{Name: "path", Type: ArgsTypeString, Required: true},
						{Name: "size", Type: ArgsTypeUInt32},
					},
				},
			},
		},
	}

	var payload map[string]any
	if err := json.Unmarshal([]byte(`{
		"name": "read",
		"timeout": 120,
		"public": true,
		"items": [{"path": "a.txt", "size": 42}]
	}`), &payload); err != nil {
		t.Fatal(err)
	}

	if err := Validate(payload, args); err != nil {
		t.Fatalf("Validate() returned unexpected error: %v", err)
	}
}

func TestValidateRequiredField(t *testing.T) {
	args := Args{
		Type: ArgsTypeObject,
		Properties: []*Args{
			{Name: "shell", Type: ArgsTypeString, Required: true},
		},
	}

	err := Validate(map[string]any{}, args)
	if err == nil || !strings.Contains(err.Error(), "$.shell is required") {
		t.Fatalf("Validate() error = %v, want required shell error", err)
	}
}

func TestValidateRejectsWrongType(t *testing.T) {
	args := Args{
		Type: ArgsTypeObject,
		Properties: []*Args{
			{Name: "timeout", Type: ArgsTypeInt64, Required: true},
		},
	}

	err := Validate(map[string]any{"timeout": "120"}, args)
	if err == nil || !strings.Contains(err.Error(), "$.timeout must be int64") {
		t.Fatalf("Validate() error = %v, want int64 error", err)
	}
}

func TestValidateRejectsNonIntegerJSONNumber(t *testing.T) {
	args := Args{
		Type: ArgsTypeObject,
		Properties: []*Args{
			{Name: "timeout", Type: ArgsTypeInt64, Required: true},
		},
	}

	err := Validate(map[string]any{"timeout": 1.5}, args)
	if err == nil || !strings.Contains(err.Error(), "$.timeout must be int64") {
		t.Fatalf("Validate() error = %v, want int64 error", err)
	}
}

func TestValidateRejectsEnumMismatch(t *testing.T) {
	args := Args{
		Type: ArgsTypeObject,
		Properties: []*Args{
			{Name: "name", Type: ArgsTypeString, Required: true, IsEnum: true, Enum: []any{"search", "webpage"}},
		},
	}

	err := Validate(map[string]any{"name": "delete"}, args)
	if err == nil || !strings.Contains(err.Error(), "$.name must be one of") {
		t.Fatalf("Validate() error = %v, want enum error", err)
	}
}

func TestValidateArrayItemPath(t *testing.T) {
	args := Args{
		Type: ArgsTypeObject,
		Properties: []*Args{
			{
				Name: "files",
				Type: ArgsTypeArray,
				Items: &Args{
					Type: ArgsTypeObject,
					Properties: []*Args{
						{Name: "path", Type: ArgsTypeString, Required: true},
					},
				},
			},
		},
	}

	err := Validate(map[string]any{"files": []any{map[string]any{"path": "ok"}, map[string]any{"path": 3}}}, args)
	if err == nil || !strings.Contains(err.Error(), "$.files[1].path must be string") {
		t.Fatalf("Validate() error = %v, want array item path error", err)
	}
}

func TestValidateRejectsNegativeUnsigned(t *testing.T) {
	args := Args{
		Type: ArgsTypeObject,
		Properties: []*Args{
			{Name: "size", Type: ArgsTypeUInt64, Required: true},
		},
	}

	err := Validate(map[string]any{"size": -1.0}, args)
	if err == nil || !strings.Contains(err.Error(), "$.size must be uint64") {
		t.Fatalf("Validate() error = %v, want uint64 error", err)
	}
}

func TestValidateRejectsNullValue(t *testing.T) {
	args := Args{
		Type: ArgsTypeObject,
		Properties: []*Args{
			{Name: "name", Type: ArgsTypeString},
		},
	}

	err := Validate(map[string]any{"name": nil}, args)
	if err == nil || !strings.Contains(err.Error(), "$.name must not be null") {
		t.Fatalf("Validate() error = %v, want null error", err)
	}
}
