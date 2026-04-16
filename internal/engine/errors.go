package engine

import "errors"

var ErrInvalidParams = errors.New("invalid parameters")
var ErrNoAvailableLLMProvider = errors.New("no available LLM provider found")
var ErrAccessDenied = errors.New("access denied")
