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

type buildableWithMetadata interface {
	DisplayName() string
	Name() string
	Description() string
	MetaData() schema.MetaData
}

func BuildOptionsWithMetadata[T buildableWithMetadata](options []T, selected []string) []Option {
	return buildOptions[T](options, selected, func(option T) Option {
		metadata := option.MetaData()

		displayName := option.DisplayName()
		description := option.Description()
		if metadata.DescriptionForUser != "" {
			description = metadata.DescriptionForUser
		}
		if metadata.DisplayNameForUser != "" {
			displayName = metadata.DisplayNameForUser
		}

		return Option{
			Key:   fmt.Sprintf("%s\n%s", displayName, description),
			Value: option.Name(),
		}
	})
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
