package keys

// Wizard: entrypoint
const (
	EntrypointSelectPrompt = "wizard.entrypoint.select_prompt"
	EntrypointSelectHelp   = "wizard.entrypoint.select_help"
	EntrypointDatabase     = "wizard.entrypoint.database"
	EntrypointConfig       = "wizard.entrypoint.config"
	EntrypointEditor       = "wizard.entrypoint.editor"
	EntrypointQuit         = "wizard.entrypoint.quit"
)

// Wizard: config landing
const (
	ConfigLandingSelectPrompt = "wizard.config_landing.select_prompt"
	ConfigLandingSelectHelp   = "wizard.config_landing.select_help"
	ConfigLandingChat         = "wizard.config_landing.chat"
	ConfigLandingLLM          = "wizard.config_landing.llm"
	ConfigLandingTool         = "wizard.config_landing.tool"
	ConfigLandingHachimi      = "wizard.config_landing.hachimi"
	ConfigLandingApp          = "wizard.config_landing.app"
	ConfigLandingQuit         = "wizard.config_landing.quit"
)

// Wizard: provider actions (shared)
const (
	ActionDeleteProvider = "wizard.action.delete_provider"
	ActionEditProvider   = "wizard.action.edit_provider"
	ActionSetDefault     = "wizard.action.set_default"
	ActionQuit           = "wizard.action.quit"
	ActionWhatToDo       = "wizard.action.what_to_do"
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
	ConfigLLMSelectPrompt  = "wizard.config.llm_select_prompt"
	ConfigLLMSelectHelp    = "wizard.config.llm_select_help"
	ConfigLLMDeleteConfirm = "wizard.config.llm_delete_confirm"
	ConfigLLMDeleteSuccess = "wizard.config.llm_delete_success"
)

// Wizard: config tool
const (
	ConfigToolSelectPrompt  = "wizard.config.tool_select_prompt"
	ConfigToolSelectHelp    = "wizard.config.tool_select_help"
	ConfigToolDeleteConfirm = "wizard.config.tool_delete_confirm"
	ConfigToolDeleteSuccess = "wizard.config.tool_delete_success"
	ConfigToolNoConfig      = "wizard.config.tool_no_config"
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
	DatabaseSelectPrompt = "wizard.database.select_prompt"
	DatabaseSelectHelp   = "wizard.database.select_help"
	DatabaseSessions     = "wizard.database.sessions"
	DatabaseUsers        = "wizard.database.users"
	DatabaseSouls        = "wizard.database.souls"
	DatabaseQuit         = "wizard.database.quit"
)

// Wizard: database soul
const (
	SoulSelectPrompt    = "wizard.soul.select_prompt"
	SoulEditAction      = "wizard.soul.edit_action"
	SoulDeleteAction    = "wizard.soul.delete_action"
	SoulEditContent     = "wizard.soul.edit_content"
	SoulEditContentHelp = "wizard.soul.edit_content_help"
	SoulUpdatedSuccess  = "wizard.soul.updated_success"
	SoulDeleteConfirm   = "wizard.soul.delete_confirm"
	SoulDeletedSuccess  = "wizard.soul.deleted_success"
	SoulCreatedSuccess  = "wizard.soul.created_success"
	SoulNameInput       = "wizard.soul.name_input"
	SoulNameHelp        = "wizard.soul.name_help"
	SoulContentInput    = "wizard.soul.content_input"
	SoulContentHelp     = "wizard.soul.content_help"
	SoulScopeInput      = "wizard.soul.scope_input"
	SoulScopeHelp       = "wizard.soul.scope_help"
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
