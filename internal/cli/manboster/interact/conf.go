package interact

import (
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/database"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/spec/cli"
)

type entrypointType string

const (
	entrypointConfig   entrypointType = "config"
	entrypointDatabase entrypointType = "database"
	entrypointEditor   entrypointType = "editor"
	entrypointQuit     entrypointType = "quit"
)

func (t entrypointType) Name() string {
	return string(t)
}

func (t entrypointType) DisplayName() string {
	switch t {
	case entrypointDatabase:
		return i18n.T(keys.CliConfigEntrypointDatabase)
	case entrypointConfig:
		return i18n.T(keys.CliConfigEntrypointConfig)
	case entrypointEditor:
		return i18n.T(keys.CliConfigEntrypointEditor)
	case entrypointQuit:
		return i18n.T(keys.CliConfigEntrypointQuit)
	default:
		return ""
	}
}

func runConfigEntrypoint(p cli.Provider) error {
	selections := []entrypointType{entrypointConfig, entrypointDatabase, entrypointEditor, entrypointQuit}
	options := cli.BuildOptions[entrypointType](selections, nil)
	mark := false

	form := newConfigForm[entrypointType]()
	form.Register(entrypointEditor, func() error {
		return openEditor(config.Path("config.yaml"))
	})

	form.Register(entrypointConfig, func() error {
		err := config.Init()
		if err != nil {
			return err
		}

		cfg := config.Read()
		cfg, err = runConfig(p, cfg)
		if err != nil {
			return err
		}

		err = cfg.Validate()
		if err != nil {
			return err
		}

		return config.Write(cfg, config.Path("config.yaml"))
	})

	form.Register(entrypointDatabase, func() error {
		err := config.Init()
		if err != nil {
			return err
		}
		cfg := config.Read()

		database.DBInstance = &database.Client{}
		err = database.DBInstance.Init(cfg.App.DBPath)
		if err != nil {
			return err
		}

		repo := repository.New(database.DBInstance.Instance())
		return runDatabaseConfig(p, repo)
	})

	form.Register(entrypointQuit, func() error {
		mark = true
		return nil
	})

	for {
		err := handle[entrypointType](p, form, options, i18n.T(keys.CliConfigEntrypointSelectPrompt), i18n.T(keys.CliConfigEntrypointSelectHelp))
		if err != nil {
			return err
		}
		if mark {
			return nil
		}
	}
}
