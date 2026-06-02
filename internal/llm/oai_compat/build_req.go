package oai_compat

import (
	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/llm"
	"github.com/sashabaranov/go-openai"
)

func (s *Service) buildRequest(msg llm.Message, model llm.Model) []openai.ChatCompletionMessage {
	if msg.Type&(llm.MessageToolCallResponse) != 0 {
		// going to check tool call resp and get it back to llm!
		for _, resp := range msg.ToolCallResponse {
			if resp.Result == "" {
				resp.Result = "{}"
			}

			var res []openai.ChatCompletionMessage
			res = append(res, openai.ChatCompletionMessage{
				Role:       openai.ChatMessageRoleTool,
				Content:    resp.Result,
				ToolCallID: resp.ID,
				Name:       resp.ToolName,
			})
			return res
		}
	}

	ccMsg := openai.ChatCompletionMessage{
		Role: string(msg.Role),
	}

	if msg.Role == llm.RoleSystem {
		ccMsg.Content = msg.Parts[0].Text.Text
		return []openai.ChatCompletionMessage{ccMsg}
	}

	// get there is a request or not
	if msg.Type&(llm.MessageToolCallRequest) != 0 {
		if model.Capabilities.Input&llm.CapabilityToolCall == 0 {
			color.Yellow("unsupported tool call!")
		} else {
			for _, req := range msg.ToolCallRequest {
				ccMsg.ToolCalls = append(ccMsg.ToolCalls, openai.ToolCall{
					ID:   req.ID,
					Type: openai.ToolTypeFunction,
					Function: openai.FunctionCall{
						Name:      req.ToolName,
						Arguments: req.ToolArgs.(string),
					},
				})
			}
		}
	}

	// get there is a reasoning or not
	if msg.Type&(llm.MessageThinking) != 0 && msg.Thinking != nil {
		ccMsg.ReasoningContent = msg.Thinking.Thinking
	}

	if msg.Type&(llm.MessageText|llm.MessageFile|llm.MessageImage) != 0 {
		for _, part := range msg.Parts {
			switch part.PartsType {
			case llm.MessagePartsText:
				ccMsg.MultiContent = append(ccMsg.MultiContent, openai.ChatMessagePart{
					Type: openai.ChatMessagePartTypeText,
					Text: part.Text.Text,
				})
			case llm.MessagePartsImage:
				if model.Capabilities.Input&llm.CapabilityImage == 0 {
					color.Yellow("[Manboster LLM Provider] unsupported model")
					continue
				}

				ccMsg.MultiContent = append(ccMsg.MultiContent, openai.ChatMessagePart{
					Type: openai.ChatMessagePartTypeImageURL,
					ImageURL: &openai.ChatMessageImageURL{
						URL: part.Image.Content,
					},
				})
			default:
			}
		}
	}
	return []openai.ChatCompletionMessage{ccMsg}
}
