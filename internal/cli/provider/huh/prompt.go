package huh

import "github.com/charmbracelet/huh"

func (h Huh) Prompt(title string, description string, t string, f string) (bool, error) {
	approved := false
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title(title).Description(description).Affirmative(t).Negative(f),
		)).Run()
	if err != nil {
		return false, err
	}
	return approved, nil
}
