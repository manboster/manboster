package telegram

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
	"gopkg.in/telebot.v3"
)

func (s *Service) HandleCallback(ctx context.Context, c telebot.Context, onMsg func(msg *chat.Message)) error {
	var msg *chat.Message
	msg = msg.Build(&Service{})
	s.msgBaseParser(msg, c)
	msg.MessageType = chat.MessageSelectionCallback
	msg.SelectionCallback = &chat.SelectionCallbackPayload{
		SelectionSessionId: c.Callback().Data,
		SelectionValue:     c.Callback().Unique,
	}
	onMsg(msg)
	return nil
}
