package telegram

import (
	"context"
	"errors"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
	"gopkg.in/telebot.v3"
)

// Start starts your Telegram Service
func (s *Service) Start(ctx context.Context, onMsg func(msg *chat.Message)) error {

	// validate
	if s.cfg.BotToken == "" {
		return ErrBotTokenRequired
	}

	stopDone := make(chan struct{}, 1) // make a channel to align with

	// set the bot and start it.
	settings := telebot.Settings{
		Token:  s.cfg.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, c telebot.Context) {
			color.Red("[Manboster Telegram Provider] We encountered an error: %q", err)
			stopDone <- struct{}{}
		},
	}

	// start the bot
	b, err := telebot.NewBot(settings)
	if err != nil {
		return err
	}
	s.tgInstance = b

	go func() {
		select {
		case <-ctx.Done():
			color.Yellow("[Manboster Telegram Provider] Context cancelled, shutting down...")
			s.tgInstance.Stop()
		case <-stopDone:
			s.tgInstance.Stop()
		}
	}()

	// Handler for Message Resp calling.
	s.tgInstance.Handle(telebot.OnText, func(c telebot.Context) error {
		return s.HandleText(ctx, c, onMsg)
	})
	s.tgInstance.Handle(telebot.OnCallback, func(c telebot.Context) error { return s.HandleCallback(ctx, c, onMsg) })

	color.Blue("[Manboster Telegram Provider] Starting the telegram bot...")

	go s.tgInstance.Start()

	select {
	case <-ctx.Done():
		return nil
	case <-stopDone:
		return errors.New("telegram provider is facing an problem")
	}
}

func (s *Service) Type(ctx context.Context, chatId telebot.ChatID) {
	// send immediately
	_ = s.tgInstance.Notify(chatId, telebot.Typing)

	ticker := time.NewTicker(4 * time.Second) // send every 4 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_ = s.tgInstance.Notify(chatId, telebot.Typing)
		}
	}
}
