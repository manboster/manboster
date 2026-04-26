package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

// HandleText handles text messages.
func (e *Engine) HandleText(ctx context.Context, instance chat.Provider, msg *chat.Message, event llm.Event) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	text := event.Message.Parts[0].Text.Text
	// fmt.Println(text)
	textWithoutThinking := util.StripThink(text)
	// fmt.Println(textWithoutThinking)

	// If there is a thinking context
	if util.ExtractThinkContent(text) != "" {
		respMessage.MessageType = chat.MessageThinkingText
		respMessage.Text = &chat.TextPayload{
			Text: util.ExtractThinkContent(text),
		}
		err := e.SendMessage(ctx, instance, respMessage)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Engine] Error while sending message: %q\n", err))
		}
	}

	respMessage.MessageType = chat.MessageText
	respMessage.Text = &chat.TextPayload{
		Text: textWithoutThinking,
	}
	return e.SendMessage(ctx, instance, respMessage)
}
