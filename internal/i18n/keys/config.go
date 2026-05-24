package keys

// Config args descriptions: LLM providers
const (
	LLMOpenRouterAPIKeyDesc     = "config.llm.openrouter.api_key_desc"
	LLMOAICompatNameDesc        = "config.llm.oai_compat.name_desc"
	LLMOAICompatDisplayNameDesc = "config.llm.oai_compat.display_name_desc"
	LLMOAICompatBaseURLDesc     = "config.llm.oai_compat.base_url_desc"
	LLMOAICompatAPIKeyDesc      = "config.llm.oai_compat.api_key_desc"
)

// Config args descriptions: Chat providers
const (
	ChatTelegramBotTokenDesc          = "config.chat.telegram.bot_token_desc"
	ChatTelegramCollapseLengthDesc    = "config.chat.telegram.collapse_length_desc"
	ChatTelegramReactionStatusDesc    = "config.chat.telegram.reaction_status_desc"
)

// Config args descriptions: Hachimi
const (
	HachimiGGUFModelTypeDesc    = "config.hachimi.gguf.model_type_desc"
	HachimiGGUFContextLenDesc   = "config.hachimi.gguf.context_length_desc"
)

// Config args descriptions: Tools
const (
	ToolFileWriteModeDesc = "config.tool.file.write_mode_desc"
	ToolBrowserModeDesc   = "config.tool.browser.mode_desc"
)

// Shared confirm/deny button labels
const (
	BtnYes      = "btn.yes"
	BtnNo       = "btn.no"
	BtnSkip     = "btn.skip"
	BtnContinue = "btn.continue"
	BtnExit     = "btn.exit"
	BtnExitNow  = "btn.exit_now"
	BtnRetry    = "btn.retry"
)

// Shared question subtitles
const (
	QuestionWantToRetry   = "question.want_to_retry"
	QuestionWantToContinue = "question.want_to_continue"
)

// Hachimi GGUF setup prompts
const (
	HachimiGGUFSetupPrompt   = "config.hachimi.gguf.setup_prompt"
	HachimiGGUFSetupQuestion = "config.hachimi.gguf.setup_question"
	HachimiGGUFURLInput      = "config.hachimi.gguf.url_input"
	HachimiGGUFURLHelp       = "config.hachimi.gguf.url_help"
	HachimiGGUFSHA256Input   = "config.hachimi.gguf.sha256_input"
	HachimiGGUFSHA256Help    = "config.hachimi.gguf.sha256_help"
	HachimiGGUFSetupSuccess  = "config.hachimi.gguf.setup_success"
	HachimiGGUFAutoSetMsg    = "config.hachimi.gguf.auto_set_msg"
)

// OAI compat setup
const (
	OAICompatCredentialError    = "config.llm.oai_compat.credential_error"
	OAICompatCredentialErrorMsg = "config.llm.oai_compat.credential_error_msg"
	OAICompatModelSelectPrompt  = "config.llm.oai_compat.model_select_prompt"
	OAICompatModelSelectHelp    = "config.llm.oai_compat.model_select_help"
	OAICompatOtherModel         = "config.llm.oai_compat.other_model"
)
