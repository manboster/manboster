package handler

import "errors"

var ErrInvalidMessageType = errors.New("invalid message type or session is expired")
