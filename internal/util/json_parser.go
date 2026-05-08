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
