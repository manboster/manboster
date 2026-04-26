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

// buildParameters converts args to JSON Schema.
func buildParameters(root *schema.Args) map[string]any {
	properties := make(map[string]any)
	var required []string

	for _, prop := range root.Properties {
		paramDef := buildParamDef(prop)
		properties[prop.Name] = paramDef

		if prop.Required {
			required = append(required, prop.Name)
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

func buildParamDef(arg *schema.Args) map[string]any {
	def := map[string]any{
		"type":        arg.Type,
		"description": arg.Description,
	}

	if arg.IsEnum && len(arg.Enum) > 0 {
		def["enum"] = arg.Enum
	}

	if arg.Type == schema.ArgsTypeObject && len(arg.Properties) > 0 {
		subProps := make(map[string]any)
		var subRequired []string
		for _, p := range arg.Properties {
			subProps[p.Name] = buildParamDef(p)
			if p.Required {
				subRequired = append(subRequired, p.Name)
			}
		}
		def["properties"] = subProps
		if len(subRequired) > 0 {
			def["required"] = subRequired
		}
	}

	if arg.Type == schema.ArgsTypeArray && arg.Items != nil {
		def["items"] = buildParamDef(arg.Items)
	}

	return def
}
