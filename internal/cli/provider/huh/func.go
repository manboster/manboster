package huh

import (
	"github.com/charmbracelet/huh"
	"github.com/manboster/manboster/spec/cli"
)

func BuildHuhOptions(options []cli.Option) []huh.Option[string] {
	var ops []huh.Option[string]
	for _, o := range options {
		opt := huh.NewOption[string](o.Key, o.Value)
		if o.Selected {
			opt.Selected(o.Selected)
		}
		ops = append(ops, opt)
	}
	return ops
}

func BuildProviderOption(options []cli.Option, resp string) (cli.Option, bool) {
	for _, option := range options {
		if option.Value == resp {
			return option, true
		}
	}
	return cli.Option{}, false
}

func BuildProviderOptions(options []cli.Option, resp []string) []cli.Option {
	var ops []cli.Option
	for _, r := range resp {
		for _, o := range options {
			if o.Value == r {
				ops = append(ops, o)
				break
			}
		}
	}
	return ops
}
