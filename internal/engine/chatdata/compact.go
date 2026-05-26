package chatdata

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config/prompt"
	"github.com/manboster/manboster/internal/engine/hook"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

// Compact compacts chat data and then open a new session.
func (s *Service) Compact(ctx context.Context, instance chat.Provider, mesg *chat.Message, sessionId string) error {
	// overwrite context because if not specified there will be a problem that context canceled with runner.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	msg := s.sessionManager.GetMessages(sessionId)
	provider, model, _ := s.sessionManager.GetModel(sessionId)
	p, m := util.GetModelWithFallback(ctx, s.llmProviders, provider, model)

	var i int
	if len(msg) <= 10 {
		return ErrNoNeedToCompact
	}
	for i = len(msg) - 1; msg[i].Type&(llm.MessageToolCallResponse|llm.MessageToolCallRequest) != 0; i-- {
	}

	hookProviders := hook.Reg.GetProviders(hook.EngineBeforeCompact)
	for _, hookP := range hookProviders {
		ho, ok := hookP.(hook.EngineBeforeCompactHookProvider)
		if !ok {
			color.Yellow(fmt.Sprintf("[Manboster ChatData] Failed to assert in before compact hook provider, please check function is valid or not."))
			continue
		}
		err := ho.PolyfillFunc(ctx, sessionId)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster ChatData] Failed to polyfill session in before compact hook: %v", err))
		}
	}

	messagesToCompact := msg // compact data
	var compactString strings.Builder
	for _, message := range messagesToCompact {
		compactString.WriteString(util.ConvertLLMMessageToString(message) + "\n")
	}
	event, err := s.gateway.LLMQuickChat(ctx, p, m, prompt.CompactSystemPrompt, compactString.String())
	if err != nil {
		return err
	}

	if event == nil {
		color.Red("[Manboster ChatData] event is nil")
		return ErrCompactChatFailed
	}
	if event.EventType&llm.EventMessage == 0 || event.Message == nil {
		color.Red("event message is nil")
		return ErrCompactChatFailed
	}
	if event.Message.Type&llm.MessageText == 0 || len(event.Message.Parts) == 0 {
		color.Red("event message parts is nil")
		return ErrCompactChatFailed
	}
	if event.Message.Parts[0].PartsType != llm.MessagePartsText || event.Message.Parts[0].Text == nil {
		return ErrCompactChatFailed
	}

	compactedMessage := event.Message.Parts[0].Text.Text
	newSessionID := util.RandomString(8)
	err = s.repo.CreateSoul(ctx, types.Soul{
		Priority: 1,
		Name:     "previous-message-" + newSessionID,
		Content:  "# Previous chat messages:\n" + compactedMessage,
		Scope: []string{
			"session:" + newSessionID,
		},
	})
	if err != nil {
		return err
	}
	s.sessionManager.DeleteSession(sessionId)
	err = s.repo.CreateSession(ctx, types.Session{
		SessionID:        newSessionID,
		LLMProviderModel: model,
		LLMProvider:      provider,
		ActivatedSouls: []string{
			"system", "previous-message-" + newSessionID,
		},
	})
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster ChatData] Failed to create session in repository when compacting: %v", err))
		return err
	}

	hookProviders = hook.Reg.GetProviders(hook.EngineAfterCompact)
	for _, hookP := range hookProviders {
		ho, ok := hookP.(hook.EngineAfterCompactHookProvider)
		if !ok {
			color.Yellow(fmt.Sprintf("[Manboster ChatData] Failed to assert in hook provider, please check function is valid or not."))
			continue
		}
		err := ho.PolyfillFunc(ctx, sessionId, newSessionID)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster ChatData] Failed to polyfill session in hook: %v", err))
		}
	}

	err = s.repo.ReplaceChatSessions(ctx, sessionId, newSessionID)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster ChatData] Failed to replace session in repository when compacting: %v", err))
		return err
	}

	respMessage := mesg.Clone()
	respMessage.MessageType = chat.MessageText
	respMessage.Text = &chat.TextPayload{
		Text: fmt.Sprintf(i18n.T(keys.EngineChatDataCompactSuccess), sessionId, newSessionID, compactedMessage, newSessionID),
	}
	return instance.SendMessage(ctx, respMessage)
}
