package huh

import (
	"github.com/charmbracelet/huh"
)

func (h Huh) Input(title string, description string, validate func(input string) error) (any, error) {
	var data string
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title(title).Description(description).Value(&data).Validate(validate),
		)).Run()
	if err != nil {
		return nil, err
	}
	return data, nil
}
