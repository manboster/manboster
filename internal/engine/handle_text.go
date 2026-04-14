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
	"github.com/manboster/manboster/internal/util"
)

// HandleText handles text messages.
func (e *Engine) HandleText(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	color.Blue("[Manboster Engine] Now handling text message...")

	msg.MessageType = chat.MessageText
	sessionId := e.sessionManager.ID(instance.Name(), msg.ChatID)

	// TODO: replace it to channel queue
	lock := e.sessionManager.GetSessionLocks(sessionId)
	lock.Lock()
	defer lock.Unlock()

	if err := ctx.Err(); err != nil {
		return err
	}

	// make a cancelable context
	cancelCtx, cancel := context.WithCancel(ctx)
	defer func(sid string) {
		cancel()
		e.sessionManager.Deactivate(sid)
	}(sessionId)
	e.sessionManager.Activate(sessionId, cancel)
	_, avail := e.sessionManager.GetSession(sessionId)
	if !avail {
		e.sessionManager.AppendMessage(sessionId, llm.Message{
			Role: llm.RoleSystem,
			Text: &llm.MessageTextPayload{
				Text: config.InitialSystemPrompt, // TODO: prompt engineering
			},
			Type: llm.MessageText,
		})
	}

	// who say...
	chatName := "(Private Chat)"
	if msg.ChatType != chat.ChatsPersonal {
		chatName = msg.ChatName
	}
	promptTxt := fmt.Sprintf("%s said in %s, [%s]:\n%s", msg.Username, chatName, msg.CreatedAt, msg.Text.Text)
	msgData := llm.Message{
		Role: llm.RoleUser,
		Text: &llm.MessageTextPayload{
			Text: promptTxt,
		},
		Type: llm.MessageText,
	}
	e.sessionManager.AppendMessage(sessionId, msgData)

	tries := 1
	var event *llm.Event
	var err error
	// try 3 times
	for tries <= 3 {
		if cancelCtx.Err() != nil {
			return cancelCtx.Err()
		}

		color.Blue(fmt.Sprintf("[Manboster Engine] Fetching message response from LLMProvider %s, try %d times", e.llmProviders[0].Name(), tries))
		// we make timeout requests.
		timeoutCtx, cancel := context.WithTimeout(cancelCtx, 2*time.Minute)
		models := e.llmProviders[0].Models()
		data, _ := e.sessionManager.GetSession(sessionId)
		event, err = e.llmProviders[0].Chat(timeoutCtx, models[0].Name, data.Messages)
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
		color.Red(fmt.Sprintf("[Manboster Engine] Failed to get message from LLMProvider %s after 3 tries, get error: %s", e.llmProviders[0].Name(), err.Error()))
		// now we have to wrap this into friendly prompt
		tips := fmt.Sprintf("%+v", err)
		text := fmt.Sprintf("[Manboster] Failed to get message from LLMProvider %s after trying 3 times, get error: %s\nYou can resend your message or check the API's availability.", e.llmProviders[0].Name(), err.Error())
		if strings.Contains(tips, "429") {
			text = fmt.Sprintf("[Manboster] %s has been suffering a very high traffic and triggered rate limit, please try again later or change provider's models.", e.llmProviders[0].Name())
		} else if strings.Contains(tips, "500") || strings.Contains(tips, "502") || strings.Contains(tips, "503") || strings.Contains(tips, "501") {
			text = fmt.Sprintf("[Manboster] %s has been down, please check your provider's status page, or change providers and try again later.", e.llmProviders[0].Name())
		} else if strings.Contains(tips, "context deadline exceeded") {
			text = fmt.Sprintf("[Manboster] It seems that there is a connection issue between you and provider %s, please check your internet connection and try again.", e.llmProviders[0].Name())
		} else if strings.Contains(tips, "403") || strings.Contains(tips, "401") {
			text = fmt.Sprintf("[Manboster] Access denied or unauthorized in provider %s, please check your API key or other credentials is valid.", e.llmProviders[0].Name())
		} else if strings.Contains(tips, "cancel") {
			text = fmt.Sprintf("[Manboster] You cancelled provider %s's request.", e.llmProviders[0].Name())
		}
		msg.Text = &chat.TextPayload{
			Text: text,
		}
	} else {
		textWithoutThinking := util.StripThink(event.Message.Text.Text)
		e.sessionManager.AppendMessage(sessionId, llm.Message{
			Text: &llm.MessageTextPayload{
				Text: textWithoutThinking,
			},
			Role: event.Message.Role,
			Type: llm.MessageText,
		})

		if util.ExtractThinkContent(event.Message.Text.Text) != "" {
			msg.MessageType = chat.MessageThinkingText
			msg.Text = &chat.TextPayload{
				Text: util.ExtractThinkContent(event.Message.Text.Text),
			}
			err := instance.SendMessage(ctx, msg)
			if err != nil {
				return err
			}
		}

		msg.MessageType = chat.MessageText
		msg.Text = &chat.TextPayload{
			Text: textWithoutThinking,
		}
	}

	tries = 1
	for tries <= 5 {
		if cancelCtx.Err() != nil {
			return cancelCtx.Err()
		}

		timeoutCtx, cancel := context.WithTimeout(cancelCtx, 10*time.Second)
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
