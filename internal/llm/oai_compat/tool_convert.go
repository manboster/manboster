package oai_compat

import (
	"strings"

	"github.com/manboster/manboster/internal/tool"
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
				Parameters:  t.Args().ToJSONSchema(),
			},
		}
		oaiTools = append(oaiTools, oaiTool)
	}

	return oaiTools
}
