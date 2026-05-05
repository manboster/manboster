package interactive

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/helper"
)

type databaseConfigSessionPageSelection string

const (
	databaseConfigSessionPageQuit   databaseConfigSessionPageSelection = "quit"
	databaseConfigSessionPageEdit   databaseConfigSessionPageSelection = "edit"
	databaseConfigSessionPageDelete databaseConfigSessionPageSelection = "delete"
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
		color.Blue("Bye!")
		return nil
	case databaseConfigSessionPageDelete:
		color.Blue("Bye!")
		return nil
	case databaseConfigSessionPageQuit:
		return nil
	default:
		return fmt.Errorf("unexpected database session page: %s", se)
	}
}
