package interact

import (
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/spec/cli"
)

type databaseConfigLandingSelection string

const (
	databaseConfigLandingUser    databaseConfigLandingSelection = "user"
	databaseConfigLandingSession databaseConfigLandingSelection = "session"
	databaseConfigLandingSoul    databaseConfigLandingSelection = "soul"
	databaseConfigLandingQuit    databaseConfigLandingSelection = "quit"
)

func (s databaseConfigLandingSelection) Name() string {
	return string(s)
}

func (s databaseConfigLandingSelection) DisplayName() string {
	switch s {
	case databaseConfigLandingSession:
		return "Sessions\nIt's chat sessions you made in daily chatting routine, you can purge it to save disk or do more."
	case databaseConfigLandingUser:
		return "Users\nIt helps you to grant a user to an admin or degrade it. It's all up to you to decide."
	case databaseConfigLandingSoul:
		return "Souls\nSouls is the key system prompt for you to personalize your Manbo Lobster."
	case databaseConfigLandingQuit:
		return "Quit\nBye!"
	default:
		return ""
	}
}

func runDatabaseConfig(p cli.Provider, repo repository.Repository) error {
	se := []databaseConfigLandingSelection{databaseConfigLandingUser, databaseConfigLandingSession, databaseConfigLandingSoul, databaseConfigLandingQuit}
	options := cli.BuildOptions[databaseConfigLandingSelection](se, nil)

	form := newConfigForm[databaseConfigLandingSelection]()

	form.Register(databaseConfigLandingUser, func() error {
		return runDatabaseUserConfig(p, repo)
	})

	form.Register(databaseConfigLandingSession, func() error {
		return runDatabaseSessionConfig(p, repo)
	})

	form.Register(databaseConfigLandingSoul, func() error {
		return runDatabaseSoulConfig(p, repo)
	})

	form.Register(databaseConfigLandingQuit, nilFunc)

	return handle[databaseConfigLandingSelection](p, form, options, "Please select what to configure in database", "Please choose what you want to configure in database field.")
}
