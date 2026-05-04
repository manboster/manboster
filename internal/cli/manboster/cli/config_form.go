package cli

import (
	"os"

	"github.com/charmbracelet/huh"
	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/config"
	"github.com/manboster/manboster/internal/database"
	"github.com/manboster/manboster/internal/repository"
)

func configFormRun() error {
	for {
		se, err := configStartupForm()
		if err != nil {
			return err
		}
		err = config.Init()
		if err != nil {
			color.Red("It seems that there is no configuration available in your device, please run 'manboster onboard' first!")
			return err
		}
		switch se {
		case SelectionEditor:
			configCmdOpenRun(nil, nil)
			color.Blue("Opened via your default Editor, please edit in the editor, save it and then restart the instance!")
			os.Exit(0)
		case SelectionConfig:

		case SelectionDatabase:
			cli := database.Client{}
			path := config.Read().App.DBPath
			err := cli.Init(path)
			if err != nil {
				color.Red("It seems that there is no database available in your device, please run 'manboster' first!")
				return err
			}

			repo := repository.New(cli.Instance())
			s := newDatabaseConfigService(repo)
			err = s.configDatabaseLandingForm()
			if err != nil {
				return err
			}
		case SelectionQuit:
			color.Blue("Bye!")
			os.Exit(0)
		}
	}

	return nil
}

func configStartupForm() (Selection, error) {
	var s Selection
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[Selection]().Options(
				huh.NewOption("Database\nThis will open your database and manage chats, users and sessions. If you want to manage chat sessions, purge unused sessions or more, please choose this.", SelectionDatabase),
				huh.NewOption("Configuration\nThis will affect your model and chat provider's configuration and it displays the changes in config.yaml. If you want to manage models, default models, or modify application settings, please choose this.", SelectionConfig),
				huh.NewOption("Open Configuration yaml file in system's default editor(For advanced users only)", SelectionEditor),
				huh.NewOption("Quit Manboster Configuration Wizard", SelectionQuit),
			).Title("Please select what to configure").Description("Welcome to Manboster Configuration Wizard! Please choose which field you want to configure.").Value(&s),
		)).Run()
	if err != nil {
		return "", err
	}
	return s, nil
}

func configChatConfigRun() error {
	return nil
}
