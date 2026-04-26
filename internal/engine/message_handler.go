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
	e.sessionManager.AppendEvent(sessionId, msgData)
	err = e.chatDataService.Write(ctx, msgData, sessionId)
	if err != nil {
		color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to write message data to repository, your chat data would not be saved! sessionId: %s, chatId: %s, provider: %s, error: %q", sessionId, msg.ChatID, instance.Name(), err))
	}

	provider, model, _ := e.sessionManager.GetModel(sessionId)
	msgList := e.sessionManager.GetMessages(sessionId)
	p, m := util.GetModelWithFallback(ctx, e.llmProviders, provider, model)
	respMessage := msg.Clone()

	// get total tokens in order to compact
	totToken, err := e.repo.GetTotalToken(ctx, sessionId)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] Error while getting total tokens from repository: %q", err))
	}
	// checkout whether a need to compact or not
	if uint64(totToken) > llm.CalculateCompactTokens(m) {
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

	var souls []string
	s, avail := e.sessionManager.GetSoul(sessionId)
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

	// jsonify, _ := json.MarshalIndent(msgList, "", "  ")
	// fmt.Printf(string(jsonify))

	maxCount := 3
	maxRepeatCount := 5
	count := 0
	repeatCount := 0
	repeatFingerPrint := ""

	for {
		event, err := e.LLMChat(ctx, p, m, msgList)
		errChat := e.HandleLLMChatError(ctx, instance, msg, p.DisplayName(), m.DisplayName, err)
		if errChat != nil {
			color.Yellow(fmt.Sprintf("[Manboster Engine] Error while sending message: %q\n", errChat))
		}
		if err != nil {
			break
		}
		if count >= maxCount || repeatCount >= maxRepeatCount {
			respMsg := msg.Clone()
			respMsg.MessageType = chat.MessageText
			respMsg.Text = &chat.TextPayload{
				Text: "Models tried the same tool call too much times or model calls return failed too much times, we broke it out in order to avoid consuming useless tokens and time.",
			}
			return e.SendMessage(ctx, instance, respMsg)
		}

		// write the whole event immediately
		e.sessionManager.AppendEvent(sessionId, *event)
		msgList = append(msgList, *event.Message)

		// add model cost
		util.CalculateCost(event, m)
		err = e.chatDataService.Write(ctx, *event, sessionId)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to write message data to repository, your chat data would not be saved! sessionId: %s, chatId: %s, provider: %s, error: %q", sessionId, respMessage.ChatID, instance.Name(), err))
		}

		// handle text message
		if event.EventType&llm.EventMessage != 0 && event.Message != nil && len(event.Message.Parts) > 0 {
			err = e.HandleText(ctx, instance, msg, *event)
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
				toolNameArgs = append(toolNameArgs, fmt.Sprintf("%s:%v", req.ToolName, req.ToolArgs))
			}
			if repeatFingerPrint == strings.Join(toolNameArgs, "%") {
				repeatCount++
			} else {
				repeatCount = 0
				repeatFingerPrint = strings.Join(toolNameArgs, "%")
			}

			respEvent, successExecution, err := e.HandleToolCall(ctx, instance, msg, *event)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Error while sending message: %q\n", err))
			}
			if !successExecution {
				count++
			} else {
				count = 0
			}

			e.sessionManager.AppendEvent(sessionId, respEvent)
			msgList = append(msgList, *respEvent.Message)

			// in order to avoid complexity?
			err = e.chatDataService.Write(ctx, respEvent, sessionId)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to write message data to repository, your chat data would not be saved! sessionId: %s, chatId: %s, provider: %s, error: %q", sessionId, respMessage.ChatID, instance.Name(), err))
			}
		}

	}

	return nil
}
