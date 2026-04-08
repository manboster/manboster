package engine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
)

// HandleText handles text messages.
func (e *Engine) HandleText(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	color.Blue("[Manboster Engine] Now handling text message...")

	msg.MessageType = chat.MessageText
	sessionId := fmt.Sprintf("%s:%s", instance.Name(), msg.ChatID)
	sessionData := e.sessionManager.GetSession(sessionId)
	if len(sessionData.Messages) == 0 {
		sessionData.Messages = append(sessionData.Messages, llm.Message{
			Role: llm.RoleSystem,
			Text: config.InitialSystemPrompt, // TODO: prompt engineering
			Type: llm.MessageText,
		})
	}
	msgData := append(sessionData.Messages, llm.Message{
		Role: llm.RoleUser,
		Text: msg.Text.Text,
		Type: llm.MessageText,
	})

	tries := 1
	var event *llm.Event
	var err error
	// try 5 times
	for tries <= 5 {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		// we make timeout requests.
		timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		event, err = e.llmProviders[0].Chat(timeoutCtx, msgData)
		cancel()

		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Failed to get message from LLMProvider %s after %d tries, get error: %s", e.llmProviders[0].Name(), tries, err.Error()))
			time.Sleep(time.Second * time.Duration(tries+1))
			tries++
		} else {
			color.Blue(fmt.Sprintf("[Manboster Engine] Got message feedback from LLMProvider %s", e.llmProviders[0].Name()))
			break
		}
	}
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] Failed to get message from LLMProvider %s after 5 tries, get error: %s", e.llmProviders[0].Name(), err.Error()))
		// now we have to wrap this into friendly prompt
		tips := fmt.Sprintf("%+v", err)
		if strings.Contains(tips, "429") {
			msg.Text = &chat.TextPayload{
				Text: fmt.Sprintf("[Manboster] %s has been suffering a very high traffic and triggered rate limit, please try again later or change provider's models.", e.llmProviders[0].Name()),
			}
		} else if strings.Contains(tips, "500") || strings.Contains("502", tips) || strings.Contains("503", tips) || strings.Contains("501", tips) {
			msg.Text = &chat.TextPayload{
				Text: fmt.Sprintf("[Manboster] %s has been down, please check your provider's status page, or change providers and try again later.", e.llmProviders[0].Name()),
			}
		} else if strings.Contains(tips, "context deadline exceeded") {
			msg.Text = &chat.TextPayload{
				Text: fmt.Sprintf("[Manboster] It seems that there is a connection issue between you and provider %s, please check your internet connection and try again.", e.llmProviders[0].Name()),
			}
		} else if strings.Contains(tips, "403") || strings.Contains(tips, "401") {
			msg.Text = &chat.TextPayload{
				Text: fmt.Sprintf("[Manboster] Access denied or unauthorized in provider %s, please check your API key or other credentials is valid.", e.llmProviders[0].Name()),
			}
		} else {
			msg.Text = &chat.TextPayload{
				Text: fmt.Sprintf("[Manboster Engine] Failed to get message from LLMProvider %s after trying 5 times, get error: %s\nYou can resend your message or check the API's availability.", e.llmProviders[0].Name(), err.Error()),
			}
		}
	} else {
		msg.Text = &chat.TextPayload{
			Text: event.Message.Text,
		}
		msgData = append(msgData, llm.Message{
			Text: event.Message.Text,
			Role: event.Message.Role,
			Type: llm.MessageText,
		})
	}

	sessionData.Messages = msgData
	e.sessionManager.SetSession(sessionId, sessionData)

	tries = 1
	for tries <= 5 {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		err = instance.SendMessage(timeoutCtx, msg)
		cancel()

		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Tried %d times sending via %s, got error: %q", tries, instance.Name(), err))
			time.Sleep(time.Second * time.Duration(tries+1))
			tries++
			continue
		} else {
			color.Green(fmt.Sprintf("[Manboster Engine] Tried %d times sending via %s, success.", tries, instance.Name()))
			return nil
		}
	}
	return err
}
