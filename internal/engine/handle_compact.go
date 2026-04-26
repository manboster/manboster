package engine

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/engine/chatdata"
	"github.com/manboster/manboster/spec/chat"
)

// HandleCompact compacts data
func (e *Engine) HandleCompact(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	text := ""
	respMessage.Text = &chat.TextPayload{
		Text: fmt.Sprintf("Compacting session `%s`, please wait for a while...", sessionId),
	}
	err := instance.SendMessage(ctx, respMessage)
	if err != nil {
		return err
	}

	err = e.chatDataService.Compact(ctx, instance, msg, sessionId)
	if err != nil {
		if errors.Is(err, chatdata.ErrNoNeedToCompact) {
			text = "The context is too small now and there is no need to compact."
		} else {
			color.Yellow(fmt.Sprintf("[Manboster Engine] We encountered an error when compacting message: %q", err))
			text = "We encountered an error when compacting this session's message."
		}
		respMessage.MessageType = chat.MessageText
		respMessage.Text = &chat.TextPayload{
			Text: text,
		}
		return instance.SendMessage(ctx, respMessage)
	}
	return nil
}
