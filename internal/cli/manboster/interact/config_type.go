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
	Key:      i18n.T(keys.OptionAddNew),
	Value:    _ADD_,
	Selected: false,
}

var quitOption = cli.Option{
	Key:      i18n.T(keys.OptionQuit),
	Value:    _QUIT_,
	Selected: false,
}

var purgeOption = cli.Option{
	Key:      i18n.T(keys.OptionPurge),
	Value:    _PURGE_,
	Selected: false,
}

var createOption = cli.Option{
	Key:      i18n.T(keys.OptionCreateNew),
	Value:    _CREATE_,
	Selected: false,
}
