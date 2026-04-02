package telegram

import (
	"github.com/manboster/manboster/internal/chat"
	"gopkg.in/telebot.v3"
)

func init() {
	chat.Register("telegram", func() chat.Provider {
		return &Service{}
	})
}

type Service struct {
	tgInstance *telebot.Bot
}

func NewService(tgInstance *telebot.Bot) *Service {
	return &Service{
		tgInstance: tgInstance,
	}
}

func (s *Service) New() chat.Provider {
	return &Service{}
}

func (s *Service) Name() string {
	return "telegram"
}
