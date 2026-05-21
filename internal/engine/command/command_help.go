package command

import (
	"context"
	"strings"

	"github.com/manboster/manboster/spec/chat"
)

// cmdHelp is the helper command of manboster
func (h *Handler) cmdHelp(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	var str strings.Builder
	str.WriteString("Manboster Help Commands:\n\n")
	str.WriteString("/version - Get current version information of this Manboster instance\n")
	str.WriteString("/id - Get current user and chat's id\n")
	str.WriteString("/help - Get helper information of Manboster\n")
	str.WriteString("/op - Grant a user operator permission, just reply who you want to grant, root permission required. You can also use it by `/op [userid]`.\n")
	str.WriteString("/deop - Ungrant a user operator permission, just reply who you want to ungrant, root permission required. You can also use it by `/deop [userid]`.\n")
	str.WriteString("/status - Get current and overall status of session, chats and context.\n")
	str.WriteString("/save - Save the current session and start a new session. You can change it anytime by using `/session`.\n")
	str.WriteString("/new - Delete the current session data and start a new session.\n")
	str.WriteString("/compact - Compact the current session by summarizing context and start a new session. If the context is about to overflow, it will be done automatically.\n")
	str.WriteString("/model - Get and select current chat's chatting model by `/model [model id]`.\n")
	str.WriteString("/models - Select current chat's chatting model in an interactive way.\n")
	str.WriteString("/session - Get and select current chat's session by `/session [session id]`\n")
	str.WriteString("/sessions - Select current chat's session in an interactive way\n")
	str.WriteString("/provider - Get and select current chat's provider by `/provider [provider id]`\n")
	str.WriteString("/providers - Select current chat's provider in an interactive way\n")
	str.WriteString("/start - Display the start welcome message\n")
	str.WriteString("/pair - Use /pair xxxxxx to pair with your Lobster\n")
	str.WriteString("/cancel - Cancel current chat's pending request")

	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{
		Text: str.String(),
	}
	return instance.SendMessage(ctx, msg)
}
