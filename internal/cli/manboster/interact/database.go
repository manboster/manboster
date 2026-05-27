package interact

import (
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
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
		return i18n.T(keys.CliConfigDatabaseSessions)
	case databaseConfigLandingUser:
		return i18n.T(keys.CliConfigDatabaseUsers)
	case databaseConfigLandingSoul:
		return i18n.T(keys.CliConfigDatabaseSouls)
	case databaseConfigLandingQuit:
		return i18n.T(keys.CliConfigDatabaseQuit)
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

	return handle[databaseConfigLandingSelection](p, form, options, i18n.T(keys.CliConfigDatabaseSelectPrompt), i18n.T(keys.CliConfigDatabaseSelectHelp))
}
