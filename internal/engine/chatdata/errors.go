package chatdata

import "errors"

var ErrNoNeedToCompact = errors.New("no need to compact")
var ErrCompactChatFailed = errors.New("compact chat failed")
