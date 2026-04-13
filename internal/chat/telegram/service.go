package telegram

import (
	"sync"

	"github.com/manboster/manboster/internal/chat"
	"gopkg.in/telebot.v3"
)

type Service struct {
	tgInstance *telebot.Bot
	sendMutex  sync.Mutex
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

func (s *Service) Ability() chat.AbilityType {
	// return chat.AbilityAll
	// Now select function is a problem so we deleted without select info
	return chat.AbilityNoSelect
}
