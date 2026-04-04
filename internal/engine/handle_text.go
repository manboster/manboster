package engine

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/llm"
)

// HandleText handles text messages.
func (e *Engine) HandleText(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	sessionId := fmt.Sprintf("%s:%s", instance.Name(), msg.ChatID)
	sessionData := e.sessionManager.GetSession(sessionId)
	if len(sessionData.Messages) == 0 {
		sessionData.Messages = append(sessionData.Messages, llm.Message{
			Role: llm.RoleTypeSystem,
			Text: "You're an assistant named Manboster. You are chatting with people. The one who is chatting with you is your owner.", // TODO: prompt engineering
			Type: llm.MessageTypeText,
		})
	}
	msgData := append(sessionData.Messages, llm.Message{
		Role: llm.RoleTypeUser,
		Text: msg.Text.Text,
		Type: llm.MessageTypeText,
	})

	tries := 0
	var mesg *llm.Message
	var err error
	// try 5 times
	for tries < 5 {
		mesg, err = e.llmProviders[0].Chat(ctx, msgData)
		if err != nil {
			color.Red("Failed to get message from LLMProvider %s after %d tries, get error: %s", e.llmProviders[0].Name(), tries+1, err.Error())
			tries++
		} else {
			break
		}
	}
	if err != nil {
		color.Red(fmt.Sprintf("Failed to get message from LLMProvider %s after 5 tries, get error: %s", e.llmProviders[0].Name(), err.Error()))
		msg.Text = &chat.TextPayload{
			Text: fmt.Sprintf("[Manboster]Failed to get message from LLMProvider %s after trying 5 times, get error: %s\nYou can resend your message or check the API's availability.", e.llmProviders[0].Name(), err.Error()),
		}
	} else {
		msg.Text = &chat.TextPayload{
			Text: mesg.Text,
		}
		msgData = append(msgData, llm.Message{
			Text: mesg.Text,
			Role: mesg.Role,
			Type: llm.MessageTypeText,
		})
	}

	sessionData.Messages = msgData
	e.sessionManager.SetSession(sessionId, sessionData)

	err = instance.SendMessage(ctx, msg)
	if err != nil {
		color.Red(err.Error())
		return nil
	}
	return nil
}
