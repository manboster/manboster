package engine

import (
	"context"
	"fmt"
	"strings"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/config"
)

// cmdId displays id information of current chat
func (e *Engine) cmdId(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	respText := strings.Builder{}
	respText.WriteString(fmt.Sprintf("Message ID: %s\n", msg.MessageID))
	respText.WriteString(fmt.Sprintf("Message User ID: %s\n", msg.UserID))
	respText.WriteString(fmt.Sprintf("Message Chat ID: %s\n", msg.ChatID))
	respText.WriteString(fmt.Sprintf("Chat Provider: %s\n", instance.Name()))
	if msg.Reply != nil {
		respText.WriteString(fmt.Sprintf("Message Replying ID: %s\n", msg.Reply.MessageID))
		respText.WriteString(fmt.Sprintf("Message Replying Chat ID: %s\n", msg.Reply.ChatID))
		respText.WriteString(fmt.Sprintf("Message Replying User ID: %s\n", msg.Reply.UserID))
	}
	msg.Text = &chat.TextPayload{
		Text: respText.String(),
	}

	return instance.SendMessage(ctx, msg)
}

// cmdHelp is the helper command of manboster
func (e *Engine) cmdHelp(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
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

// cmdVersion when user execute version commands, it will run.
func (e *Engine) cmdVersion(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{
		Text: fmt.Sprintf("Manboster: Your Personal Manbo Lobster!\nManboster version `%s %s@%s`, build at `%s`\nCheckout our latest releases here:\nhttps://github.com/manboster/manboster", config.Version, config.CurrentVersion, config.BuildCommit, config.BuildTime),
	}
	return instance.SendMessage(ctx, msg)
}

func (e *Engine) cmdDefault(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	msg.Text = &chat.TextPayload{
		Text: "We are sorry but this is an Invalid Command. Please check your grammatical mistakes.",
	}
	return instance.SendMessage(ctx, msg)
}

func (e *Engine) cmdStart(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	msg.MessageType = chat.MessageText
	var txt strings.Builder
	txt.WriteString("Welcome to use Manboster!\n")
	txt.WriteString("If this is your first use, please send something and trigger pair process.\n")
	txt.WriteString("If this Lobster is not yours, please contact owner to get access.\n")
	txt.WriteString("You can also use the following commands:\n")
	txt.WriteString("/help - show the whole help command\n")
	txt.WriteString("/id - get current information\n")
	txt.WriteString("/cancel - cancel the request.\n")
	txt.WriteString("/status - get current status of this conversation.\n")
	msg.Text = &chat.TextPayload{
		Text: txt.String(),
	}
	return instance.SendMessage(ctx, msg)
}
