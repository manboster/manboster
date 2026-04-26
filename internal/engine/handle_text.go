package engine

import (
	"context"
	"fmt"
	"strings"

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

	maxCount := 5
	count := 0
	for {
		event, err := e.LLMChat(ctx, p, m, msgList)
		count++
		errChat := e.HandleLLMChatError(ctx, instance, msg, p.DisplayName(), m.DisplayName, err)
		if errChat != nil {
			color.Yellow(fmt.Sprintf("[Manboster Engine] Error while sending message: %q\n", errChat))
		}
		if err != nil || count >= maxCount {
			break
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
			err = e.SendMessage(ctx, instance, respMessage)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Error while sending message: %q\n", err))
			}
		}

		if event.EventType&llm.EventMessage != 0 && event.Message != nil && event.Message.Type&llm.MessageToolCallRequest == 0 {
			break
		}

		// handling tool call request
		if event.EventType&llm.EventMessage != 0 && event.Message != nil && event.Message.Type&llm.MessageToolCallRequest != 0 {
			var respEvent llm.Event
			respEvent.EventType = llm.EventMessage
			respEvent.Message = &llm.Message{
				Role: llm.RoleToolCall,
				Type: llm.MessageToolCallResponse,
			}
			for _, req := range event.Message.ToolCallRequest {
				resp := ""
				safeName := strings.ReplaceAll(req.ToolName, "_", ".")
				resp, err = e.HandleToolExec(ctx, safeName, req.ToolArgs.(string))
				if err != nil {
					resp = err.Error()
				}
				color.Blue(fmt.Sprintf("[Manboster Engine] Tool call %s responded with response: %q", safeName, resp))
				callMsg := msg.Clone()
				callMsg.MessageType = chat.MessageText
				toolProvider, av := e.toolMaps[safeName]
				callMsg.Reply = nil
				if !av {
					callMsg.Text = &chat.TextPayload{
						Text: fmt.Sprintf("Model called tool `%s` but not found.", safeName),
					}
				} else {
					callMsg.Text = &chat.TextPayload{
						Text: fmt.Sprintf("Model called tool `%s`(`%s`)!", toolProvider.DisplayName(), safeName),
					}
				}
				err = e.SendMessage(ctx, instance, callMsg)

				respEvent.Message.ToolCallResponse = append(respEvent.Message.ToolCallResponse, llm.MessageToolCallResponsePayload{
					ID:       req.ID,
					ToolName: req.ToolName,
					Result:   resp,
				})
			}

			e.sessionManager.AppendEvent(sessionId, respEvent)
			msgList = append(msgList, *respEvent.Message)

			// in order to avoid complexity?
			err := e.chatDataService.Write(ctx, *event, sessionId)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to write message data to repository, your chat data would not be saved! sessionId: %s, chatId: %s, provider: %s, error: %q", sessionId, respMessage.ChatID, instance.Name(), err))
			}
			err = e.chatDataService.Write(ctx, respEvent, sessionId)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to write message data to repository, your chat data would not be saved! sessionId: %s, chatId: %s, provider: %s, error: %q", sessionId, respMessage.ChatID, instance.Name(), err))
			}

		}

	}

	return nil
}
