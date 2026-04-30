package telegram

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/spec/chat"
	"gopkg.in/telebot.v3"
)

// DeleteMessage deletes the message
func (s *Service) DeleteMessage(ctx context.Context, msg *chat.Message) error {
	if msg.MessageType&chat.MessageUnknown == 0 {
		return ErrInvalidMessageType
	}

	msgId := 0
	_, err := fmt.Sscanf(msg.ChatID, "%d", &msgId)
	if err != nil {
		return err
	}
	chatId := 0
	_, err = fmt.Sscanf(msg.ChatID, "%d", &chatId)
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
