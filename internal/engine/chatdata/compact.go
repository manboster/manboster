package chatdata

import (
	"context"
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/util"
)

// Compact compacts chat data and then open a new session.
func (s *Service) Compact(ctx context.Context, instance chat.Provider, mesg *chat.Message, sessionId string) error {
	msg := s.sessionManager.GetMessages(sessionId)
	provider, model, _ := s.sessionManager.GetModel(sessionId)
	pIndex, _ := util.GetModelIndexWithFallback(ctx, s.llmProviders, provider, model)

	count := 0
	splitIndex := 0
	isChecked := false
	var i int
	for i = len(msg) - 1; i >= 0; i-- {
		if msg[i].Role == llm.RoleAssistant && !isChecked {
			isChecked = true
		}
		if msg[i].Role == llm.RoleUser && isChecked {
			count++
			isChecked = false
		}
		if count >= 5 {
			splitIndex = i
			break
		}
	}
	if count < 5 || splitIndex == 0 {
		return ErrNoNeedToCompact
	}

	messagesToCompact := msg[:splitIndex] // compact data
	recentMessages := msg[splitIndex:]    // messages reserved

	var compactString strings.Builder
	for _, message := range messagesToCompact {
		compactString.WriteString(util.ConvertLLMMessageToString(message) + "\n")
	}

	var appendMessages []llm.Message
	message := llm.Message{
		Role: llm.RoleSystem,
		Type: llm.MessageText,
		Parts: []llm.MessageParts{
			{
				PartsType: llm.MessagePartsText,
				Text: &llm.MessageTextPayload{
					Text: config.CompactSystemPrompt,
				},
			},
		}}
	uMessage := llm.Message{
		Role: llm.RoleUser,
		Type: llm.MessageText,
		Parts: []llm.MessageParts{
			{
				PartsType: llm.MessagePartsText,
				Text: &llm.MessageTextPayload{
					Text: compactString.String(),
				},
			},
		},
	}
	appendMessages = append(appendMessages, message)
	appendMessages = append(appendMessages, uMessage)

	event, err := s.llmProviders[pIndex].Chat(ctx, model, appendMessages)
	if err != nil {
		return err
	}

	if event == nil {
		fmt.Println("event is nil")
		return ErrCompactChatFailed
	}
	if event.EventType&llm.EventMessage == 0 || event.Message == nil {
		fmt.Println("event message is nil")
		return ErrCompactChatFailed
	}
	if event.Message.Type&llm.MessageText == 0 || len(event.Message.Parts) == 0 {
		fmt.Println("event message parts is nil")
		return ErrCompactChatFailed
	}
	if event.Message.Parts[0].PartsType != llm.MessagePartsText || event.Message.Parts[0].Text == nil {
		return ErrCompactChatFailed
	}

	compactedMessage := event.Message.Parts[0].Text.Text
	newSessionID := util.RandomString(32)
	err = s.CreateWithSystemPrompt(ctx, newSessionID, "<previous_chat>"+compactedMessage+"</previous_chat>")
	if err != nil {
		return err
	}
	for _, m := range recentMessages {
		err := s.Write(ctx, llm.Event{
			EventType: llm.EventMessage,
			Message:   &m,
		}, newSessionID)
		if err != nil {
			return err
		}
	}

	respMessage := mesg.Clone()
	respMessage.MessageType = chat.MessageText
	respMessage.Text = &chat.TextPayload{
		Text: fmt.Sprintf("Compacted session `%s` and created a new session `%s`.\nCompacted summary data:\n```%s```\nOld session is saved. If you want to bind other chats to this new session, please run `/session %s` in that chat.", sessionId, newSessionID, compactedMessage, newSessionID),
	}
	return instance.SendMessage(ctx, respMessage)
}
