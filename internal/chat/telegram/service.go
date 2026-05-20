package telegram

import (
	"sync"

	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/config"
	"gopkg.in/telebot.v3"
)

type Service struct {
	tgInstance *telebot.Bot
	cfg        *Config
	sendMutex  sync.Mutex
	manager    *Manager
}

func (s *Service) New() chat.Provider {
	return &Service{}
}

func (s *Service) Name() string {
	return "telegram"
}

func (s *Service) DisplayName() string {
	return "Telegram"
}

func (s *Service) Ability() chat.AbilityType {
	// return chat.AbilityAll
	// Now select function is a problem so we deleted without select info
	return chat.AbilityNoSelect
}

func (s *Service) Config() config.Provider {
	return &Config{}
}
