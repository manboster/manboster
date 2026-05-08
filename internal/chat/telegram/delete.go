package telegram

import (
	"context"
	"strconv"

	"github.com/manboster/manboster/spec/chat"
	"gopkg.in/telebot.v3"
)

// DeleteMessage deletes the message
func (s *Service) DeleteMessage(ctx context.Context, msg *chat.Message) error {
	if msg.MessageType&chat.MessageUnknown == 0 {
		return ErrInvalidMessageType
	}

	msgId, err := strconv.Atoi(msg.MessageID)
	if err != nil {
		return err
	}
	chatId, err := strconv.Atoi(msg.ChatID)
	if err != nil {
		return err
	}

	return s.tgInstance.Delete(&telebot.Message{
		ID: msgId,
		Chat: &telebot.Chat{
			ID: int64(chatId),
		},
	})
}
