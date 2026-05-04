package cli

import "github.com/charmbracelet/huh"

func configFormRun() error {
	_, err := configStartupForm()
	if err != nil {
		return err
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
