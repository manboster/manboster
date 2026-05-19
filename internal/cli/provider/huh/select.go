package huh

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/x/term"
	"github.com/manboster/manboster/spec/cli"
)

func (h Huh) Select(title string, description string, options []cli.Option, selected string, validate func(option cli.Option) error) (cli.Option, error) {
	ClearScreen()
	opts := BuildHuhOptions(options)
	var sel string
	if selected != "" {
		sel = selected
	}

	err := huh.NewForm(huh.NewGroup(
		huh.NewSelect[string]().Title(title).Description(description).Options(opts...).Value(&sel).Validate(func(s string) error {
			opt, ok := BuildProviderOption(options, sel)
			if ok {
				return validate(opt)
			}
			return fmt.Errorf("invalid option: %s", s)
		}),
	)).Run()
	if err != nil {
		return cli.Option{}, err
	}

	opt, ok := BuildProviderOption(options, sel)
	if !ok {
		return cli.Option{}, fmt.Errorf("unknown option: %s", sel)
	}
	return opt, nil
}

func (h Huh) MultiSelect(title string, description string, options []cli.Option, selected []string, validate func(options []cli.Option) error) ([]cli.Option, error) {
	ClearScreen()
	opts := BuildHuhOptions(options)
	var sels []string
	if selected != nil {
		sels = selected
	}

	_, termHeight, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		termHeight = 24 // fallback
	}
	// Height() sets the visible rows for the option list only.
	// Reserve space for: title (1), description (1), top/bottom borders (2),
	// help line (1), and some breathing room (5) = 10 total.
	height := termHeight - 10
	if height < 3 {
		height = 3
	}

	err = huh.NewForm(huh.NewGroup(
		huh.NewMultiSelect[string]().Title(title).Description(description).Options(opts...).Value(&sels).Height(height).Validate(func(s []string) error {
			opts := BuildProviderOptions(options, sels)
			return validate(opts)
		}),
	)).WithHeight(termHeight).Run()
	if err != nil {
		return nil, err
	}

	optis := BuildProviderOptions(options, sels)
	return optis, nil
}
