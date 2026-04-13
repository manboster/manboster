package telegram

import (
	"context"
	"errors"
	"time"

	"github.com/fatih/color"
	"github.com/go-viper/mapstructure/v2"
	"github.com/manboster/manboster/internal/chat"
	"gopkg.in/telebot.v3"
)

// Start starts your Telegram Service
func (s *Service) Start(ctx context.Context, conf any, onMsg func(msg *chat.Message)) error {
	// get config
	var cfg Config
	err := mapstructure.Decode(conf, &cfg)
	if err != nil {
		return err
	}

	// validate
	if cfg.BotToken == "" {
		return ErrBotTokenRequired
	}

	// set the bot and start it.
	settings := telebot.Settings{
		Token:  cfg.BotToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		OnError: func(err error, c telebot.Context) {
			color.Red("[Manboster Telegram Provider] We encountered an error: %v", err)
		},
	}

	// start the bot
	b, err := telebot.NewBot(settings)
	if err != nil {
		return err
	}
	s.tgInstance = b

	stopDone := make(chan struct{}) // make a channel to align with
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

	color.Blue("[Manboster Telegram Provider] Starting the telegram bot...")

	s.tgInstance.Start()

	select {
	case <-ctx.Done():
		return nil
	case <-stopDone:
		return errors.New("telegram provider is facing an problem")
	}
}

func (s *Service) Notify(ctx context.Context, chatID string, action chat.ActionType) error {
	recipient, err := recipientParser(chatID)
	if err != nil {
		return err
	}
	return s.tgInstance.Notify(recipient, telebot.ChatAction(action))
}

func (s *Service) Type(chatId telebot.ChatID, ctx context.Context) {
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
