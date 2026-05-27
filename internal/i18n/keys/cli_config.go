package keys

// cli.config: entrypoint
const (
	CliConfigEntrypointSelectPrompt = "cli.config.entrypoint.select_prompt"
	CliConfigEntrypointSelectHelp   = "cli.config.entrypoint.select_help"
	CliConfigEntrypointDatabase     = "cli.config.entrypoint.database"
	CliConfigEntrypointConfig       = "cli.config.entrypoint.config"
	CliConfigEntrypointEditor       = "cli.config.entrypoint.editor"
	CliConfigEntrypointQuit         = "cli.config.entrypoint.quit"
)

// cli.config: config landing
const (
	CliConfigLandingSelectPrompt = "cli.config.landing.select_prompt"
	CliConfigLandingSelectHelp   = "cli.config.landing.select_help"
	CliConfigLandingChat         = "cli.config.landing.chat"
	CliConfigLandingLLM          = "cli.config.landing.llm"
	CliConfigLandingTool         = "cli.config.landing.tool"
	CliConfigLandingHachimi      = "cli.config.landing.hachimi"
	CliConfigLandingApp          = "cli.config.landing.app"
	CliConfigLandingQuit         = "cli.config.landing.quit"
)

// cli.config: provider actions (shared)
const (
	CliConfigActionDeleteProvider = "cli.config.action.delete_provider"
	CliConfigActionEditProvider   = "cli.config.action.edit_provider"
	CliConfigActionSetDefault     = "cli.config.action.set_default"
	CliConfigActionQuit           = "cli.config.action.quit"
	CliConfigActionWhatToDo       = "cli.config.action.what_to_do"
)

// cli.config: config chat
const (
	CliConfigChatSelectPrompt  = "cli.config.chat.select_prompt"
	CliConfigChatSelectHelp    = "cli.config.chat.select_help"
	CliConfigChatDeleteConfirm = "cli.config.chat.delete_confirm"
	CliConfigChatDeleteSuccess = "cli.config.chat.delete_success"
)

// cli.config: config LLM
const (
	CliConfigLLMSelectPrompt  = "cli.config.llm.select_prompt"
	CliConfigLLMSelectHelp    = "cli.config.llm.select_help"
	CliConfigLLMDeleteConfirm = "cli.config.llm.delete_confirm"
	CliConfigLLMDeleteSuccess = "cli.config.llm.delete_success"
)

// cli.config: config tool
const (
	CliConfigToolSelectPrompt  = "cli.config.tool.select_prompt"
	CliConfigToolSelectHelp    = "cli.config.tool.select_help"
	CliConfigToolDeleteConfirm = "cli.config.tool.delete_confirm"
	CliConfigToolDeleteSuccess = "cli.config.tool.delete_success"
	CliConfigToolNoConfig      = "cli.config.tool.no_config"
)

// cli.config: config hachimi
const (
	CliConfigHachimiSelectPrompt  = "cli.config.config.hachimi.select_prompt"
	CliConfigHachimiSelectHelp    = "cli.config.config.hachimi.select_help"
	CliConfigHachimiSetDefault    = "cli.config.config.hachimi.set_default"
	CliConfigHachimiDeleteConfirm = "cli.config.config.hachimi.delete_confirm"
	CliConfigHachimiDeleteSuccess = "cli.config.config.hachimi.delete_success"
)

// cli.config: database landing
const (
	CliConfigDatabaseSelectPrompt = "cli.config.database.select_prompt"
	CliConfigDatabaseSelectHelp   = "cli.config.database.select_help"
	CliConfigDatabaseSessions     = "cli.config.database.sessions"
	CliConfigDatabaseUsers        = "cli.config.database.users"
	CliConfigDatabaseSouls        = "cli.config.database.souls"
	CliConfigDatabaseQuit         = "cli.config.database.quit"
)

// cli.config: database soul
const (
	CliConfigSoulSelectPrompt    = "cli.config.soul.select_prompt"
	CliConfigSoulEditAction      = "cli.config.soul.edit_action"
	CliConfigSoulDeleteAction    = "cli.config.soul.delete_action"
	CliConfigSoulEditContent     = "cli.config.soul.edit_content"
	CliConfigSoulEditContentHelp = "cli.config.soul.edit_content_help"
	CliConfigSoulUpdatedSuccess  = "cli.config.soul.updated_success"
	CliConfigSoulDeleteConfirm   = "cli.config.soul.delete_confirm"
	CliConfigSoulDeletedSuccess  = "cli.config.soul.deleted_success"
	CliConfigSoulCreatedSuccess  = "cli.config.soul.created_success"
	CliConfigSoulNameInput       = "cli.config.soul.name_input"
	CliConfigSoulNameHelp        = "cli.config.soul.name_help"
	CliConfigSoulContentInput    = "cli.config.soul.content_input"
	CliConfigSoulContentHelp     = "cli.config.soul.content_help"
	CliConfigSoulScopeInput      = "cli.config.soul.scope_input"
	CliConfigSoulScopeHelp       = "cli.config.soul.scope_help"
)

// cli.config: config run (command_run)
const (
	CliConfigRunDaemonRunning   = "cli.config.run.daemon_running"
	CliConfigRunInitError       = "cli.config.run.init_error"
	CliConfigRunError           = "cli.config.run.error"
	CliConfigEditNotFound       = "cli.config.run.edit_not_found"
	CliConfigEditCreatePrompt   = "cli.config.run.edit_create_prompt"
	CliConfigEditCancelled      = "cli.config.run.edit_cancelled"
	CliConfigEditOpenError      = "cli.config.run.edit_open_error"
	CliConfigOpenNotFound       = "cli.config.run.open_not_found"
	CliConfigOpenError          = "cli.config.run.open_error"
	CliConfigOnboardError       = "cli.config.run.onboard_error"
	CliConfigOnboardValidateErr = "cli.config.run.onboard_validate_error"
	CliConfigOnboardWriteError  = "cli.config.run.onboard_write_error"
	CliConfigOnboardSuccess     = "cli.config.run.onboard_success"
)

// cli.config: shared option labels
const (
	CliConfigOptionAddNew    = "cli.config.option.add_new"
	CliConfigOptionQuit      = "cli.config.option.quit"
	CliConfigOptionPurge     = "cli.config.option.purge"
	CliConfigOptionCreateNew = "cli.config.option.create_new"
	CliConfigOptionBye       = "cli.config.option.bye"
)

// cli.config: database user
const (
	CliConfigUserSelectPrompt   = "cli.config.user.select_prompt"
	CliConfigUserPromoteAction  = "cli.config.user.promote_action"
	CliConfigUserDegradeAction  = "cli.config.user.degrade_action"
	CliConfigUserDeleteAction   = "cli.config.user.delete_action"
	CliConfigUserNoUsers        = "cli.config.user.no_users"
	CliConfigUserPromoteConfirm = "cli.config.user.promote_confirm"
	CliConfigUserPromoteSuccess = "cli.config.user.promote_success"
	CliConfigUserDegradeConfirm = "cli.config.user.degrade_confirm"
	CliConfigUserDegradeSuccess = "cli.config.user.degrade_success"
	CliConfigUserDeleteConfirm  = "cli.config.user.delete_confirm"
	CliConfigUserDeleteSuccess  = "cli.config.user.delete_success"
)

// cli.config: database session
const (
	CliConfigSessionSelectPrompt    = "cli.config.session.select_prompt"
	CliConfigSessionEditAction      = "cli.config.session.edit_action"
	CliConfigSessionDeleteAction    = "cli.config.session.delete_action"
	CliConfigSessionSelectProvider  = "cli.config.session.select_provider"
	CliConfigSessionSelectModel     = "cli.config.session.select_model"
	CliConfigSessionUpdateSuccess   = "cli.config.session.update_success"
	CliConfigSessionDeleteConfirm   = "cli.config.session.delete_confirm"
	CliConfigSessionDeleteBounds    = "cli.config.session.delete_bounds"
	CliConfigSessionDeleteSuccess   = "cli.config.session.delete_success"
	CliConfigSessionPurgeConfirm    = "cli.config.session.purge_confirm"
	CliConfigSessionPurgeSuccess    = "cli.config.session.purge_success"
	CliConfigSessionPurgeError      = "cli.config.session.purge_error"
	CliConfigSessionPurgeClean      = "cli.config.session.purge_clean"
	CliConfigSessionChatDeleteError = "cli.config.session.chat_delete_error"
	CliConfigSessionDataDeleteError = "cli.config.session.data_delete_error"
)
