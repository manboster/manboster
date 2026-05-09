package telegram

import (
	"context"
	"fmt"
	"strconv"

	"github.com/manboster/manboster/spec/chat"
	"gopkg.in/telebot.v3"
)

// EditMessage edits a message specified
func (s *Service) EditMessage(ctx context.Context, msg *chat.Message) error {
	if !s.manager.Avail() {
		return fmt.Errorf("telegram instance is currently available")
	}
	if msg.MessageType&chat.MessageUnknown == 0 {
		return ErrInvalidMessageType
	}

	msgId := 0
	msgId, err := strconv.Atoi(msg.MessageID)
	if err != nil {
		return err
	}
	chatId := int64(0)
	chatId, err = strconv.ParseInt(msg.ChatID, 10, 64)
	if err != nil {
		return err
	}

	var m = &telebot.Message{
		ID: msgId,
		Chat: &telebot.Chat{
			ID: chatId,
		},
	}

	if msg.MessageType&chat.MessageText != 0 {
		_, err = s.tgInstance.Edit(m, s.Converter(msg.Text.Text, false, false), &telebot.SendOptions{
			ParseMode: telebot.ModeHTML,
		})
	}

	return err
}
