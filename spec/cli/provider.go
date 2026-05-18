package cli

import "time"

type Option struct {
	Key      string
	Value    string
	Selected bool
}

// Provider defines cli provider's functions
type Provider interface {
	Select(title string, description string, options []Option, selected string, validate func(option Option) error) (Option, error)
	MultiSelect(title string, description string, options []Option, selected []string, validate func(options []Option) error) ([]Option, error)
	Input(title string, description string, defaultVal string, validate func(input string) error) (any, error)
	Prompt(content string, title string, t string, f string) (bool, error)
	Display(content string, timeout time.Duration) error
	Alert(title string, description string) error
}
