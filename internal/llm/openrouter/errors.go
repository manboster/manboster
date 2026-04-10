package openrouter

import "errors"

var ErrNoResponse = errors.New("no response")
var ErrDuplicatedModel = errors.New("duplicated model")
var ErrModelNameRequired = errors.New("model name required")
