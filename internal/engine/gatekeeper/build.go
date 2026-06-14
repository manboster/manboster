package gatekeeper

import (
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/config/prompt"
	"github.com/manboster/manboster/internal/hachimi"
)

func BuildSessionId(name string, chatId string, sid string) string {
	return fmt.Sprintf("%s:%s:%s", name, chatId, sid)
}

func buildToolId(sessionId string, toolName string, executeGroup string) string {
	return fmt.Sprintf("%s:%s:%s", sessionId, toolName, executeGroup)
}

func buildTemplate(status hachimi.ResponseStatusType) string {
	return strings.Replace(prompt.DescribeSafetyPrompt, "{{verdict}}", fmt.Sprintf("%s", status), -1)
}
