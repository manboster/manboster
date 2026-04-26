package engine

import "errors"

var ErrNoAvailableLLMProvider = errors.New("no available LLM provider found")
var ErrAccessDenied = errors.New("access denied")
