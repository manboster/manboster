package cli

func buildOptions[T buildable](options []T, selected []string, buildFunc func(option T) Option) []Option {
	var opts []Option

	for index, option := range options {
		opt := buildFunc(option)
		if len(selected) > 0 {
			for i, selectedOption := range selected {
				if selectedOption == opt.Value {
					opt.Selected = true
					selected = append(selected[:i], selected[i+1:]...)
					break
				}
			}
		} else if selected == nil {
			if index == 0 {
				opt.Selected = true
			}
		}
		opts = append(opts, opt)
	}
	return opts
}

type buildable interface {
	DisplayName() string
	Name() string
}
