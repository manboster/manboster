package telegram

import (
	"github.com/manboster/manboster/internal/chat"
	chatType "github.com/manboster/manboster/spec/chat"
)

func init() {
	chat.Register("telegram", func() chatType.Provider {
		return &Service{}
	})
}
