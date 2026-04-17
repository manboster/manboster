package telegram

import (
	"github.com/manboster/manboster/internal/chat"
)

func init() {
	chat.Register("telegram", func() chat.Provider {
		return &Service{}
	})
}
