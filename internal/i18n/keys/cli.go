package keys

// CLI: manboster app
const (
	CLIManbosterWelcome             = "cli.manboster.welcome"
	CLIManbosterLoading             = "cli.manboster.loading"
	CLIManbosterDaemonRunningError  = "cli.manboster.daemon_running_error"
	CLIManbosterQuiting             = "cli.manboster.quiting"
	CLIManbosterConfigNotFound      = "cli.manboster.config_not_found"
	CLIManbosterConfigCreatedSuccess = "cli.manboster.config_created_success"
	CLIManbosterConfigInitError     = "cli.manboster.config_init_error"
	CLIManbosterReadingConfig       = "cli.manboster.reading_config"
	CLIManbosterLoaderError         = "cli.manboster.loader_error"
	CLIManbosterGoodbye             = "cli.manboster.goodbye"
)

// CLI: daemon
const (
	CLIDaemonStartError      = "cli.daemon.start_error"
	CLIDaemonStartSuccess    = "cli.daemon.start_success"
	CLIDaemonStopError       = "cli.daemon.stop_error"
	CLIDaemonStopSuccess     = "cli.daemon.stop_success"
	CLIDaemonRestartMessage  = "cli.daemon.restart_message"
	CLIDaemonStatusRunning   = "cli.daemon.status_running"
	CLIDaemonStatusStopped   = "cli.daemon.status_stopped"
	CLIDaemonStatusNotRunning = "cli.daemon.status_not_running"
	CLIDaemonStatusError     = "cli.daemon.status_error"
	CLIDaemonLogError        = "cli.daemon.log_error"
	CLIDaemonLogReading      = "cli.daemon.log_reading"
	CLIDaemonNoConfig        = "cli.daemon.no_config"
)

// CLI: config wizard
const (
	CLIConfigDaemonRunningError  = "cli.config.daemon_running_error"
	CLIConfigInitError           = "cli.config.init_error"
	CLIConfigError               = "cli.config.error"
	CLIConfigEditNotFound        = "cli.config.edit_not_found"
	CLIConfigEditCreatePrompt    = "cli.config.edit_create_prompt"
	CLIConfigEditCancelled       = "cli.config.edit_cancelled"
	CLIConfigEditInitError       = "cli.config.edit_init_error"
	CLIConfigEditOpenError       = "cli.config.edit_open_error"
	CLIConfigOpenNotFound        = "cli.config.open_not_found"
	CLIConfigOpenInitError       = "cli.config.open_init_error"
	CLIConfigOpenError           = "cli.config.open_error"
	CLIConfigOnboardError        = "cli.config.onboard_error"
	CLIConfigOnboardValidateError = "cli.config.onboard_validate_error"
	CLIConfigOnboardWriteError   = "cli.config.onboard_write_error"
	CLIConfigOnboardSuccess      = "cli.config.onboard_success"
	CLIConfigSelectPrompt        = "cli.config.select_prompt"
	CLIConfigSelectHelp          = "cli.config.select_help"
	CLIConfigWizardError         = "cli.config.wizard_error"
	CLIConfigChatSelectPrompt    = "cli.config.chat_select_prompt"
	CLIConfigChatSelectHelp      = "cli.config.chat_select_help"
	CLIConfigChatDeleteConfirm   = "cli.config.chat_delete_confirm"
	CLIConfigChatDeleteSuccess   = "cli.config.chat_delete_success"
	CLIConfigChatDeleteCancelled = "cli.config.chat_delete_cancelled"
	CLIConfigLLMSelectPrompt     = "cli.config.llm_select_prompt"
	CLIConfigLLMSelectHelp       = "cli.config.llm_select_help"
	CLIConfigLLMDeleteConfirm    = "cli.config.llm_delete_confirm"
	CLIConfigLLMDeleteSuccess    = "cli.config.llm_delete_success"
	CLIConfigLLMDeleteCancelled  = "cli.config.llm_delete_cancelled"
	CLIConfigToolSelectPrompt    = "cli.config.tool_select_prompt"
	CLIConfigToolSelectHelp      = "cli.config.tool_select_help"
	CLIConfigToolDeleteConfirm   = "cli.config.tool_delete_confirm"
	CLIConfigToolDeleteSuccess   = "cli.config.tool_delete_success"
	CLIConfigToolDeleteCancelled = "cli.config.tool_delete_cancelled"
	CLIConfigToolNoConfig        = "cli.config.tool_no_config"
	CLIConfigHachimiSelectPrompt  = "cli.config.hachimi_select_prompt"
	CLIConfigHachimiSelectHelp    = "cli.config.hachimi_select_help"
	CLIConfigHachimiDeleteConfirm = "cli.config.hachimi_delete_confirm"
	CLIConfigHachimiDeleteSuccess = "cli.config.hachimi_delete_success"
	CLIConfigHachimiSetDefault    = "cli.config.hachimi_set_default"
)

// CLI: onboard wizard
const (
	CLIOnboardExistingConfigPrompt = "cli.onboard.existing_config_prompt"
	CLIOnboardUserCancelled        = "cli.onboard.user_cancelled"
	CLIOnboardWarningRejected      = "cli.onboard.warning_rejected"
	CLIOnboardWelcome              = "cli.onboard.welcome"
	CLIOnboardWelcomeMsg           = "cli.onboard.welcome_msg"
	CLIOnboardWizardErrorPrompt    = "cli.onboard.wizard_error_prompt"
	CLIOnboardWizardErrorRetry     = "cli.onboard.wizard_error_retry"
	CLIOnboardSuccess              = "cli.onboard.success"
	CLIOnboardSuccessMsg           = "cli.onboard.success_msg"
	CLIOnboardChatSelectPrompt     = "cli.onboard.chat_select_prompt"
	CLIOnboardChatAddedCount       = "cli.onboard.chat_added_count"
	CLIOnboardChatNoMoreProviders  = "cli.onboard.chat_no_more_providers"
	CLIOnboardLLMSelectPrompt      = "cli.onboard.llm_select_prompt"
	CLIOnboardLLMAddedCount        = "cli.onboard.llm_added_count"
	CLIOnboardAppSelectProvider    = "cli.onboard.app_select_provider"
	CLIOnboardAppSelectModel       = "cli.onboard.app_select_model"
	CLIOnboardHachimiFeaturePrompt = "cli.onboard.hachimi_feature_prompt"
	CLIOnboardHachimiEnableQuestion = "cli.onboard.hachimi_enable_question"
	CLIOnboardHachimiAddedCount    = "cli.onboard.hachimi_added_count"
	CLIOnboardHachimiSelectDefault = "cli.onboard.hachimi_select_default"
	CLIOnboardHachimiSelectProvider = "cli.onboard.hachimi_select_provider"
	CLIOnboardToolSelectPrompt     = "cli.onboard.tool_select_prompt"
	CLIOnboardToolSelectHelp       = "cli.onboard.tool_select_help"
	CLIOnboardPreviewTitle         = "cli.onboard.preview_title"
	CLIOnboardPreviewChatCount     = "cli.onboard.preview_chat_count"
	CLIOnboardPreviewLLMCount      = "cli.onboard.preview_llm_count"
	CLIOnboardPreviewToolCount     = "cli.onboard.preview_tool_count"
	CLIOnboardPreviewHachimiEnabled = "cli.onboard.preview_hachimi_enabled"
	CLIOnboardPreviewHachimiDisabled = "cli.onboard.preview_hachimi_disabled"
	CLIOnboardPreviewConfirm       = "cli.onboard.preview_confirm"
	CLIOnboardWriteConfirm         = "cli.onboard.write_confirm"
	CLIOnboardWritePathPrompt      = "cli.onboard.write_path_prompt"
	CLIOnboardWritePathHelp        = "cli.onboard.write_path_help"
	CLIOnboardWriteSuccess         = "cli.onboard.write_success"
)

// Config validation
const (
	ConfigValidateUnsupportedVersion      = "config.validate.unsupported_version"
	ConfigValidateOutdatedVersion         = "config.validate.outdated_version"
	ConfigValidateMissingChat             = "config.validate.missing_chat"
	ConfigValidateMissingLLM              = "config.validate.missing_llm"
	ConfigValidateHachimiNoProviders      = "config.validate.hachimi_no_providers"
	ConfigValidateMissingDBPath           = "config.validate.missing_db_path"
	ConfigValidateMissingDefaultLLMProvider = "config.validate.missing_default_llm_provider"
	ConfigValidateMissingDefaultLLMModel  = "config.validate.missing_default_llm_model"
)

// Engine: onboard
const (
	EngineOnboardRetryLimitExceeded = "engine.onboard.retry_limit_exceeded"
	EngineOnboardPairCodeMessage    = "engine.onboard.pair_code_message"
)

// App: main
const (
	AppWelcome              = "app.welcome"
	AppLoading              = "app.loading"
	AppDaemonRunning        = "app.daemon_running"
	AppDaemonRunningQuit    = "app.daemon_running_quit"
	AppConfigNotFound       = "app.config_not_found"
	AppConfigCreated        = "app.config_created"
	AppConfigInitError      = "app.config_init_error"
	AppReadingConfig        = "app.reading_config"
	AppLoaderError          = "app.loader_error"
	AppGoodbye              = "app.goodbye"
)

// Daemon: execute
const (
	DaemonStartError        = "daemon.start_error"
	DaemonStartSuccess      = "daemon.start_success"
	DaemonStopSuccess       = "daemon.stop_success"
	DaemonStopStopped       = "daemon.stop_stopped"
	DaemonStopError         = "daemon.stop_error"
	DaemonRestartMessage    = "daemon.restart_message"
	DaemonStatusRunning     = "daemon.status_running"
	DaemonStatusNotRunning  = "daemon.status_not_running"
	DaemonStatusError       = "daemon.status_error"
	DaemonNoConfig          = "daemon.no_config"
	DaemonLogError          = "daemon.log_error"
	DaemonLogReading        = "daemon.log_reading"
)

// Wizard: shared alerts
const (
	WizardTitle             = "wizard.title"
	WizardErrorAlert        = "wizard.error_alert"
	WizardWelcome           = "wizard.welcome"
	WizardSuccess           = "wizard.success"
	WizardErrorRetry        = "wizard.error_retry"
	WizardConfigError       = "wizard.config_error"
)

// Wizard: entrypoint
const (
	EntrypointSelectPrompt  = "wizard.entrypoint.select_prompt"
	EntrypointSelectHelp    = "wizard.entrypoint.select_help"
	EntrypointDatabase      = "wizard.entrypoint.database"
	EntrypointConfig        = "wizard.entrypoint.config"
	EntrypointEditor        = "wizard.entrypoint.editor"
	EntrypointQuit          = "wizard.entrypoint.quit"
)

// Wizard: config landing
const (
	ConfigLandingSelectPrompt  = "wizard.config_landing.select_prompt"
	ConfigLandingSelectHelp    = "wizard.config_landing.select_help"
	ConfigLandingChat          = "wizard.config_landing.chat"
	ConfigLandingLLM           = "wizard.config_landing.llm"
	ConfigLandingTool          = "wizard.config_landing.tool"
	ConfigLandingHachimi       = "wizard.config_landing.hachimi"
	ConfigLandingApp           = "wizard.config_landing.app"
	ConfigLandingQuit          = "wizard.config_landing.quit"
)

// Wizard: provider actions (shared)
const (
	ActionDeleteProvider    = "wizard.action.delete_provider"
	ActionEditProvider      = "wizard.action.edit_provider"
	ActionSetDefault        = "wizard.action.set_default"
	ActionQuit              = "wizard.action.quit"
	ActionWhatToDo          = "wizard.action.what_to_do"
)

// Wizard: onboard existing config
const (
	OnboardExistingConfig   = "wizard.onboard.existing_config"
	OnboardExistingContinue = "wizard.onboard.existing_continue"
	OnboardUserCancelled    = "wizard.onboard.user_cancelled"
	OnboardWarningRejected  = "wizard.onboard.warning_rejected"
)

// Wizard: onboard chat
const (
	OnboardChatSelectPrompt    = "wizard.onboard.chat_select_prompt"
	OnboardChatAddedCount      = "wizard.onboard.chat_added_count"
	OnboardChatNoMoreProviders = "wizard.onboard.chat_no_more_providers"
	OnboardChatConfigError     = "wizard.onboard.chat_config_error"
)

// Wizard: onboard LLM
const (
	OnboardLLMSelectPrompt  = "wizard.onboard.llm_select_prompt"
	OnboardLLMAddedCount    = "wizard.onboard.llm_added_count"
	OnboardLLMConfigError   = "wizard.onboard.llm_config_error"
)

// Wizard: onboard app
const (
	OnboardAppSelectProvider = "wizard.onboard.app_select_provider"
	OnboardAppSelectHelp     = "wizard.onboard.app_select_help"
	OnboardAppSelectModel    = "wizard.onboard.app_select_model"
)

// Wizard: onboard hachimi
const (
	OnboardHachimiFeaturePrompt  = "wizard.onboard.hachimi_feature_prompt"
	OnboardHachimiEnableQuestion = "wizard.onboard.hachimi_enable_question"
	OnboardHachimiAddedCount     = "wizard.onboard.hachimi_added_count"
	OnboardHachimiSelectDefault  = "wizard.onboard.hachimi_select_default"
	OnboardHachimiSelectHelp     = "wizard.onboard.hachimi_select_help"
	OnboardHachimiSelectProvider = "wizard.onboard.hachimi_select_provider"
	OnboardHachimiNoMore         = "wizard.onboard.hachimi_no_more"
	OnboardHachimiConfigError    = "wizard.onboard.hachimi_config_error"
)

// Wizard: onboard tool
const (
	OnboardToolSelectPrompt = "wizard.onboard.tool_select_prompt"
	OnboardToolSelectHelp   = "wizard.onboard.tool_select_help"
)

// Wizard: onboard preview
const (
	OnboardPreviewTitle          = "wizard.onboard.preview_title"
	OnboardPreviewRestart        = "wizard.onboard.preview_restart"
	OnboardPreviewChatCount      = "wizard.onboard.preview_chat_count"
	OnboardPreviewLLMCount       = "wizard.onboard.preview_llm_count"
	OnboardPreviewToolCount      = "wizard.onboard.preview_tool_count"
	OnboardPreviewHachimiEnabled = "wizard.onboard.preview_hachimi_enabled"
	OnboardPreviewHachimiDisabled = "wizard.onboard.preview_hachimi_disabled"
	OnboardPreviewContinue       = "wizard.onboard.preview_continue"
	OnboardPreviewConfirm        = "wizard.onboard.preview_confirm"
	OnboardPreviewProblem        = "wizard.onboard.preview_problem"
)

// Wizard: write config
const (
	OnboardWriteExisting    = "wizard.onboard.write_existing"
	OnboardWriteConfirm     = "wizard.onboard.write_confirm"
	OnboardWritePathPrompt  = "wizard.onboard.write_path_prompt"
	OnboardWritePathHelp    = "wizard.onboard.write_path_help"
	OnboardWriteError       = "wizard.onboard.write_error"
	OnboardWriteSuccess     = "wizard.onboard.write_success"
)

// Wizard: config chat
const (
	ConfigChatSelectPrompt  = "wizard.config.chat_select_prompt"
	ConfigChatSelectHelp    = "wizard.config.chat_select_help"
	ConfigChatDeleteConfirm = "wizard.config.chat_delete_confirm"
	ConfigChatDeleteSuccess = "wizard.config.chat_delete_success"
)

// Wizard: config LLM
const (
	ConfigLLMSelectPrompt   = "wizard.config.llm_select_prompt"
	ConfigLLMSelectHelp     = "wizard.config.llm_select_help"
	ConfigLLMDeleteConfirm  = "wizard.config.llm_delete_confirm"
	ConfigLLMDeleteSuccess  = "wizard.config.llm_delete_success"
)

// Wizard: config tool
const (
	ConfigToolSelectPrompt   = "wizard.config.tool_select_prompt"
	ConfigToolSelectHelp     = "wizard.config.tool_select_help"
	ConfigToolDeleteConfirm  = "wizard.config.tool_delete_confirm"
	ConfigToolDeleteSuccess  = "wizard.config.tool_delete_success"
	ConfigToolNoConfig       = "wizard.config.tool_no_config"
)

// Wizard: config hachimi
const (
	ConfigHachimiSelectPrompt  = "wizard.config.hachimi_select_prompt"
	ConfigHachimiSelectHelp    = "wizard.config.hachimi_select_help"
	ConfigHachimiSetDefault    = "wizard.config.hachimi_set_default"
	ConfigHachimiDeleteConfirm = "wizard.config.hachimi_delete_confirm"
	ConfigHachimiDeleteSuccess = "wizard.config.hachimi_delete_success"
)

// Wizard: database landing
const (
	DatabaseSelectPrompt    = "wizard.database.select_prompt"
	DatabaseSelectHelp      = "wizard.database.select_help"
	DatabaseSessions        = "wizard.database.sessions"
	DatabaseUsers           = "wizard.database.users"
	DatabaseSouls           = "wizard.database.souls"
	DatabaseQuit            = "wizard.database.quit"
)

// Wizard: database soul
const (
	SoulSelectPrompt        = "wizard.soul.select_prompt"
	SoulEditAction          = "wizard.soul.edit_action"
	SoulDeleteAction        = "wizard.soul.delete_action"
	SoulEditContent         = "wizard.soul.edit_content"
	SoulEditContentHelp     = "wizard.soul.edit_content_help"
	SoulUpdatedSuccess      = "wizard.soul.updated_success"
	SoulDeleteConfirm       = "wizard.soul.delete_confirm"
	SoulDeletedSuccess      = "wizard.soul.deleted_success"
	SoulCreatedSuccess      = "wizard.soul.created_success"
	SoulNameInput           = "wizard.soul.name_input"
	SoulNameHelp            = "wizard.soul.name_help"
	SoulContentInput        = "wizard.soul.content_input"
	SoulContentHelp         = "wizard.soul.content_help"
	SoulScopeInput          = "wizard.soul.scope_input"
	SoulScopeHelp           = "wizard.soul.scope_help"
)

// Wizard: config run (command_run)
const (
	ConfigRunDaemonRunning   = "wizard.config_run.daemon_running"
	ConfigRunInitError       = "wizard.config_run.init_error"
	ConfigRunError           = "wizard.config_run.error"
	ConfigEditNotFound       = "wizard.config_run.edit_not_found"
	ConfigEditCreatePrompt   = "wizard.config_run.edit_create_prompt"
	ConfigEditCancelled      = "wizard.config_run.edit_cancelled"
	ConfigEditOpenError      = "wizard.config_run.edit_open_error"
	ConfigOpenNotFound       = "wizard.config_run.open_not_found"
	ConfigOpenError          = "wizard.config_run.open_error"
	ConfigOnboardError       = "wizard.config_run.onboard_error"
	ConfigOnboardValidateErr = "wizard.config_run.onboard_validate_error"
	ConfigOnboardWriteError  = "wizard.config_run.onboard_write_error"
	ConfigOnboardSuccess     = "wizard.config_run.onboard_success"
)

// Wizard: shared option labels
const (
	OptionAddNew    = "wizard.option.add_new"
	OptionQuit      = "wizard.option.quit"
	OptionPurge     = "wizard.option.purge"
	OptionCreateNew = "wizard.option.create_new"
	OptionBye       = "wizard.option.bye"
)

// Wizard: database user
const (
	UserSelectPrompt   = "wizard.user.select_prompt"
	UserPromoteAction  = "wizard.user.promote_action"
	UserDegradeAction  = "wizard.user.degrade_action"
	UserDeleteAction   = "wizard.user.delete_action"
	UserNoUsers        = "wizard.user.no_users"
	UserPromoteConfirm = "wizard.user.promote_confirm"
	UserPromoteSuccess = "wizard.user.promote_success"
	UserDegradeConfirm = "wizard.user.degrade_confirm"
	UserDegradeSuccess = "wizard.user.degrade_success"
	UserDeleteConfirm  = "wizard.user.delete_confirm"
	UserDeleteSuccess  = "wizard.user.delete_success"
)

// Wizard: database session
const (
	SessionSelectPrompt    = "wizard.session.select_prompt"
	SessionEditAction      = "wizard.session.edit_action"
	SessionDeleteAction    = "wizard.session.delete_action"
	SessionSelectProvider  = "wizard.session.select_provider"
	SessionSelectModel     = "wizard.session.select_model"
	SessionUpdateSuccess   = "wizard.session.update_success"
	SessionDeleteConfirm   = "wizard.session.delete_confirm"
	SessionDeleteSuccess   = "wizard.session.delete_success"
	SessionPurgeConfirm    = "wizard.session.purge_confirm"
	SessionPurgeSuccess    = "wizard.session.purge_success"
	SessionPurgeError      = "wizard.session.purge_error"
	SessionPurgeClean      = "wizard.session.purge_clean"
	SessionChatDeleteError = "wizard.session.chat_delete_error"
	SessionDataDeleteError = "wizard.session.data_delete_error"
)

// Wizard: onboard warning
const (	OnboardWarningRiskTitle      = "wizard.onboard.warning_risk_title"
	OnboardWarningRiskPrompt     = "wizard.onboard.warning_risk_prompt"
	OnboardWarningUnstableTitle  = "wizard.onboard.warning_unstable_title"
	OnboardWarningUnstablePrompt = "wizard.onboard.warning_unstable_prompt"
	OnboardWarningAccept         = "wizard.onboard.warning_accept"
	OnboardWarningExit           = "wizard.onboard.warning_exit"
)
