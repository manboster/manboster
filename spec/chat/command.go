package chat

// CommandType defines command's types.
type CommandType string

const (
	CommandUnknown   CommandType = ""          // No Command Available
	CommandVersion   CommandType = "version"   // Get Version Data
	CommandHelp      CommandType = "help"      // Get Helper Messages
	CommandOp        CommandType = "op"        // Grant a user
	CommandDeOp      CommandType = "deop"      // Ungrant a user
	CommandId        CommandType = "id"        // display ids
	CommandStatus    CommandType = "status"    // display current status
	CommandSave      CommandType = "save"      // save this chat to database
	CommandNew       CommandType = "new"       // delete this and create a new chat
	CommandCompact   CommandType = "compact"   // summary this chat and create a new chat with summarized items
	CommandModel     CommandType = "model"     // select models you want
	CommandModels    CommandType = "models"    // select models using interactive selection field, when AbilityType & AbilitySendSelect != 0
	CommandSession   CommandType = "session"   // select sessions you want
	CommandSessions  CommandType = "sessions"  // select sessions using interactive selection field, when AbilityType & AbilitySendSelect != 0
	CommandProvider  CommandType = "provider"  // select providers you want
	CommandProviders CommandType = "providers" // select providers using interactive selection field, when AbilityType & AbilitySendSelect != 0
	CommandStart     CommandType = "start"     // start command gives tips to you, when it's the first run, it will grant you the root access to this application.
	CommandPair      CommandType = "pair"      // pair Manboster with pair code
	CommandCancel    CommandType = "cancel"    // cancel this request
)

func IsPublicCommand(commandType CommandType) bool {
	return commandType == CommandId || commandType == CommandPair || commandType == CommandVersion
}

// IsSessionRequiredCommand returns whether the command need a Session ID. CommandCancel is included because it needs a sessionId lookup
func IsSessionRequiredCommand(commandType CommandType) bool {
	return commandType == CommandCompact || commandType == CommandCancel || commandType == CommandNew || commandType == CommandSave || commandType == CommandStatus || commandType == CommandProvider || commandType == CommandProviders || commandType == CommandModels || commandType == CommandModel || commandType == CommandSessions || commandType == CommandSession
}
