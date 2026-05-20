package interact

import "github.com/manboster/manboster/spec/cli"

const _ADD_ = "add"
const _DELETE_ = "delete"
const _QUIT_ = "quit"
const _EDIT_ = "edit"
const _PURGE_ = "purge"
const _CREATE_ = "create"

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

var purgeOption = cli.Option{
	Key:      "Purge Unused Session Data (Save Space)",
	Value:    _PURGE_,
	Selected: false,
}

var createOption = cli.Option{
	Key:      "Create a new one",
	Value:    _CREATE_,
	Selected: false,
}
