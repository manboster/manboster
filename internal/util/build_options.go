package util

import (
	"github.com/manboster/manboster/spec/cli"
	"github.com/manboster/manboster/spec/config"
)

func BuildOptions[T configurable](providers []T, selected []string) []cli.Option {
	var options []cli.Option
	var configProviders []config.Provider
	for _, p := range providers {
		configProviders = append(configProviders, p.Config())
	}

	for i, provider := range configProviders {
		option := cli.Option{
			Key:   provider.DisplayName(),
			Value: provider.Name(),
		}

		if i == 0 && (len(selected) == 0 || selected == nil) {
			option.Selected = true
		} else if selected != nil && len(selected) > 0 {
			for j, selectedItem := range selected {
				if selectedItem == provider.Name() {
					option.Selected = true
					selected = append(selected[:j], selected[j+1:]...)
					break
				}
			}
		}

		options = append(options, option)
	}
	return options
}

type configurable interface {
	Config() config.Provider
}
