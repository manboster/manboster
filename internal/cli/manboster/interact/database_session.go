package interact

import (
	"context"
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/spec/cli"
)

type databaseSessionPageAction string

const (
	databaseSessionPageEdit   databaseSessionPageAction = _EDIT_
	databaseSessionPageDelete databaseSessionPageAction = _DELETE_
	databaseSessionPageQuit   databaseSessionPageAction = _QUIT_
)

func (a databaseSessionPageAction) Name() string { return string(a) }
func (a databaseSessionPageAction) DisplayName() string {
	switch a {
	case databaseSessionPageEdit:
		return "Edit this session's provider and model"
	case databaseSessionPageDelete:
		return "Delete this session"
	case databaseSessionPageQuit:
		return "Quit"
	default:
		return ""
	}
}

func runDatabaseSessionConfig(p cli.Provider, repo repository.Repository) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sessions, err := repo.GetSessions(ctx)
	if err != nil {
		return err
	}
	chats, err := repo.GetAllChats(ctx)
	if err != nil {
		return err
	}
	chatsMap := make(map[string][]types.Chat)
	for _, c := range chats {
		chatsMap[c.SessionID] = append(chatsMap[c.SessionID], c)
	}
	purgeNum := len(sessions) - len(chatsMap)

	// purge first, then existing sessions, then quit
	options := []cli.Option{purgeOption}
	for _, sess := range sessions {
		var label strings.Builder
		label.WriteString(fmt.Sprintf("`%s`, used `%s:%s`, created at %s",
			sess.SessionID, sess.LLMProvider, sess.LLMProviderModel,
			sess.CreatedAt.Format("2006-01-02 15:04:05"),
		))
		if cm, ok := chatsMap[sess.SessionID]; ok {
			label.WriteString(fmt.Sprintf(", %d chats", len(cm)))
		}
		options = append(options, cli.Option{Key: label.String(), Value: sess.SessionID})
	}
	options = append(options, quitOption)

	summary := fmt.Sprintf("%d sessions loaded, %d sessions can be purged.", len(sessions), purgeNum)

	var option cli.Option
	for {
		option, err = p.Select("Select a session to manage.", summary, options, option.Value, func(o cli.Option) error {
			return nil
		})
		if err != nil {
			return err
		}

		if option.Value == _QUIT_ {
			return nil
		}

		if option.Value == _PURGE_ {
			if purgeNum <= 0 {
				if err := p.Alert("Manboster Configuration Wizard", "No need to purge sessions, your session list is clean and smart!"); err != nil {
					return err
				}
				continue
			}
			confirm, err := p.Prompt(fmt.Sprintf("Do you really want to DELETE %d unused sessions? YOUR ACTION IS IRREVERSIBLE!", purgeNum), "Do you want to continue?", "Yes", "No")
			if err != nil {
				return err
			}
			if !confirm {
				continue
			}
			var purgeErrors []string
			for _, sess := range sessions {
				if _, ok := chatsMap[sess.SessionID]; !ok {
					if err := repo.DeleteSession(ctx, sess.SessionID); err != nil {
						purgeErrors = append(purgeErrors, fmt.Sprintf("session %s: %q", sess.SessionID, err))
					}
					if err := repo.DeleteChatData(ctx, sess.SessionID); err != nil {
						purgeErrors = append(purgeErrors, fmt.Sprintf("chat data %s: %q", sess.SessionID, err))
					}
				}
			}
			if len(purgeErrors) > 0 {
				if err := p.Alert("Manboster Configuration Wizard", fmt.Sprintf("Some errors occurred during purge:\n%s", strings.Join(purgeErrors, "\n"))); err != nil {
					return err
				}
			}
			return p.Alert("Manboster Configuration Wizard", fmt.Sprintf("Successfully deleted %d unused sessions!", purgeNum))
		}

		// find selected session
		var selectedSession types.Session
		selectedIndex := -1
		for i, sess := range sessions {
			if sess.SessionID == option.Value {
				selectedSession = sess
				selectedIndex = i
				break
			}
		}
		if selectedIndex == -1 {
			return fmt.Errorf("unknown session selected: %s", option.Value)
		}

		var detail strings.Builder
		detail.WriteString(fmt.Sprintf("Session: `%s`\nProvider: `%s:%s`\nCreated: %s\nUpdated: %s\n",
			selectedSession.SessionID,
			selectedSession.LLMProvider, selectedSession.LLMProviderModel,
			selectedSession.CreatedAt.Format("2006-01-02 15:04:05"),
			selectedSession.UpdatedAt.Format("2006-01-02 15:04:05"),
		))
		if cm, ok := chatsMap[selectedSession.SessionID]; ok {
			detail.WriteString(fmt.Sprintf("Bind %d chats: ", len(cm)))
			for _, c := range cm {
				detail.WriteString(fmt.Sprintf("%s:%s ", c.ChatProvider, c.ChatID))
			}
			detail.WriteString("\n")
		}

		se := []databaseSessionPageAction{databaseSessionPageEdit, databaseSessionPageDelete, databaseSessionPageQuit}
		opts := cli.BuildOptions[databaseSessionPageAction](se, nil)
		form := newConfigForm[databaseSessionPageAction]()

		form.Register(databaseSessionPageEdit, func() error {
			cfg := config.Read()
			llmOptions := cli.BuildStringOptions(func() []string {
				var names []string
				for _, c := range cfg.LLMs {
					names = append(names, c.Provider)
				}
				return names
			}(), nil)
			providerOpt, err := p.Select("Select the LLM provider for this session.", "", llmOptions, selectedSession.LLMProvider, func(o cli.Option) error { return nil })
			if err != nil {
				return err
			}
			modelRaw, err := p.Input("Model name", "Enter the model name to use in this session.", selectedSession.LLMProviderModel, false, func(input string) error {
				if strings.TrimSpace(input) == "" {
					return fmt.Errorf("model name is required")
				}
				return nil
			})
			if err != nil {
				return err
			}
			if err := repo.UpdateSession(ctx, selectedSession.SessionID, map[string]interface{}{
				"llm_provider":       providerOpt.Value,
				"llm_provider_model": fmt.Sprintf("%v", modelRaw),
			}); err != nil {
				return err
			}
			return p.Alert("Manboster Configuration Wizard", "Successfully updated the session!")
		})

		form.Register(databaseSessionPageDelete, func() error {
			txt := ""
			if cm, ok := chatsMap[selectedSession.SessionID]; ok {
				txt = fmt.Sprintf("\nThis session is bound to %d chats. Deleting it will also delete the chat information.", len(cm))
			}
			confirm, err := p.Prompt(fmt.Sprintf("Do you want to delete session %s? YOUR ACTION IS IRREVERSIBLE!%s", selectedSession.SessionID, txt), "Do you want to continue?", "Yes", "No")
			if err != nil {
				return err
			}
			if !confirm {
				return errQuit
			}
			if err := repo.DeleteSession(ctx, selectedSession.SessionID); err != nil {
				return err
			}
			if err := repo.DeleteChatData(ctx, selectedSession.SessionID); err != nil {
				if alertErr := p.Alert("Manboster Configuration Wizard", fmt.Sprintf("Error deleting chat data: %q", err)); alertErr != nil {
					return alertErr
				}
			}
			if cm, ok := chatsMap[selectedSession.SessionID]; ok {
				for _, c := range cm {
					if err := repo.DeleteChat(ctx, c.ChatID, c.ChatProvider); err != nil {
						if alertErr := p.Alert("Manboster Configuration Wizard", fmt.Sprintf("Error deleting chat %s: %q", c.ChatID, err)); alertErr != nil {
							return alertErr
						}
					}
				}
			}
			return p.Alert("Manboster Configuration Wizard", "Successfully deleted session!")
		})

		form.Register(databaseSessionPageQuit, nilFunc)

		err = handleWithPrompt[databaseSessionPageAction](p, form, opts, detail.String(), "What do you want to do with this session?")
		if err != nil {
			return err
		}
	}
}
