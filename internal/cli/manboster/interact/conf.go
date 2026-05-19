package interact

import (
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/database"
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
		return "Database\nThis will open your database and manage chats, users and sessions. If you want to manage chat sessions, purge unused sessions or more, please choose this."
	case entrypointConfig:
		return "Configuration\nThis will affect your model and chat provider's configuration and it displays the changes in config.yaml. If you want to manage models, default models, or modify application settings, please choose this."
	case entrypointEditor:
		return "Open Configuration yaml file in system's default editor\n(For advanced users only)"
	case entrypointQuit:
		return "Quit Manboster Configuration Wizard\nBye!"
	default:
		return ""
	}
}

func runConfigEntrypoint(p cli.Provider) error {
	selections := []entrypointType{entrypointConfig, entrypointDatabase, entrypointEditor, entrypointQuit}
	options := cli.BuildOptions[entrypointType](selections, nil)

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

		err = database.DBInstance.Init(cfg.App.DBPath)
		if err != nil {
			return err
		}

		repo := repository.New(database.DBInstance.Instance())
		return runDatabaseConfig(p, repo)
	})

	form.Register(entrypointQuit, func() error {
		return nil
	})

	var option cli.Option
	for {
		var err error
		option, err = p.Select("Please select what to configure!", "Welcome to Manboster Configuration Wizard! Please choose which field you want to configure.", options, option.Value, func(option cli.Option) error {
			return nil
		})
		if err != nil {
			return err
		}

		err = form.Handle(entrypointType(option.Value))
		if err != nil {
			return err
		}

		if option.Value == string(entrypointQuit) {
			color.Yellow("Bye!")
			return nil
		}
	}
}
