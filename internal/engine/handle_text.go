package engine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/util"
)

// HandleText handles text messages.
func (e *Engine) HandleText(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	color.Blue("[Manboster Engine] Now handling text message...")

	// now, notify process!
	err := instance.Notify(ctx, msg, chat.ActionPending)
	color.Blue("[Manboster Engine] Notified provider pending status")
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Engine] Error while notifying provider %s: %q", instance.DisplayName(), err))
	}

	respMessage := msg.Clone()

	if err := ctx.Err(); err != nil {
		return err
	}

	message, err := e.promptService.BuildLLMMessage(ctx, msg, sessionId, e.safeguardService.UserType(ctx, instance.Name(), msg.UserID))
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] Error while building LLM message: %q", err))
		return err
	}
	// enhanced prompt engineering in order to avoid injection with some effort.
	msgData := llm.Event{
		EventType: llm.EventMessage,
		Message:   message,
	}
	e.sessionManager.AppendEvent(sessionId, msgData)
	err = e.chatDataService.Write(ctx, msgData, sessionId)
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to write message data to repository, your chat data would not be saved! sessionId: %s, chatId: %s, provider: %s, error: %q", sessionId, msg.ChatID, instance.Name(), err))
	}

	var event *llm.Event

	provider, model, _ := e.sessionManager.GetModel(sessionId)
	msgList := e.sessionManager.GetMessages(sessionId)
	pIndex, mIndex := util.GetModelIndexWithFallback(ctx, e.llmProviders, provider, model)
	llmProviderDisplayName := e.llmProviders[pIndex].DisplayName()
	llmModelDisplayName := e.llmProviders[pIndex].Models()[mIndex].DisplayName
	llmModelName := e.llmProviders[pIndex].Models()[mIndex].Name

	totToken, err := e.repo.GetTotalToken(ctx, sessionId)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] Error while getting total tokens from repository: %q", err))
	}

	if uint64(totToken) > llm.CalculateCompactTokens(e.llmProviders[pIndex].Models()[mIndex]) {
		err := e.HandleCompact(ctx, instance, msg, sessionId)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Error while compacting data: %q", err))
			return err
		}
		// get new session id
		resp, err := e.repo.GetChat(ctx, msg.ChatID, instance.Name())
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Error while getting new session id: %q", err))
			return err
		}
		sessionId = resp.SessionID
	}
	// fmt.Printf("%+q \n%+q", data.Events, msgList)

	var errTrying error = nil
	// try 3 times
	times := 3
	// tries def
	tries := 1
	for tries <= times {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		color.Blue(fmt.Sprintf("[Manboster Engine] Fetching message response from LLMProvider %s, model %s, try %d times", llmProviderDisplayName, llmModelDisplayName, tries))
		// we make timeout requests.
		timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)

		event, errTrying = e.llmProviders[pIndex].Chat(timeoutCtx, llmModelName, msgList)

		cancel()

		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Failed to get message from LLMProvider %s, model %s after %d tries, get error: %s", llmProviderDisplayName, llmModelDisplayName, tries, err.Error()))
			time.Sleep(time.Second * time.Duration(tries+1))
			tries++
		} else {
			color.Blue(fmt.Sprintf("[Manboster Engine] Got message feedback from LLMProvider %s, model %s", llmProviderDisplayName, llmModelDisplayName))
			break
		}
	}
	if errTrying != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] Failed to get message from LLMProvider %s, model %s, after %d tries, get error: %s", llmProviderDisplayName, llmModelDisplayName, times, err.Error()))
		// now we have to wrap this into friendly prompt
		tips := fmt.Sprintf("%+v", err)
		text := fmt.Sprintf("[Manboster] Failed to get message from LLMProvider %s, model %s after trying %d times, get error: %s\nYou can resend your message or check the API's availability.", llmProviderDisplayName, llmModelDisplayName, times, err.Error())
		if strings.Contains(tips, "429") {
			text = fmt.Sprintf("[Manboster] Provider %s, Model %s has been suffering a very high traffic and triggered rate limit, please try again later or change provider's models.", llmProviderDisplayName, llmModelDisplayName)
		} else if strings.Contains(tips, "500") || strings.Contains(tips, "502") || strings.Contains(tips, "503") || strings.Contains(tips, "501") {
			text = fmt.Sprintf("[Manboster] Provider %s has been down, please check your provider's status page, or change providers and try again later.", llmProviderDisplayName)
		} else if strings.Contains(tips, "context deadline exceeded") {
			text = fmt.Sprintf("[Manboster] It seems that there is a connection issue between you and provider %s, please check your internet connection and try again.", llmProviderDisplayName)
		} else if strings.Contains(tips, "403") || strings.Contains(tips, "401") {
			text = fmt.Sprintf("[Manboster] Access denied or unauthorized in provider %s, please check your API key or other credentials is valid.", llmProviderDisplayName)
		} else if strings.Contains(tips, "cancel") {
			text = fmt.Sprintf("[Manboster] You cancelled provider %s's request.", llmProviderDisplayName)
		}
		respMessage.Text = &chat.TextPayload{
			Text: text,
		}
	} else {
		if event.EventType&llm.EventMessage != 0 && event.Message != nil && len(event.Message.Parts) > 0 {
			text := event.Message.Parts[0].Text.Text
			// fmt.Println(text)
			textWithoutThinking := util.StripThink(text)
			// fmt.Println(textWithoutThinking)
			e.sessionManager.AppendEvent(sessionId, *event)
			err := e.chatDataService.Write(ctx, *event, sessionId)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to write message data to repository, your chat data would not be saved! sessionId: %s, chatId: %s, provider: %s, error: %q", sessionId, respMessage.ChatID, instance.Name(), err))
			}

			if util.ExtractThinkContent(text) != "" {
				respMessage.MessageType = chat.MessageThinkingText
				respMessage.Text = &chat.TextPayload{
					Text: util.ExtractThinkContent(text),
				}
				err := instance.SendMessage(ctx, respMessage)
				if err != nil {
					return err
				}
			}

			respMessage.MessageType = chat.MessageText
			respMessage.Text = &chat.TextPayload{
				Text: textWithoutThinking,
			}
		}
	}

	tries = 1
	for tries <= 5 {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		err = instance.SendMessage(timeoutCtx, respMessage)
		cancel()

		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Tried %d times sending via %s, got error: %q", tries, instance.Name(), err))
			time.Sleep(time.Second * time.Duration(tries+1))
			tries++
			continue
		} else {
			color.Green(fmt.Sprintf("[Manboster Engine] Tried %d times sending via %s, success.", tries, instance.Name()))
			err := instance.Notify(ctx, msg, chat.ActionSuccess)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return err
}
