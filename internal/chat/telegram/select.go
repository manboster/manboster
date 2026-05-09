package telegram

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"gopkg.in/telebot.v3"
)

// Select give user a plenty of selections and wait for them to reply.
func (s *Service) Select(ctx context.Context, sessionId string, message *chat.Message) error {
	if !s.manager.Avail() {
		return fmt.Errorf("telegram instance is currently available")
	}
	if message.MessageType&(chat.MessageSelection|chat.MessageTextImageAndFile) == 0 {
		return ErrInvalidMessageType
	}
	menu := &telebot.ReplyMarkup{}

	recp, err := recipientParser(message.ChatID)
	if err != nil {
		return err
	}

	// define buttons
	var btns []telebot.Btn
	if message.Selection == nil {
		return ErrInvalidSelectionMessage
	}

	maxLen := 0
	for _, slc := range message.Selection.Selection {
		btn := menu.Data(slc.Name, slc.Value, sessionId)
		btns = append(btns, btn)
		if maxLen < len(slc.Name) {
			maxLen = len(slc.Name)
		}
	}

	if maxLen < 10 {
		menu.Inline(menu.Split(2, btns)...)
	} else {
		menu.Inline(menu.Split(1, btns)...)
	}

	text, err := util.EscapeMarkdownToTelegramHTML(message.Text.Text)
	// send menu selection
	send, err := s.tgInstance.Send(recp, text, &telebot.SendOptions{
		ReplyMarkup: menu,
		ParseMode:   telebot.ModeHTML,
	})
	if err != nil {
		return err
	}
	message.MessageID = fmt.Sprintf("%d", send.ID)
	// jsonify, _ := json.MarshalIndent(message, "", " ")
	// fmt.Println(string(jsonify))
	return nil
}
