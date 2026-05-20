package cli

import (
	"fmt"

	"github.com/manboster/manboster/spec/schema"
)

func BuildOptions[T buildable](options []T, selected []string) []Option {
	return buildOptions(options, selected, func(option T) Option {
		return Option{
			Key:   option.DisplayName(),
			Value: option.Name(),
		}
	})
}

type buildableWithDescription interface {
	DisplayName() string
	Name() string
	Description() string
}

func BuildOptionsWithDescription[T buildableWithDescription](options []T, selected []string) []Option {
	return buildOptions[T](options, selected, func(option T) Option {
		return Option{
			Key:   fmt.Sprintf("%s\n%s", option.DisplayName(), option.Description()),
			Value: option.Name(),
		}
	})
}

type buildableWithMetadata interface {
	DisplayName() string
	Name() string
	Metadata() schema.MetaData
}

type buildableModel interface {
	GetDisplayName() string
	GetName() string
}

func BuildModelOptions[T buildableModel](options []T, selected []string) []Option {
	var opts []*BuildableType
	for _, option := range options {
		opts = append(opts, NewBuildableType(option.GetName(), option.GetDisplayName()))
	}
	return buildOptions[*BuildableType](opts, selected, func(option *BuildableType) Option {
		return Option{
			Key:   option.Name(),
			Value: option.DisplayName(),
		}
	})
}

func BuildStringOptions(options []string, selected []string) []Option {
	var opts []*BuildableType
	for _, option := range options {
		opts = append(opts, NewBuildableType(option, option))
	}
	return buildOptions[*BuildableType](opts, selected, func(option *BuildableType) Option {
		return Option{
			Key:   option.Name(),
			Value: option.DisplayName(),
		}
	})
}
