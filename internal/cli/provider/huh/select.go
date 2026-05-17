package huh

import "github.com/manboster/manboster/spec/cli"

func (h Huh) Select(title string, description string, options []cli.Option[any], validate func(option cli.Option[any]) bool) (cli.Option[any], error) {
	//TODO implement me
	panic("implement me")
}

func (h Huh) MultiSelect(title string, description string, options []cli.Option[any], validate func(options []cli.Option[any]) bool) ([]cli.Option[any], error) {
	//TODO implement me
	panic("implement me")
}
