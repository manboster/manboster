package huh

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/spec/cli"
)

func (h Huh) Select(title string, description string, options []cli.Option, validate func(option cli.Option) error) (cli.Option, error) {
	opts := BuildHuhOptions(options)
	var sel string

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

func (h Huh) MultiSelect(title string, description string, options []cli.Option, validate func(options []cli.Option) error) ([]cli.Option, error) {
	opts := BuildHuhOptions(options)
	var sels []string

	err := huh.NewForm(huh.NewGroup(
		huh.NewMultiSelect[string]().Title(title).Description(description).Options(opts...).Value(&sels).Validate(func(s []string) error {
			opts := BuildProviderOptions(options, sels)
			return validate(opts)
		}),
	)).Run()
	if err != nil {
		return nil, err
	}

	optis := BuildProviderOptions(options, sels)
	return optis, nil
}
