package gateway

import "errors"

var ErrInvalidMessageType = errors.New("invalid message type")
var ErrTimeout = errors.New("timeout")
