package telegram

import "gopkg.in/telebot.v3"

type Service struct {
	tgInstance *telebot.Bot
}

func NewService(tgInstance *telebot.Bot) *Service {
	return &Service{
		tgInstance: tgInstance,
	}
}
