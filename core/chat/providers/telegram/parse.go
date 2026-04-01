package telegram

import (
	"fmt"
	"strconv"

	"gopkg.in/telebot.v3"
)

func recipientParser(chatID string) (telebot.ChatID, error) {
	id, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return telebot.ChatID(0), fmt.Errorf("invalid chat id: %w", err)
	}

	// define recipients
	recp := telebot.ChatID(id)
	return recp, nil
}
