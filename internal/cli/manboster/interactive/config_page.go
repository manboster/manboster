package interactive

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/config"
)

func (s *databaseConfigService) runConfigDatabaseSessionSelect(ctx context.Context) error {
	var options []huh.Option[string]

	var selection string

	defer helper.ClearScreen()

	for _, sess := range s.sessions {
		var outputMsg strings.Builder
		outputMsg.WriteString(fmt.Sprintf("%d) `%s`, used `%s:%s` created at %s, updated at %s.\n", sess.ID, sess.SessionID, sess.LLMProvider, sess.LLMProviderModel, sess.CreatedAt.Format("2006-01-02 15:04:05"), sess.UpdatedAt.Format("2006-01-02 15:04:05")))
		cm, avail := s.chatsMap[sess.SessionID]
		if avail {
			outputMsg.WriteString(fmt.Sprintf("Bind %d chats: ", len(cm)))
			for _, c := range cm {
				outputMsg.WriteString(fmt.Sprintf("%s:%s ", c.ChatProvider, c.ChatID))
			}
		}
		options = append(options, huh.NewOption(outputMsg.String(), sess.SessionID))
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Options(
				options...,
			).Title("Please choose session you want to edit").Description("Please choose session you want to edit.").Value(&selection),
		)).Run()
	if err != nil {
		return err
	}

	for _, sess := range s.sessions {
		if sess.SessionID == selection {
			var outputMsg strings.Builder
			outputMsg.WriteString("Selected session:\n")
			outputMsg.WriteString(fmt.Sprintf("%d) `%s`, used `%s:%s` created at %s, updated at %s.\n", sess.ID, sess.SessionID, sess.LLMProvider, sess.LLMProviderModel, sess.CreatedAt.Format("2006-01-02 15:04:05"), sess.UpdatedAt.Format("2006-01-02 15:04:05")))
			cm, avail := s.chatsMap[sess.SessionID]
			if avail {
				outputMsg.WriteString(fmt.Sprintf("Bind %d chats: ", len(cm)))
				for _, c := range cm {
					outputMsg.WriteString(fmt.Sprintf("%s:%s ", c.ChatProvider, c.ChatID))
				}
			}
			helper.DisplayText(outputMsg.String())
			break
		}
	}

	var se databaseConfigSessionPageSelection
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[databaseConfigSessionPageSelection]().Options(
				huh.NewOption("Edit this session's provider and model", databaseConfigSessionPageEdit),
				huh.NewOption("Delete this session", databaseConfigSessionPageDelete),
				huh.NewOption("Quit", databaseConfigSessionPageQuit),
			).Description("What do you want to do?").Value(&se),
		)).Run()

	switch se {
	case databaseConfigSessionPageEdit:
		llmConfigs := config.Read().LLMs
		provider, err := LLMProviderInstanceForm(ctx, llmConfigs, "Select the default provider you want to use in this session:", "")
		if err != nil {
			return err
		}
		model, err := SelectModelForm(ctx, provider.Models(), "Select the default model you want to use in this session:", "")
		if err != nil {
			return err
		}
		err = s.editConfigSessionDatabase(ctx, selection, provider.Name(), model.Name)
		if err != nil {
			return err
		}
		color.Blue("Successfully updated the default model and provider!")
		time.Sleep(1 * time.Second)
		return nil
	case databaseConfigSessionPageDelete:
		txt := ""
		cm, avail := s.chatsMap[selection]
		if avail {
			txt = fmt.Sprintf("\nThis session bind with %d chats, if you delete it, the chats information will be deleted too.", len(cm))
		}
		if helper.ContinueConfirm(ctx, fmt.Sprintf("Do you want to continue delete session %s? YOUR ACTION IS IRREVERSIBLE!%s", selection, txt)) {
			err := s.deleteConfigSessionDatabase(ctx, selection)
			if err != nil {
				return err
			}
			color.Blue("Successfully deleted session!")
			time.Sleep(1 * time.Second)
			return nil
		}
		return nil
	case databaseConfigSessionPageQuit:
		return nil
	default:
		return fmt.Errorf("unexpected database session page: %s", se)
	}
}
