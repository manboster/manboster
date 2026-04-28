package telegram

import (
	"context"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
	"gopkg.in/telebot.v3"
)

func (s *Service) HandleCallback(ctx context.Context, c telebot.Context, onMsg func(msg *chat.Message)) error {
	var msg *chat.Message
	msg = msg.Build(&Service{})
	color.Blue("[Manboster Telegram Provider] Received a callback")
	s.msgBaseParser(msg, c)
	//jsonify, _ := json.MarshalIndent(c.Callback(), "", " ")
	//fmt.Println(string(jsonify))

	d := c.Callback().Data
	if strings.Contains(d, "\f") {
		d = strings.Replace(d, "\f", "", -1)
		dList := strings.Split(d, "|")
		// fmt.Println(c.Callback().Sender.ID)
		if len(dList) == 2 {
			msg.MessageType = chat.MessageSelectionCallback
			msg.SelectionCallback = &chat.SelectionCallbackPayload{
				SelectionSessionId: dList[1],
				SelectionValue:     dList[0],
				SelectionBy:        strconv.FormatInt(c.Callback().Sender.ID, 10),
			}
			err := s.tgInstance.Delete(c.Callback().Message)
			if err != nil {
				color.Yellow("[Manboster Telegram Provider] Failed to delete reply message")
			}
			// c.DeleteAfter(30 * time.Second)
			err = s.tgInstance.Respond(c.Callback(), &telebot.CallbackResponse{
				CallbackID: c.Callback().ID,
				Text:       "Manboster Processing...",
			})
			if err != nil {
				return err
			}
			onMsg(msg)
		} else {
			color.Yellow("[Manboster Telegram Provider] Could not parse callback data message")
		}
	}

	return nil
}
