package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/llm"
)

var bracketRegex = regexp.MustCompile(`\[.*?]\s*`)

func DescribeToHachimi(req llm.MessageToolCallRequestPayload, provider tool.Provider) string {
	var descStr strings.Builder

	metadata := provider.MetaData()
	purgedDescription := bracketRegex.ReplaceAllString(metadata.Description, "")
	descStr.WriteString(fmt.Sprintf("Model wants to call %s, description: %s, required min user type: %s, ", provider.Name(), purgedDescription, metadata.MinUserType))

	var args map[string]interface{}
	err := json.Unmarshal([]byte(fmt.Sprintf("%v", req.ToolArgs)), &args)
	if err != nil {
		color.Yellow("[Manboster Handler] Failed to unmarshal tool call result")
		return ""
	}

	schema := provider.Args()
	if schema != nil && len(schema.Properties) > 0 {
		descStr.WriteString("with the following arguments:\n")
		for _, prop := range schema.Properties {
			val := args[prop.Name]
			descStr.WriteString(fmt.Sprintf("  - %s", prop.Name))
			if prop.Description != "" {
				descStr.WriteString(fmt.Sprintf(" (%s)", prop.Description))
			}
			descStr.WriteString(fmt.Sprintf(": %v", val))
			if prop.IsEnum && len(prop.Enum) > 0 {
				descStr.WriteString(fmt.Sprintf(" (options: %v)", prop.Enum))
			}
			descStr.WriteString("\n")
		}
	} else {
		descStr.WriteString(fmt.Sprintf("\n%s", JSONParseFull(args)))
	}

	return descStr.String()
}

func JSONParseFull(j map[string]interface{}) string {
	var respStr strings.Builder

	l := len(j)
	c := 0
	for k, v := range j {
		c++
		if v == nil {
			respStr.WriteString(fmt.Sprintf("`%s`: null", k))
			if l != c {
				respStr.WriteString("\n")
			}
			continue
		}

		vOf := reflect.ValueOf(v)
		var valStr string

		switch vOf.Kind() {
		case reflect.Map:
			if nested, ok := v.(map[string]interface{}); ok {
				valStr = fmt.Sprintf("{\n%s\n}", indent(JSONParseFull(nested), "  "))
			} else {
				valStr = fmt.Sprintf("%v", v)
			}
		case reflect.Struct:
			valStr = fmt.Sprintf("%v", v)
		case reflect.Slice, reflect.Array:
			if vOf.Len() == 0 {
				valStr = "[]"
			} else {
				var elems []string
				for i := 0; i < vOf.Len(); i++ {
					elem := vOf.Index(i)
					if elem.Kind() == reflect.Interface && !elem.IsNil() {
						elem = elem.Elem()
					}
					if elem.IsValid() {
						switch elem.Kind() {
						case reflect.Map:
							if nested, ok := elem.Interface().(map[string]interface{}); ok {
								elems = append(elems, fmt.Sprintf("{\n%s\n}", indent(JSONParseFull(nested), "  ")))
							} else {
								elems = append(elems, fmt.Sprintf("%v", elem.Interface()))
							}
						default:
							elems = append(elems, fmt.Sprintf("%v", elem.Interface()))
						}
					} else {
						elems = append(elems, "null")
					}
				}
				if len(elems) == 1 {
					valStr = elems[0]
				} else {
					valStr = fmt.Sprintf("[\n%s\n]", indent(strings.Join(elems, ",\n"), "  "))
				}
			}
		default:
			valStr = fmt.Sprintf("%v", v)
		}

		respStr.WriteString(fmt.Sprintf("`%s`: %s", k, valStr))
		if l != c {
			respStr.WriteString("\n")
		}
	}

	return strings.TrimSpace(respStr.String())
}

func indent(s string, prefix string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}
