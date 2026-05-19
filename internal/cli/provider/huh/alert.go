package huh

import (
	"github.com/charmbracelet/huh"
)

func (h Huh) Alert(title string, description string) error {
	ClearScreen()
	return huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title(title).Description(description).Next(true),
		)).Run()
}
