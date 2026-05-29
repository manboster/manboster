package gatekeeper

import "fmt"

func BuildSessionId(name string, chatId string, sid string) string {
	return fmt.Sprintf("%s:%s:%s", name, chatId, sid)
}

func buildToolId(sessionId string, toolName string, executeGroup string) string {
	return fmt.Sprintf("%s:%s:%s", sessionId, toolName, executeGroup)
}
