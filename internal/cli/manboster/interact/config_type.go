package interact

import (
	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/spec/cli"
)

const _ADD_ = "add"
const _DELETE_ = "delete"
const _QUIT_ = "quit"
const _EDIT_ = "edit"
const _PURGE_ = "purge"
const _CREATE_ = "create"

var addOption = cli.Option{
	Key:      i18n.T(keys.CliConfigOptionAddNew),
	Value:    _ADD_,
	Selected: false,
}

var quitOption = cli.Option{
	Key:      i18n.T(keys.CliConfigOptionQuit),
	Value:    _QUIT_,
	Selected: false,
}

var purgeOption = cli.Option{
	Key:      i18n.T(keys.CliConfigOptionPurge),
	Value:    _PURGE_,
	Selected: false,
}

var createOption = cli.Option{
	Key:      i18n.T(keys.CliConfigOptionCreateNew),
	Value:    _CREATE_,
	Selected: false,
}
