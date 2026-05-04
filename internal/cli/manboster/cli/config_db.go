package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
)

type databaseConfigService struct {
	repo     repository.Repository
	sessions []types.Session
	chats    []types.Chat
}

type databaseConfigLandingSelection string

const (
	databaseConfigLandingUser    databaseConfigLandingSelection = "user"
	databaseConfigLandingSession databaseConfigLandingSelection = "session"
	databaseConfigLandingSoul    databaseConfigLandingSelection = "soul"
	databaseConfigLandingQuit    databaseConfigLandingSelection = "quit"
)

func newDatabaseConfigService(repo repository.Repository) *databaseConfigService {
	return &databaseConfigService{repo: repo}
}

func (s *databaseConfigService) configDatabaseLandingForm() error {
	for {
		se, err := s.runConfigDatabaseLandingSelection()
		if err != nil {
			return err
		}
		switch se {
		case databaseConfigLandingUser:
			return nil
		case databaseConfigLandingSession:
			err := s.runConfigDatabaseSessionSelection()
			if err != nil {
				return err
			}
			return nil
		case databaseConfigLandingSoul:
			return nil
		case databaseConfigLandingQuit:
			color.Blue("Bye!")
			return nil
		default:
			return fmt.Errorf("unexpected database landing form: %s", se)
		}
	}
}

func (s *databaseConfigService) runConfigDatabaseLandingSelection() (databaseConfigLandingSelection, error) {
	var se databaseConfigLandingSelection
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[databaseConfigLandingSelection]().Options(
				huh.NewOption("Sessions\nIt's chat sessions you made in daily chatting routine, you can purge it to save disk or do more.", databaseConfigLandingSession),
				huh.NewOption("Users\nIt helps you to grant a user to an admin or degrade it. It's all up to you to decide.", databaseConfigLandingUser),
				huh.NewOption("Souls\nSouls is the key system prompt for you to personalize your Manbo Lobster.", databaseConfigLandingSoul),
				huh.NewOption("Quit Manboster Configuration Wizard", databaseConfigLandingQuit),
			).Title("Please select what to configure in database").Description("Please choose which field you want to configure in database field.").Value(&se),
		)).Run()
	if err != nil {
		return "", err
	}
	return se, nil
}

type databaseConfigSessionSelection string

const (
	databaseConfigSessionPurge databaseConfigSessionSelection = "purge"
	databaseConfigSessionList  databaseConfigSessionSelection = "list"
	databaseConfigSessionQuit  databaseConfigSessionSelection = "quit"
)

func (s *databaseConfigService) runConfigDatabaseSessionSelection() error {
	var se databaseConfigSessionSelection
	err := s.printConfigDatabaseSessionList()
	if err != nil {
		return err
	}
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[databaseConfigSessionSelection]().Options(
				huh.NewOption("Purge Session Data", databaseConfigSessionPurge),
				huh.NewOption("Select Sessions", databaseConfigSessionList),
				huh.NewOption("Quit", databaseConfigSessionQuit),
			).Value(&se),
		)).Run()
	if err != nil {
		return err
	}
	helper.ClearScreen()
	switch se {
	case databaseConfigSessionPurge:
		return nil
	case databaseConfigSessionList:
		return nil
	case databaseConfigSessionQuit:
		color.Blue("Bye!")
		return nil
	default:
		return fmt.Errorf("unexpected database session selection: %s", se)
	}
}

func (s *databaseConfigService) printConfigDatabaseSessionList() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// string is sessionID
	chatsMap := make(map[string][]types.Chat)

	sessions, err := s.repo.GetSessions(ctx)
	if err != nil {
		return err
	}
	s.sessions = sessions

	chats, err := s.repo.GetAllChats(ctx)
	if err != nil {
		return err
	}
	s.chats = chats

	for _, c := range chats {
		cm, avail := chatsMap[c.SessionID]
		if !avail {
			cm = []types.Chat{c}
			chatsMap[c.SessionID] = cm
			continue
		}
		cm = append(cm, c)
		chatsMap[c.SessionID] = cm
	}

	var outputMsg strings.Builder
	for _, sess := range s.sessions {
		outputMsg.WriteString(fmt.Sprintf("%d) `%s`, used `%s:%s` created at %s, updated at %s.\n", sess.ID, sess.SessionID, sess.LLMProvider, sess.LLMProviderModel, sess.CreatedAt.Format("2006-01-02 15:04:05"), sess.UpdatedAt.Format("2006-01-02 15:04:05")))
		cm, avail := chatsMap[sess.SessionID]
		if avail {
			outputMsg.WriteString(fmt.Sprintf("Bind %d chats: ", len(cm)))
			for _, c := range cm {
				outputMsg.WriteString(fmt.Sprintf("%s:%s ", c.ChatProvider, c.ChatID))
			}
			outputMsg.WriteString("\n")
		}
	}

	outputMsg.WriteString(fmt.Sprintf("%d sessions loaded, %d sessions can be purged.", len(s.sessions), len(sessions)-len(chatsMap)))

	helper.DisplayText(outputMsg.String())
	return nil
}
