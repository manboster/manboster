package util

import (
	"fmt"
	"reflect"
	"strings"
)

func JSONParse(j map[string]interface{}) string {
	var respStr strings.Builder

	l := len(j)
	c := 0
	for k, v := range j {
		c++
		if v == nil {
			respStr.WriteString(fmt.Sprintf("%s: null\n", k))
			continue
		}

		vOf := reflect.ValueOf(v)
		var valStr string

		switch vOf.Kind() {
		case reflect.Map:
			valStr = "[object]"
		case reflect.Struct:
			valStr = "[struct]"
		case reflect.Slice, reflect.Array:
			if vOf.Len() == 0 {
				valStr = "[]"
			} else {
				firstElem := vOf.Index(0)
				if firstElem.Kind() == reflect.Interface && !firstElem.IsNil() {
					firstElem = firstElem.Elem()
				}

				if firstElem.IsValid() {
					switch firstElem.Kind() {
					case reflect.Map:
						valStr = "[object array]"
					case reflect.Struct:
						valStr = "[struct array]"
					default:
						valStr = fmt.Sprintf("%v", v)
					}
				} else {
					valStr = fmt.Sprintf("%v", v)
				}
			}
		default:
			valStr = fmt.Sprintf("%v", v)
		}

		if k != "shell" {
			runeValStr := []rune(valStr)
			if len(runeValStr) > 100 {
				runeValStr = append(append(runeValStr[:25], []rune("......")...), runeValStr[len(runeValStr)-25:]...)
			}
			valStr = string(runeValStr)
		}

		respStr.WriteString(fmt.Sprintf("`%s`: `%s`", k, valStr))
		if l != c {
			respStr.WriteString("; ")
		}
	}

	return strings.TrimSpace(respStr.String())
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
