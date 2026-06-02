package soul

import (
	"context"
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/internal/util"
	"github.com/manboster/manboster/spec/chat"
	"github.com/manboster/manboster/spec/llm"
)

// BuildLLMMessage build from a chat message to a llm message, make it easier to handle in engine
func (s *Service) BuildLLMMessage(ctx context.Context, msg *chat.Message, sessionId string, userType types.UserType) (*llm.Message, error) {
	var respString strings.Builder
	// who said...
	chatName := "(Private Chat)"
	if msg.ChatType != chat.ChatsPersonal {
		chatName = msg.ChatName
	}

	// generate nonce
	nonceMetadata := util.RandomString(8)
	nonceInput := util.RandomString(8)

	forwardPrompt := ""
	if msg.Forward != nil {
		if msg.Forward.ChatName == "" {
			msg.Forward.ChatName = "(unknown chat)"
		}
		forwardPrompt = fmt.Sprintf("Forwarded from %s", msg.Forward.ChatName)
		if msg.Forward.Username != "" {
			forwardPrompt += fmt.Sprintf("(%s said", msg.Forward.Username)
		}
		if msg.Forward.UserID != "" {
			forwardPrompt += fmt.Sprintf(", ID: %s", msg.Forward.UserID)
		}
		forwardPrompt += ")\n"
	}

	// append prompt
	respString.WriteString(fmt.Sprintf("[chat metadata %s]\n%s%s(UID:%s, Role:%s) said in %s, [%s]:\n", nonceMetadata, forwardPrompt, msg.Username, msg.UserID, userType, msg.CreatedAt.Format("2006-01-02 15:04:05T-07"), chatName))

	if msg.Reply != nil {
		replyMsg := msg.Reply
		respString.WriteString(fmt.Sprintf("Replied: %s(UID:%s) said in %s, [%s]:%s\n", replyMsg.Username, replyMsg.UserID, replyMsg.CreatedAt.Format("2006-01-02 15:04:05T-07"), chatName, forwardPrompt))
		str := s.ChatMessageToString(replyMsg)
		respString.WriteString(fmt.Sprintf("[reply user input %s]\n%s\n", nonceInput, str))
	}

	// check msg data
	str := s.ChatMessageToString(msg)
	respString.WriteString(fmt.Sprintf("[user input %s]\n%s\n", nonceInput, str))

	respMsg := &llm.Message{}
	respMsg.Role = llm.RoleUser
	respMsg.Type = llm.MessageText

	if msg.MessageType&chat.MessageText != 0 {
		respMsg.Parts = []llm.MessageParts{
			{
				PartsType: llm.MessagePartsText,
				Text: &llm.MessageTextPayload{
					Text: respString.String(),
				},
			},
		}
	}

	if msg.MessageType&chat.MessageImage != 0 && msg.Image != nil {
		for _, c := range msg.Image.Content {
			respMsg.Parts = append(respMsg.Parts, llm.MessageParts{
				PartsType: llm.MessagePartsImage,
				Image: &llm.MessageImagePayload{
					Content: c,
				},
			})
		}
	}

	return respMsg, nil
}
