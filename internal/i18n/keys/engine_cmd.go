package keys

// Engine: command - pair
const (
	CmdPairNoCode       = "engine.cmd.pair_no_code"
	CmdPairInvalidNum   = "engine.cmd.pair_invalid_num"
	CmdPairInvalidRange = "engine.cmd.pair_invalid_range"
	CmdPairSuccess      = "engine.cmd.pair_success"
	CmdPairNoNeed       = "engine.cmd.pair_no_need"
)

// Engine: command - start
const (
	CmdStartWelcome  = "engine.cmd.start_welcome"
	CmdStartFirstUse = "engine.cmd.start_first_use"
	CmdStartNotYours = "engine.cmd.start_not_yours"
	CmdStartCommands = "engine.cmd.start_commands"
)

// Engine: command - help
const (
	CmdHelpHeader    = "engine.cmd.help_header"
	CmdHelpVersion   = "engine.cmd.help_version"
	CmdHelpID        = "engine.cmd.help_id"
	CmdHelpHelp      = "engine.cmd.help_help"
	CmdHelpOp        = "engine.cmd.help_op"
	CmdHelpDeop      = "engine.cmd.help_deop"
	CmdHelpStatus    = "engine.cmd.help_status"
	CmdHelpSave      = "engine.cmd.help_save"
	CmdHelpNew       = "engine.cmd.help_new"
	CmdHelpCompact   = "engine.cmd.help_compact"
	CmdHelpModel     = "engine.cmd.help_model"
	CmdHelpModels    = "engine.cmd.help_models"
	CmdHelpSession   = "engine.cmd.help_session"
	CmdHelpSessions  = "engine.cmd.help_sessions"
	CmdHelpProvider  = "engine.cmd.help_provider"
	CmdHelpProviders = "engine.cmd.help_providers"
	CmdHelpStart     = "engine.cmd.help_start"
	CmdHelpPair      = "engine.cmd.help_pair"
	CmdHelpCancel    = "engine.cmd.help_cancel"
)

// Engine: command - default
const (
	CmdDefaultInvalid = "engine.cmd.default_invalid"
)

// Engine: command - session not active
const (
	CmdSessionNotActive = "engine.cmd.session_not_active"
)

// Engine: command - save
const (
	CmdSaveSuccess = "engine.cmd.save_success"
)

// Engine: command - new
const (
	CmdNewSuccess = "engine.cmd.new_success"
)

// Engine: command - model
const (
	CmdModelList        = "engine.cmd.model_list"
	CmdModelInvalid     = "engine.cmd.model_invalid"
	CmdModelNotFound    = "engine.cmd.model_not_found"
	CmdModelUpdateError = "engine.cmd.model_update_error"
	CmdModelSuccess     = "engine.cmd.model_success"
)

// Engine: command - provider
const (
	CmdProviderList        = "engine.cmd.provider_list"
	CmdProviderInvalid     = "engine.cmd.provider_invalid"
	CmdProviderNotFound    = "engine.cmd.provider_not_found"
	CmdProviderUpdateError = "engine.cmd.provider_update_error"
	CmdProviderSuccess     = "engine.cmd.provider_success"
)

// Engine: command - session
const (
	CmdSessionList        = "engine.cmd.session_list"
	CmdSessionDataError   = "engine.cmd.session_data_error"
	CmdSessionNotFound    = "engine.cmd.session_not_found"
	CmdSessionGetError    = "engine.cmd.session_get_error"
	CmdSessionUpdateError = "engine.cmd.session_update_error"
	CmdSessionSuccess     = "engine.cmd.session_success"
)

// Engine: command - status
const (
	CmdStatusHeader      = "engine.cmd.status_header"
	CmdStatusSummary     = "engine.cmd.status_summary"
	CmdStatusSouls       = "engine.cmd.status_souls"
	CmdStatusTokens      = "engine.cmd.status_tokens"
	CmdStatusCost        = "engine.cmd.status_cost"
	CmdStatusInputPrice  = "engine.cmd.status_input_price"
	CmdStatusOutputPrice = "engine.cmd.status_output_price"
	CmdStatusModelUsage  = "engine.cmd.status_model_usage"
	CmdStatusContext     = "engine.cmd.status_context"
	CmdStatusFullHint    = "engine.cmd.status_full_hint"
)

// Engine: command - op/deop
const (
	CmdOpNotFound      = "engine.cmd.op_not_found"
	CmdOpAlreadyAdmin  = "engine.cmd.op_already_admin"
	CmdOpCreateError   = "engine.cmd.op_create_error"
	CmdOpSuccess       = "engine.cmd.op_success"
	CmdDeopDeleteError = "engine.cmd.deop_delete_error"
	CmdDeopSuccess     = "engine.cmd.deop_success"
	CmdDeopNotFound    = "engine.cmd.deop_not_found"
	CmdDeopFindError   = "engine.cmd.deop_find_error"
)
