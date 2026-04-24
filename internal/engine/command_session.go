package engine

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/llm"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/util"
)

// cmdCancel enables user to cancel their request
func (e *Engine) cmdCancel(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	msg.MessageType = chat.MessageText
	sessData, avail := e.sessionManager.GetSession(sessionId)

	var text string
	if avail {
		if sessData.Active {
			sessData.Cancel()
			text = "[Manboster] Successfully cancelled the request."
		} else {
			text = "[Manboster] The request in this session is not active."
		}
	} else {
		text = "[Manboster] The request in this session is not active."
	}

	msg.Text = &chat.TextPayload{
		Text: text,
	}

	return instance.SendMessage(ctx, msg)
}

// cmdNew deletes old session and creates a new session
func (e *Engine) cmdNew(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	_, avail := e.sessionManager.GetSession(sessionId)
	if !avail {
		respMessage.Text = &chat.TextPayload{
			Text: "Session is not active, there is nothing to do!",
		}
		return instance.SendMessage(ctx, respMessage)
	}

	e.sessionManager.DeleteSession(sessionId)
	err := e.repo.DeleteChat(ctx, msg.ChatID, instance.Name())
	if err != nil {
		return err
	}

	// delete chat data and session data
	err = e.repo.DeleteChatData(ctx, sessionId)
	if err != nil {
		return err
	}
	err = e.repo.DeleteSession(ctx, sessionId)
	if err != nil {
		return err
	}

	sid, err := e.loadSession(ctx, instance, msg, true)
	if err != nil {
		return err
	}

	// auto migration
	err = e.repo.ReplaceChatSessions(ctx, sessionId, sid)
	if err != nil {
		return err
	}

	respMessage.Text = &chat.TextPayload{
		Text: fmt.Sprintf("Deleted session `%s` and created session: `%s` with default provider and model.\nIf you want to change provider or model, please use `/provider` or `/model`.\nIf you want to save and create a new session, please use `/save` command.", sessionId, sid),
	}
	return instance.SendMessage(ctx, respMessage)
}

// cmdSave saves old session and create a new session
func (e *Engine) cmdSave(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	_, avail := e.sessionManager.GetSession(sessionId)
	if !avail {
		respMessage.Text = &chat.TextPayload{
			Text: "Session is not active, there is nothing to do!",
		}
		return instance.SendMessage(ctx, respMessage)
	}

	e.sessionManager.DeleteSession(sessionId)
	err := e.repo.DeleteChat(ctx, msg.ChatID, instance.Name())
	if err != nil {
		return err
	}

	sid, err := e.loadSession(ctx, instance, msg, true)
	if err != nil {
		return err
	}

	respMessage.Text = &chat.TextPayload{
		Text: fmt.Sprintf("Saved session `%s` and created session: `%s` with default provider and model.\nIf you want to change provider or model, please use `/provider` or `/model`.If you want to delete session and create a new session, please use `/new` command.", sessionId, sid),
	}
	return instance.SendMessage(ctx, respMessage)
}

func (e *Engine) cmdStatus(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	usage, err := e.repo.CountChatDataTokenBySession(ctx, sessionId)
	if err != nil {
		return err
	}

	totTokens, err := e.repo.GetTotalToken(ctx, sessionId)
	if err != nil {
		return err
	}

	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	sessData, _ := e.sessionManager.GetSession(sessionId)
	p, m := util.GetModelWithFallback(ctx, e.llmProviders, sessData.Provider, sessData.Model)
	provider := p
	model := m

	llmCallTimes := 0
	for _, data := range sessData.Events {
		if data.EventType&llm.EventMessage != 0 {
			if data.Message != nil && data.Message.Role == llm.RoleAssistant {
				llmCallTimes++
			}
		}
	}

	var respString strings.Builder
	respString.WriteString("Current status of this session:\n")
	respString.WriteString(fmt.Sprintf("Current usage data: \n"))
	respString.WriteString(fmt.Sprintf("Current session id: `%s`\n", sessionId))
	respString.WriteString(fmt.Sprintf("Current chat times: %d(call LLM API %d times)\n", len(sessData.Events), llmCallTimes))
	respString.WriteString(fmt.Sprintf("Current provider: `%s`, model: `%s`\n", provider.DisplayName(), model.DisplayName))
	respString.WriteString(fmt.Sprintf("Total Tokens Cost: %d tokens\n(input: %d tokens, output: %d tokens, thinking %d tokens)\n", usage.TotalTokens, usage.PromptTokens, usage.CompletionTokens, usage.TotalTokens-usage.PromptTokens-usage.CompletionTokens))

	totPrice := 0.0
	if model.InputPrice != 0 {
		inputPrice := model.InputPrice * float64(usage.PromptTokens) / 1000000
		totPrice += inputPrice
		respString.WriteString(fmt.Sprintf("Input Price: $%.6f/mtokens, Current Estimated Input Cost: $%.6f\n", model.InputPrice, inputPrice))
	}
	if model.OutputPrice != 0 {
		outputPrice := model.OutputPrice * float64(usage.TotalTokens-usage.PromptTokens) / 1000000
		totPrice += outputPrice
		respString.WriteString(fmt.Sprintf("Output Price: $%.6f/mtokens, Current Estimated Output Cost: $%.6f\n", model.OutputPrice, outputPrice))
	}
	if totPrice > 0 {
		respString.WriteString(fmt.Sprintf("Total Estimated Cost: $%.6f\n", totPrice))
	}

	respString.WriteString(fmt.Sprintf("Current Context length: %d / %d, now occupied %.2f%%", totTokens, model.Context, float64(totTokens*100)/float64(model.Context)))

	respMessage.Text = &chat.TextPayload{
		Text: respString.String(),
	}
	return instance.SendMessage(ctx, respMessage)
}

// cmdSession return and modify session if args is empty, it would display the list of sessions. if args is not empty, it would change session to given session id by modifying database
func (e *Engine) cmdSession(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	var respString strings.Builder

	// return session ids
	if len(msg.Command.CommandArgs) == 0 {
		sessionData, err := e.repo.GetSessions(ctx)
		if err != nil {
			color.Red(fmt.Sprintf("[Manboster Engine] we encountered an error when handling session data: %q", err))
			respString.WriteString("An error was occurred when handling session data!")
			respMessage.Text = &chat.TextPayload{
				Text: respString.String(),
			}
			return instance.SendMessage(ctx, respMessage)
		}
		respString.WriteString(fmt.Sprintf("Session List(for short, we only list 20 latest sessions, if you want to get current session id, please run `/status`.):\n"))
		for _, data := range sessionData {
			respString.WriteString(fmt.Sprintf("Session ID: `%s`(Create Time: `%s`, Provider: `%s`, Model: `%s`)\n", data.SessionID, data.CreatedAt, data.LLMProvider, data.LLMProviderModel))
		}
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	// checkout session
	sid := msg.Command.CommandArgs[0]
	_, err := e.repo.GetSession(ctx, sid)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			color.Yellow(fmt.Sprintf("[Manboster Engine] we could not found any session id"))
			respString.WriteString("The session id you entered does not exist!")
		} else {
			color.Red(fmt.Sprintf("[Manboster Engine] we encountered an error when getting session: %s", err))
			respString.WriteString("An error was occurred when getting session id you entered!")
		}
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	err = e.repo.UpdateChat(ctx, msg.ChatID, instance.Name(), sid)
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Engine] we encountered an error when handling updating chat's session data: %q", err))
		respString.WriteString("An error was occurred when changing session id for this chat!")
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	respString.WriteString(fmt.Sprintf("Successfully changed session to `%s`!", sid))
	respMessage.Text = &chat.TextPayload{
		Text: respString.String(),
	}
	return instance.SendMessage(ctx, respMessage)
}

// cmdProvider returns current available providers, if args is empty, it would display the list of providers. if args is not empty, it would change providers to given provider id by modifying database
func (e *Engine) cmdProvider(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	var respString strings.Builder
	if len(msg.Command.CommandArgs) == 0 {
		respString.WriteString("Available Providers(If you want to see current provider, please run `/status`, if you want to change provider, please run `/provider [name]`, we will automatically change the first model of the provider for you):\n")
		i := 0
		for _, provider := range e.llmProviders {
			respString.WriteString(fmt.Sprintf("ID:`%d`) `%s`, %d available Models. Run `/provider %s` to change.\n", i+1, provider.DisplayName(), len(provider.Models()), provider.Name()))
			i += 1
		}
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	id := msg.Command.CommandArgs[0]
	if id == "" {
		respString.WriteString("Invalid input data!\n")
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	s, _ := e.sessionManager.GetSession(sessionId)

	if _, avail := e.llmProviders[id]; !avail {
		respString.WriteString("Current provider is what you have entered!")
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	s.Provider = e.llmProviders[id].Name()
	s.Model = e.llmProviders[id].Models()[0].Name

	e.sessionManager.SetSession(sessionId, s)
	err := e.repo.UpdateSession(ctx, sessionId, map[string]interface{}{
		"llm_provider":       s.Provider,
		"llm_provider_model": s.Model,
	})
	if err != nil {
		respString.WriteString("An error was occurred when updating provider name for this session!")
		color.Red(fmt.Sprintf("[Manboster Engine] An error was occurred when updating provider name for this session: %q", err))
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	respString.WriteString(fmt.Sprintf("Successfully changed this session's provider to `%s`, model `%s`.", s.Provider, s.Model))
	respMessage.Text = &chat.TextPayload{
		Text: respString.String(),
	}
	return instance.SendMessage(ctx, respMessage)
}

// cmdModel if args is empty, it would display the list of models. if args is not empty, is would change models to given provider id by modifying database
func (e *Engine) cmdModel(ctx context.Context, instance chat.Provider, msg *chat.Message, sid string) error {
	respMessage := msg.Clone()
	respMessage.MessageType = chat.MessageText
	var respString strings.Builder

	s, _ := e.sessionManager.GetSession(sid)

	p, _ := util.GetModelWithFallback(ctx, e.llmProviders, s.Provider, s.Model)
	// fmt.Printf("%s %s %s %s", s.Model, s.Provider, p.DisplayName(), sid)
	if len(msg.Command.CommandArgs) == 0 {
		respString.Reset()
		respString.WriteString("Available Models(If you want to see current model, please run `/status`, if you want to change model, please run `/model [id]`.):\n")
		for i, m := range p.Models() {
			respString.WriteString(fmt.Sprintf("ID:`%d`) `%s`, context: `%d`, max output tokens: `%d` input: `$%.4f`/mtokens, output: `$%.4f`/mtokens. Run `/model %s` to change.\n", i+1, m.DisplayName, m.Context, m.MaxOutputTokens, m.InputPrice, m.OutputPrice, m.Name))
		}
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	id := msg.Command.CommandArgs[0]
	if id == "" {
		respString.WriteString("Invalid input data!\n")
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		color.Red(fmt.Sprintf("[Manboster Engine] An error was occurred when parsing data"))
		return instance.SendMessage(ctx, respMessage)
	}

	flag := false
	for _, m := range p.Models() {
		if m.Name == s.Model {
			flag = true
			break
		}
	}
	if !flag {
		respString.WriteString("Could not find model named `" + s.Model + "`")
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		return instance.SendMessage(ctx, respMessage)
	}

	s.Model = id
	e.sessionManager.SetSession(sid, s)
	err := e.repo.UpdateSession(ctx, sid, map[string]interface{}{
		"llm_provider_model": s.Model,
	})
	if err != nil {
		respString.WriteString("An error was occurred when updating model name for this session!")
		respMessage.Text = &chat.TextPayload{
			Text: respString.String(),
		}
		color.Red(fmt.Sprintf("[Manboster Engine] An error was occurred when updating model name for this session: %q", err))
		return instance.SendMessage(ctx, respMessage)
	}

	respString.WriteString(fmt.Sprintf("Successfully changed this session's model to `%s`.", s.Model))
	respMessage.Text = &chat.TextPayload{
		Text: respString.String(),
	}
	return instance.SendMessage(ctx, respMessage)
}
