package telegram

import (
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
)

func init() {
	chat.Register("telegram", func() chat.Provider {
		return &Service{}
	})
	config.Register("chat:telegram", func() config.Provider {
		return &Config{}
	})
}
