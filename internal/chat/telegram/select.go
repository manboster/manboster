package telegram

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/chat"
	"gopkg.in/telebot.v3"
)

// Select give user a plenty of selections and wait for them to reply.
func (s *Service) Select(ctx context.Context, title string, message *chat.Message, selection []chat.Selection) (string, error) {
	menu := &telebot.ReplyMarkup{}

	recp, err := recipientParser(message.ChatID)
	if err != nil {
		return "", err
	}

	// define buttons
	var btns []telebot.Btn
	for _, slc := range selection {
		btn := menu.Data(slc.Name, slc.Value, "select:"+message.ChatID+":"+title)
		btns = append(btns, btn)
	}
	menu.Inline(menu.Split(3, btns)...)

	// send menu selection
	send, err := s.tgInstance.Send(recp, menu)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%d", message.ChatID, send.ID), nil
}
