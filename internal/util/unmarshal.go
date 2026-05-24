package util

import (
	"encoding/json"
	"fmt"
)

func Unmarshal[T any](args string) (T, error) {
	var arg T
	if err := json.Unmarshal([]byte(args), &arg); err != nil {
		return arg, fmt.Errorf("invalid arguments")
	}
	return arg, nil
}
