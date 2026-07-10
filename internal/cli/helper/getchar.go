package helper

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"golang.org/x/term"
)

// GetChar helps you get a character from terminal input
func GetChar() string {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}

	defer func(fd int, oldState *term.State) {
		err := term.Restore(fd, oldState)
		if err != nil {
			color.Red(err.Error())
		}
	}(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)

	_, err = os.Stdin.Read(b)
	if err != nil {
		panic(err)
	}

	char := string(b[0])

	err = term.Restore(int(os.Stdin.Fd()), oldState)
	if err != nil {
		color.Red(err.Error())
		return ""
	}

	fmt.Println()

	return char
}
