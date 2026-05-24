package interact

import (
	"errors"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/cli"
)

func handle[T ~string](p cli.Provider, f *configForm[T], options []cli.Option, title string, content string) error {
	var option cli.Option
	for {
		var err error

		option, err = p.Select(title, content, options, option.Value, func(option cli.Option) error {
			return nil
		})
		if err != nil {
			return err
		}

		err = f.Handle(T(option.Value))
		if errors.As(errQuit, &err) {
			return nil
		}
		if err != nil {
			return err
		}

		if option.Value == _QUIT_ {
			color.Yellow(i18n.T(keys.OptionBye))
			return nil
		}
	}
}

func handleWithPrompt[T ~string](p cli.Provider, f *configForm[T], options []cli.Option, content string, title string) error {
	var option cli.Option
	for {
		var err error
		err = p.Display(content, 0)
		if err != nil {
			return err
		}

		option, err = p.Select(title, "", options, option.Value, func(option cli.Option) error {
			return nil
		})
		if err != nil {
			return err
		}

		err = f.Handle(T(option.Value))
		if err != nil {
			return err
		}

		if option.Value == _QUIT_ {
			color.Yellow(i18n.T(keys.OptionBye))
			return nil
		}
	}
}
