package interactive

type configSelection string

const (
	configSelectionDatabase configSelection = "database"
	configSelectionQuit     configSelection = "quit"
	configSelectionConfig   configSelection = "config"
	configSelectionEditor   configSelection = "editor"
)

type databaseConfigSessionPageSelection string

const (
	databaseConfigSessionPageQuit   databaseConfigSessionPageSelection = "quit"
	databaseConfigSessionPageEdit   databaseConfigSessionPageSelection = "edit"
	databaseConfigSessionPageDelete databaseConfigSessionPageSelection = "delete"
)

type databaseConfigLandingSelection string

const (
	databaseConfigLandingUser    databaseConfigLandingSelection = "user"
	databaseConfigLandingSession databaseConfigLandingSelection = "session"
	databaseConfigLandingSoul    databaseConfigLandingSelection = "soul"
	databaseConfigLandingQuit    databaseConfigLandingSelection = "quit"
)

type databaseConfigSessionSelection string

const (
	databaseConfigSessionPurge  databaseConfigSessionSelection = "purge"
	databaseConfigSessionSelect databaseConfigSessionSelection = "select"
	databaseConfigSessionQuit   databaseConfigSessionSelection = "quit"
)

type configLandingSelection string

const (
	configLandingChat    configLandingSelection = "chat"
	configLandingLLM     configLandingSelection = "llm"
	configLandingTool    configLandingSelection = "tool"
	configLandingHachimi configLandingSelection = "hachimi"
	configLandingApp     configLandingSelection = "app"
	configLandingQuit    configLandingSelection = "quit"
)

type configLandingChatActionSelection string

const (
	configLandingChatActionNew  configLandingChatActionSelection = "new"
	configLandingChatActionQuit configLandingChatActionSelection = "quit"
)
