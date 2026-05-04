package engine

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

func (e *Engine) MessageHandler(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	color.Blue("[Manboster Engine] Now handling message...")

	// now, notify process!
	err := instance.Notify(ctx, msg, chat.ActionPending)
	color.Blue("[Manboster Engine] Notified provider pending status")
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Engine] Error while notifying provider %q: %q", instance.DisplayName(), err))
	}

	if err := ctx.Err(); err != nil {
		return err
	}

	message, err := e.soulService.BuildLLMMessage(ctx, msg, sessionId, e.safeguardService.UserType(ctx, instance.Name(), msg.UserID))
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] Error while building LLM message: %q", err))
		return err
	}
	// enhanced prompt engineering in order to avoid injection with some effort.
	msgData := llm.Event{
		EventType: llm.EventMessage,
		Message:   message,
	}
	e.sessionService.Manager.ChatSession.AppendEvent(sessionId, msgData)
	err = e.chatDataService.Write(ctx, msgData, sessionId)
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to write message data to repository, your chat data would not be saved! sessionId: %s, chatId: %s, provider: %s, error: %q", sessionId, msg.ChatID, instance.Name(), err))
	}

	provider, model, _ := e.sessionService.Manager.ChatSession.GetModel(sessionId)
	msgList := e.sessionService.Manager.ChatSession.GetMessages(sessionId)
	p, m := util.GetModelWithFallback(ctx, e.llmProviders, provider, model)
	respMessage := msg.Clone()

	var souls []string
	s, avail := e.sessionService.Manager.ChatSession.GetSoul(sessionId)
	if !avail {
		souls = []string{
			"system",
		}
	} else {
		souls = s
	}
	soulLLMMsg, err := e.soulService.BuildSystemMessage(ctx, souls)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] Error while building system message: %q", err))
	}
	msgList = append([]llm.Message{soulLLMMsg}, msgList...)

	maxCount := 3
	maxRepeatCount := 5

	count := 0

	repeatCount := 0
	repeatFingerPrint := ""

	repeatSingleMaxCount := 5
	repeatSingleFingerprint := map[string]int{}

	isOverflow := false

	for {
		event, err := e.gateway.LLMChat(ctx, p, m, msgList)
		errChat := e.gateway.HandleLLMChatError(ctx, instance, msg, p.DisplayName(), m.DisplayName, err)
		if errChat != nil {
			color.Yellow(fmt.Sprintf("[Manboster Engine] Error while sending message: %q\n", errChat))
		}
		if err != nil {
			break
		}

		// renew the map
		repeatSingleFingerprint = map[string]int{}

		if count >= maxCount || repeatCount >= maxRepeatCount || isOverflow {
			respMsg := msg.Clone()
			respMsg.MessageType = chat.MessageText
			respMsg.Text = &chat.TextPayload{
				Text: "Models tried the same tool call too much times or model calls return failed too much times, we broke it out in order to avoid consuming useless tokens and time.",
			}
			return e.gateway.SendMessage(ctx, instance, respMsg)
		}

		// write the whole event immediately
		e.sessionService.Manager.ChatSession.AppendEvent(sessionId, *event)
		msgList = append(msgList, *event.Message)

		// add model cost
		util.CalculateCost(event, m)
		if event.EventType&llm.EventMessage != 0 && event.Message != nil && event.Message.Type&llm.MessageToolCallRequest == 0 {
			err = e.chatDataService.Write(ctx, *event, sessionId)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to write message data to repository, your chat data would not be saved! sessionId: %s, chatId: %s, provider: %s, error: %q", sessionId, respMessage.ChatID, instance.Name(), err))
			}
		}

		// handle text message
		if event.EventType&llm.EventMessage != 0 && event.Message != nil && len(event.Message.Parts) > 0 {
			err = e.handler.HandleText(ctx, instance, msg, *event)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Error while sending message: %q\n", err))
			}
		}

		if event.EventType&llm.EventMessage != 0 && event.Message != nil && event.Message.Type&llm.MessageToolCallRequest == 0 {
			break
		}

		// handling tool call request
		if event.EventType&llm.EventMessage != 0 && event.Message != nil && event.Message.Type&llm.MessageToolCallRequest != 0 {
			var toolNameArgs []string

			for _, req := range event.Message.ToolCallRequest {
				a := fmt.Sprintf("%s:%v", req.ToolName, req.ToolArgs)
				toolNameArgs = append(toolNameArgs, a)
				_, avail := repeatSingleFingerprint[a]
				if !avail {
					repeatSingleFingerprint[a] = 1
				}
				repeatSingleFingerprint[a]++
			}

			isOverflow = false
			for _, co := range repeatSingleFingerprint {
				if co > repeatSingleMaxCount {
					isOverflow = true
					break
				}
			}
			if isOverflow {
				continue
			}

			if repeatFingerPrint == strings.Join(toolNameArgs, "%") {
				repeatCount++
			} else {
				repeatCount = 0
				repeatFingerPrint = strings.Join(toolNameArgs, "%")
			}

			respEvent, successExecution, err := e.handler.HandleToolCall(ctx, instance, msg, *event, sessionId)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Error while sending message: %q\n", err))
			}
			if !successExecution {
				count++
			} else {
				count = 0
			}

			e.sessionService.Manager.ChatSession.AppendEvent(sessionId, respEvent)
			msgList = append(msgList, *respEvent.Message)

			// it should be written with tool req and tool resp.
			err = e.chatDataService.Write(ctx, *event, sessionId)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to write message data to repository, your chat data would not be saved! sessionId: %s, chatId: %s, provider: %s, error: %q", sessionId, respMessage.ChatID, instance.Name(), err))
			}
			// in order to avoid complexity?
			err = e.chatDataService.Write(ctx, respEvent, sessionId)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to write message data to repository, your chat data would not be saved! sessionId: %s, chatId: %s, provider: %s, error: %q", sessionId, respMessage.ChatID, instance.Name(), err))
			}
		}

	}

	return nil
}
