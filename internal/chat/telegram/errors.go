package telegram

import "errors"

var ErrBotTokenRequired = errors.New("bot token is required")
var ErrInvalidSelectionMessage = errors.New("invalid selection message")
var ErrSendFailed = errors.New("send failed")
