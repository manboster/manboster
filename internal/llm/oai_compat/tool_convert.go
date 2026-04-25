package oai_compat

import (
	"strings"

	"github.com/manboster/manboster/internal/tool"
	"github.com/manboster/manboster/spec/schema"
	"github.com/sashabaranov/go-openai"
)

// ConvertTools Converts tools provider to openai tools compatible.
func (s *Service) ConvertTools(tools []tool.Provider) []openai.Tool {
	var oaiTools = make([]openai.Tool, 0, len(tools))

	for _, t := range tools {
		safeName := strings.ReplaceAll(t.Name(), ".", "_")
		oaiTool := openai.Tool{
			Type: openai.ToolTypeFunction,
			Function: &openai.FunctionDefinition{
				Name:        safeName,
				Description: t.MetaData().Description,
				Strict:      false,
				Parameters:  buildParameters(t.Args()),
			},
		}
		oaiTools = append(oaiTools, oaiTool)
	}

	return oaiTools
}

func buildParameters(args []*schema.Args) map[string]any {
	properties := make(map[string]any)
	var required []string

	for _, arg := range args {
		paramDef := map[string]any{
			"type":        arg.Type,
			"description": arg.Description,
		}

		if arg.IsEnum && len(arg.Enum) > 0 {
			paramDef["enum"] = arg.Enum
		}

		properties[arg.Name] = paramDef

		if arg.Required {
			required = append(required, arg.Name)
		}
	}

	res := map[string]any{
		"type":       "object",
		"properties": properties,
	}
	if len(required) > 0 {
		res["required"] = required
	}
	return res
}
