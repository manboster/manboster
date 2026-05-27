package keys

// Wizard: onboard existing config
const (
	OnboardExistingConfig   = "cli.onboard.existing_config"
	OnboardExistingContinue = "cli.onboard.existing_continue"
	OnboardUserCancelled    = "cli.onboard.user_cancelled"
	OnboardWarningRejected  = "cli.onboard.warning_rejected"
)

// Wizard: onboard chat
const (
	OnboardChatSelectPrompt    = "cli.onboard.chat_select_prompt"
	OnboardChatAddedCount      = "cli.onboard.chat_added_count"
	OnboardChatNoMoreProviders = "cli.onboard.chat_no_more_providers"
	OnboardChatConfigError     = "cli.onboard.chat_config_error"
)

// Wizard: onboard LLM
const (
	OnboardLLMSelectPrompt = "cli.onboard.llm_select_prompt"
	OnboardLLMAddedCount   = "cli.onboard.llm_added_count"
	OnboardLLMConfigError  = "cli.onboard.llm_config_error"
)

// Wizard: onboard app
const (
	OnboardAppSelectProvider = "cli.onboard.app_select_provider"
	OnboardAppSelectHelp     = "cli.onboard.app_select_help"
	OnboardAppSelectModel    = "cli.onboard.app_select_model"
)

// Wizard: onboard hachimi
const (
	OnboardHachimiFeaturePrompt  = "cli.onboard.hachimi_feature_prompt"
	OnboardHachimiEnableQuestion = "cli.onboard.hachimi_enable_question"
	OnboardHachimiAddedCount     = "cli.onboard.hachimi_added_count"
	OnboardHachimiSelectDefault  = "cli.onboard.hachimi_select_default"
	OnboardHachimiSelectHelp     = "cli.onboard.hachimi_select_help"
	OnboardHachimiSelectProvider = "cli.onboard.hachimi_select_provider"
	OnboardHachimiNoMore         = "cli.onboard.hachimi_no_more"
	OnboardHachimiConfigError    = "cli.onboard.hachimi_config_error"
)

// Wizard: onboard tool
const (
	OnboardToolSelectPrompt = "cli.onboard.tool_select_prompt"
	OnboardToolSelectHelp   = "cli.onboard.tool_select_help"
)

// Wizard: onboard preview
const (
	OnboardPreviewTitle           = "cli.onboard.preview_title"
	OnboardPreviewRestart         = "cli.onboard.preview_restart"
	OnboardPreviewChatCount       = "cli.onboard.preview_chat_count"
	OnboardPreviewLLMCount        = "cli.onboard.preview_llm_count"
	OnboardPreviewToolCount       = "cli.onboard.preview_tool_count"
	OnboardPreviewHachimiEnabled  = "cli.onboard.preview_hachimi_enabled"
	OnboardPreviewHachimiDisabled = "cli.onboard.preview_hachimi_disabled"
	OnboardPreviewContinue        = "cli.onboard.preview_continue"
	OnboardPreviewConfirm         = "cli.onboard.preview_confirm"
	OnboardPreviewProblem         = "cli.onboard.preview_problem"
)

// Wizard: write config
const (
	OnboardWriteExisting   = "cli.onboard.write_existing"
	OnboardWriteConfirm    = "cli.onboard.write_confirm"
	OnboardWritePathPrompt = "cli.onboard.write_path_prompt"
	OnboardWritePathHelp   = "cli.onboard.write_path_help"
	OnboardWriteError      = "cli.onboard.write_error"
	OnboardWriteSuccess    = "cli.onboard.write_success"
)

// Wizard: onboard warning
const (
	OnboardWarningRiskTitle      = "cli.onboard.warning_risk_title"
	OnboardWarningRiskPrompt     = "cli.onboard.warning_risk_prompt"
	OnboardWarningUnstableTitle  = "cli.onboard.warning_unstable_title"
	OnboardWarningUnstablePrompt = "cli.onboard.warning_unstable_prompt"
	OnboardWarningAccept         = "cli.onboard.warning_accept"
	OnboardWarningExit           = "cli.onboard.warning_exit"
)
