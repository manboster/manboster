package telegram

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
	"gopkg.in/telebot.v3"
)

func (s *Service) Handler(ctx context.Context, c telebot.Context, onMsg func(msg *chat.Message)) error {
	// var msg = chat.BuildMessage(s.New())
	var msg *chat.Message
	msg = msg.Build(&Service{})

	color.Cyan("[Manboster Telegram Provider] Got a message.")
	err := s.msgBaseParser(msg, c, onMsg)
	if errors.Is(err, ErrImageNoNeedToTrigger) {
		return nil
	}

	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Telegram Provider] Error parsing message: %s", err.Error()))
		return err
	}
	// call onMsg on index
	onMsg(msg)

	return nil
}
