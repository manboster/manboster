package keys

// Engine: handler
const (
	EngineHandlerToolCallLimit = "engine.handler.tool_call_limit"
	EngineHandlerCompactWait   = "engine.handler.compact_wait"
	EngineHandlerCompactNoNeed = "engine.handler.compact_no_need"
	EngineHandlerCompactError  = "engine.handler.compact_error"
	EngineHandlerReject        = "engine.handler.reject"
	EngineHandlerSelectionGone = "engine.handler.selection_gone"
)

// Engine: onboard handler
const (
	EngineOnboardWelcome     = "engine.onboard.welcome"
	EngineOnboardInstruction = "engine.onboard.instruction"
	EngineOnboardStep1       = "engine.onboard.step1"
	EngineOnboardStep1Note   = "engine.onboard.step1_note"
	EngineOnboardStep2       = "engine.onboard.step2"
	EngineOnboardStep3       = "engine.onboard.step3"
	EngineOnboardWish        = "engine.onboard.wish"
)

// Engine: onboard pair
const (
	EngineOnboardPairSuccess    = "engine.onboard.pair_success"
	EngineOnboardPairSuccessMsg = "engine.onboard.pair_success_msg"
	EngineOnboardPairUserError  = "engine.onboard.pair_user_error"
	EngineOnboardPairFailed     = "engine.onboard.pair_failed"
)

// Engine: gatekeeper buttons
const (
	GatekeeperContinueOnce     = "engine.gatekeeper.continue_once"
	GatekeeperContinueAll      = "engine.gatekeeper.continue_all"
	GatekeeperShutUp           = "engine.gatekeeper.shut_up"
	GatekeeperCancelOnce       = "engine.gatekeeper.cancel_once"
	GatekeeperCancelIgnore     = "engine.gatekeeper.cancel_ignore"
	GatekeeperCancelAll        = "engine.gatekeeper.cancel_all"
	GatekeeperHandleHachimi    = "engine.gatekeeper.handle_hachimi"
	GatekeeperHandleHachimiAll = "engine.gatekeeper.handle_hachimi_all"
	GatekeeperAllow            = "engine.gatekeeper.allow"
	GatekeeperDeny             = "engine.gatekeeper.deny"
)

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

// Engine: load - version tips
const (
	EngineLoadUnstable   = "engine.load.unstable"
	EngineLoadRC         = "engine.load.rc"
	EngineLoadBeta       = "engine.load.beta"
	EngineLoadAlpha      = "engine.load.alpha"
	EngineLoadCanary     = "engine.load.canary"
	EngineLoadCanaryWarn = "engine.load.canary_warn"
)

// Engine: gateway LLM error messages
const (
	GatewayLLMErrorDefault   = "engine.gateway.llm_error_default"
	GatewayLLMErrorRateLimit = "engine.gateway.llm_error_rate_limit"
	GatewayLLMErrorDown      = "engine.gateway.llm_error_down"
	GatewayLLMErrorTimeout   = "engine.gateway.llm_error_timeout"
	GatewayLLMErrorAuth      = "engine.gateway.llm_error_auth"
	GatewayLLMErrorCancelled = "engine.gateway.llm_error_cancelled"
)
const (
	GatekeeperHachimiActivated  = "engine.gatekeeper.hachimi_activated"
	GatekeeperShutUpMsg         = "engine.gatekeeper.shut_up_msg"
	GatekeeperContinueAllMsg    = "engine.gatekeeper.continue_all_msg"
	GatekeeperCancelAllMsg      = "engine.gatekeeper.cancel_all_msg"
	GatekeeperHachimiUnsafe     = "engine.gatekeeper.hachimi_unsafe"
	GatekeeperHachimiSuspicious = "engine.gatekeeper.hachimi_suspicious"
	GatekeeperHachimiReason     = "engine.gatekeeper.hachimi_reason"
	GatekeeperRejectMsg         = "engine.gatekeeper.reject_msg"
	GateKeeperValidateRejectMsg = "engine.gatekeeper.validate_msg"
)

// Engine: util describe
const (
	DescribeToHumanText = "engine.describe.to_human"
	DescribeWithParams  = "engine.describe.with_params"
	DescribeContinue    = "engine.describe.continue"
)
