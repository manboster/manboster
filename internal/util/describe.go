package util

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/llm"
)

var bracketRegex = regexp.MustCompile(`\[.*?]\s*`)

func DescribeToHachimi(req llm.MessageToolCallRequestPayload, provider tool.Provider) string {
	var descStr strings.Builder

	metadata := provider.MetaData()
	purgedDescription := bracketRegex.ReplaceAllString(metadata.Description, "")
	descStr.WriteString(fmt.Sprintf("Model wants to call %s, description: %s, required min user type: %s, ", provider.Name(), purgedDescription, metadata.MinUserType))

	var args map[string]interface{}
	err := json.Unmarshal([]byte(fmt.Sprintf("%v", req.ToolArgs)), &args)
	if err != nil {
		color.Yellow("[Manboster Handler] Failed to unmarshal tool call result")
		return ""
	}

	schema := provider.Args()
	if schema != nil && len(schema.Properties) > 0 {
		descStr.WriteString("with the following arguments:\n")
		for _, prop := range schema.Properties {
			val := args[prop.Name]
			descStr.WriteString(fmt.Sprintf("  - %s", prop.Name))
			if prop.Description != "" {
				descStr.WriteString(fmt.Sprintf(" (%s)", prop.Description))
			}
			descStr.WriteString(fmt.Sprintf(": %v", val))
			if prop.IsEnum && len(prop.Enum) > 0 {
				descStr.WriteString(fmt.Sprintf(" (options: %v)", prop.Enum))
			}
			descStr.WriteString("\n")
		}
	} else {
		descStr.WriteString(fmt.Sprintf("\n%s", JSONParseFull(args)))
	}

	return descStr.String()
}

func DescribeToHuman(req llm.MessageToolCallRequestPayload, provider tool.Provider) string {
	txt := fmt.Sprintf("Model wants to call tool `%s`(`%s`) ", provider.DisplayName(), req.ToolName)
	var result map[string]interface{}
	err := json.Unmarshal([]byte(fmt.Sprintf("%v", req.ToolArgs)), &result)
	if err != nil {
		color.Yellow("[Manboster Handler] Failed to unmarshal tool call result")
	}
	params := JSONParse(result)
	if params != "" {
		txt += fmt.Sprintf("with params: %s", params)
	}
	txt += ", do you want to continue?"
	return txt
}
