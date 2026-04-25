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
	llmProviderDisplayName := p.DisplayName()
	llmModelDisplayName := m.DisplayName
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

	event, err := e.LLMChat(ctx, p, m, msgList)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] Failed to get message from LLMProvider %q, model %q, get error: %q", llmProviderDisplayName, llmModelDisplayName, err))
		// now we have to wrap this into friendly prompt
		tips := fmt.Sprintf("%+v", err)
		text := fmt.Sprintf("[Manboster] Failed to get message from LLMProvider %q, model %q, get error: %s\nYou can resend your message or check the API's availability.", llmProviderDisplayName, llmModelDisplayName, err)
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
		return e.SendMessage(ctx, instance, respMessage)
	}

	if event.EventType&llm.EventMessage != 0 && event.Message != nil && event.Message.Type&llm.MessageToolCallRequest != 0 {
		e.sessionManager.AppendEvent(sessionId, *event)
		msgList = append(msgList, *event.Message)
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
			respEvent.Message.ToolCallResponse = append(respEvent.Message.ToolCallResponse, llm.MessageToolCallResponsePayload{
				ID:       req.ID,
				ToolName: req.ToolName,
				Result:   resp,
			})
		}

		msgList = append(msgList, *respEvent.Message)

		//jsonify, _ := json.MarshalIndent(msgList, "", " ")
		//fmt.Println(string(jsonify))

		// TODO: recursive use LLMChat to call handleTool
		event, err = e.LLMChat(ctx, p, m, msgList)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] Failed to get message from LLMProvider %q, model %q, get error: %q", llmProviderDisplayName, llmModelDisplayName, err))
			// now we have to wrap this into friendly prompt
			tips := fmt.Sprintf("%+v", err)
			text := fmt.Sprintf("[Manboster] Failed to get message from LLMProvider %q, model %q, get error: %s\nYou can resend your message or check the API's availability.", llmProviderDisplayName, llmModelDisplayName, err)
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
			return e.SendMessage(ctx, instance, respMessage)
		}

		//jsonify, _ = json.MarshalIndent(event, "", " ")
		//fmt.Println(string(jsonify))
	}

	if event.EventType&llm.EventMessage != 0 && event.Message != nil && len(event.Message.Parts) > 0 {
		// add model cost
		util.CalculateCost(event, m)

		text := event.Message.Parts[0].Text.Text
		// fmt.Println(text)
		textWithoutThinking := util.StripThink(text)
		// fmt.Println(textWithoutThinking)
		e.sessionManager.AppendEvent(sessionId, *event)
		err := e.chatDataService.Write(ctx, *event, sessionId)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Engine] Failed to write message data to repository, your chat data would not be saved! sessionId: %s, chatId: %s, provider: %s, error: %q", sessionId, respMessage.ChatID, instance.Name(), err))
		}

		// If there is a thinking context
		if util.ExtractThinkContent(text) != "" {
			respMessage.MessageType = chat.MessageThinkingText
			respMessage.Text = &chat.TextPayload{
				Text: util.ExtractThinkContent(text),
			}
			err := e.SendMessage(ctx, instance, respMessage)
			if err != nil {
				return err
			}
		}

		respMessage.MessageType = chat.MessageText
		respMessage.Text = &chat.TextPayload{
			Text: textWithoutThinking,
		}
	}

	if respMessage.Text == nil {
		return nil
	}

	return e.SendMessage(ctx, instance, respMessage)
}
