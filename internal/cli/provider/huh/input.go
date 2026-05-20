package huh

import (
	"github.com/charmbracelet/huh"
)

func (h Huh) Input(title string, description string, defaultVal string, secret bool, validate func(input string) error) (any, error) {
	defer ClearScreen()
	var data string
	if defaultVal != "" {
		data = defaultVal
	}

	i := huh.NewInput().Title(title).Description(description).Value(&data).Validate(validate)
	if secret {
		i.EchoMode(huh.EchoModePassword)
	}

	err := huh.NewForm(
		huh.NewGroup(
			i,
		)).Run()
	if err != nil {
		return nil, err
	}
	return data, nil
}
