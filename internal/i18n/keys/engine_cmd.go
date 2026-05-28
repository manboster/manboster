package keys

// Engine: command - pair
const (
	CmdPairNoCode       = "engine.cmd.pair.no_code"
	CmdPairInvalidNum   = "engine.cmd.pair.invalid_num"
	CmdPairInvalidRange = "engine.cmd.pair.invalid_range"
	CmdPairSuccess      = "engine.cmd.pair.success"
	CmdPairNoNeed       = "engine.cmd.pair.no_need"
)

// Engine: command - start
const (
	CmdStartWelcome  = "engine.cmd.start.welcome"
	CmdStartFirstUse = "engine.cmd.start.first_use"
	CmdStartNotYours = "engine.cmd.start.not_yours"
	CmdStartCommands = "engine.cmd.start.commands"
)

// Engine: command - help
const (
	CmdHelpHeader    = "engine.cmd.help.header"
	CmdHelpVersion   = "engine.cmd.help.version"
	CmdHelpID        = "engine.cmd.help.id"
	CmdHelpHelp      = "engine.cmd.help.help"
	CmdHelpOp        = "engine.cmd.help.op"
	CmdHelpDeop      = "engine.cmd.help.deop"
	CmdHelpStatus    = "engine.cmd.help.status"
	CmdHelpSave      = "engine.cmd.help.save"
	CmdHelpNew       = "engine.cmd.help.new"
	CmdHelpCompact   = "engine.cmd.help.compact"
	CmdHelpModel     = "engine.cmd.help.model"
	CmdHelpModels    = "engine.cmd.help.models"
	CmdHelpSession   = "engine.cmd.help.session"
	CmdHelpSessions  = "engine.cmd.help.sessions"
	CmdHelpProvider  = "engine.cmd.help.provider"
	CmdHelpProviders = "engine.cmd.help.providers"
	CmdHelpStart     = "engine.cmd.help.start"
	CmdHelpPair      = "engine.cmd.help.pair"
	CmdHelpCancel    = "engine.cmd.help.cancel"
)

// Engine: command - default
const (
	CmdDefaultInvalid = "engine.cmd.default.invalid"
)

// Engine: command - save
const (
	CmdSaveSuccess = "engine.cmd.save.success"
)

// Engine: command - new
const (
	CmdNewSuccess = "engine.cmd.new.success"
)

// Engine: command - model
const (
	CmdModelList        = "engine.cmd.model.list"
	CmdModelInvalid     = "engine.cmd.model.invalid"
	CmdModelNotFound    = "engine.cmd.model.not_found"
	CmdModelUpdateError = "engine.cmd.model.update_error"
	CmdModelSuccess     = "engine.cmd.model.success"
	CmdModelInfo        = "engine.cmd.model.info"
)

// Engine: command - provider
const (
	CmdProviderList        = "engine.cmd.provider.list"
	CmdProviderInvalid     = "engine.cmd.provider.invalid"
	CmdProviderNotFound    = "engine.cmd.provider.not_found"
	CmdProviderUpdateError = "engine.cmd.provider.update_error"
	CmdProviderSuccess     = "engine.cmd.provider.success"
	CmdProviderInfo        = "engine.cmd.provider.info"
)

// Engine: command - session
const (
	CmdSessionList        = "engine.cmd.session.list"
	CmdSessionDataError   = "engine.cmd.session.data_error"
	CmdSessionNotFound    = "engine.cmd.session.not_found"
	CmdSessionGetError    = "engine.cmd.session.get_error"
	CmdSessionUpdateError = "engine.cmd.session.update_error"
	CmdSessionSuccess     = "engine.cmd.session.success"
	CmdSessionNotActive   = "engine.cmd.session.not_active"
	CmdSessionInfo        = "engine.cmd.session.info"
)

// Engine: command - status
const (
	CmdStatusHeader      = "engine.cmd.status.header"
	CmdStatusSummary     = "engine.cmd.status.summary"
	CmdStatusSouls       = "engine.cmd.status.souls"
	CmdStatusTokens      = "engine.cmd.status.tokens"
	CmdStatusCost        = "engine.cmd.status.cost"
	CmdStatusInputPrice  = "engine.cmd.status.input_price"
	CmdStatusOutputPrice = "engine.cmd.status.output_price"
	CmdStatusModelUsage  = "engine.cmd.status.model_usage"
	CmdStatusContext     = "engine.cmd.status.context"
	CmdStatusFullHint    = "engine.cmd.status.full_hint"
)

// Engine: command - op/deop
const (
	CmdOpNotFound     = "engine.cmd.op.not_found"
	CmdOpAlreadyAdmin = "engine.cmd.op.already_admin"
	CmdOpCreateError  = "engine.cmd.op.create_error"
	CmdOpSuccess      = "engine.cmd.op.success"

	CmdDeopDeleteError = "engine.cmd.deop.delete_error"
	CmdDeopSuccess     = "engine.cmd.deop.success"
	CmdDeopNotFound    = "engine.cmd.deop.not_found"
	CmdDeopFindError   = "engine.cmd.deop.find_error"
)
