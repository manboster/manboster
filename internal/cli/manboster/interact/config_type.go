package interact

import "github.com/manboster/manboster/spec/cli"

const _ADD_ = "add"
const _DELETE_ = "delete"
const _QUIT_ = "quit"
const _EDIT_ = "edit"

var addOption = cli.Option{
	Key:      "Add a new one",
	Value:    _ADD_,
	Selected: false,
}

var quitOption = cli.Option{
	Key:      "Quit",
	Value:    _QUIT_,
	Selected: false,
}
