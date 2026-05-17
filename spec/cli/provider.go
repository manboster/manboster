package cli

import "time"

type Option[T any] struct {
	Key      string
	Value    T
	Selected bool
}

// Provider defines cli provider's functions
type Provider interface {
	Select(title string, description string, options []Option[any], validate func(option Option[any]) bool) (Option[any], error)
	MultiSelect(title string, description string, options []Option[any], validate func(options []Option[any]) bool) ([]Option[any], error)
	Input(title string, description string, validate func(input any) bool) (any, error)
	Prompt(title string, description string, t string, f string) (bool, error)
	Display(content string, timeout time.Duration) error
}
