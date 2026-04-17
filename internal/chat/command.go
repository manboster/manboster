package chat

// CommandPayload defines what's in commands
type CommandPayload struct {
	CommandType CommandType // Optional. Required when MessageType = MessageCommand Command's type
	CommandArgs []string    // Optional. Required when MessageType = MessageCommand Command's args
}

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
	CommandModels    CommandType = "models"    // select models you want
	CommandProviders CommandType = "providers" // select providers you want
	CommandStart     CommandType = "start"     // start command gives tips to you, when it's the first run, it will grant you the root access to this application.
	CommandPair      CommandType = "pair"      // pair Manboster with pair code
	CommandCancel    CommandType = "cancel"    // cancel this request
)

func IsPublicCommand(commandType CommandType) bool {
	return commandType == CommandId || commandType == CommandPair || commandType == CommandVersion
}

// IsSessionRequiredCommand returns whether the command need a Session ID. CommandCancel is included because it needs a sessionId lookup
func IsSessionRequiredCommand(commandType CommandType) bool {
	return commandType == CommandCompact || commandType == CommandCancel || commandType == CommandNew || commandType == CommandSave || commandType == CommandStatus
}
