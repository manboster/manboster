package huh

import (
	"github.com/charmbracelet/huh"
)

func (h Huh) Prompt(content string, title string, t string, f string) (bool, error) {
	err := h.Display(content, 0)
	defer ClearScreen()
	if err != nil {
		return false, err
	}

	approved := false
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().Title(title).Affirmative(t).Negative(f).Value(&approved),
		)).Run()
	if err != nil {
		return false, err
	}
	return approved, nil
}
