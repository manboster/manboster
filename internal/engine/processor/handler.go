package processor

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/session"
	"github.com/manboster/manboster/spec/chat"
)

func (s *Service) Process(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	var resultProcess ProcessSuggestion
	color.Cyan("[Manboster Processor] Processing message")

	// before receiving messages, we should check users' identity.
	// get user information
	resultProcess = ProcessDrop
	uType := s.safeguardService.UserType(ctx, instance.Name(), msg.UserID)
	// allowed := chat.MessageSelectionCallback | chat.MessageSelection | chat.MessageCommand | chat.MessageFromCron | chat.MessageFromCronIgnore

	// first we check personal chats
	if msg.ChatType == chat.ChatsPersonal {
		resultProcess = ProcessHandle

		// checkout onboard message available or not, MessageCommand is passthrough
		if s.onboard != nil && s.onboard.Active() && msg.MessageType&chat.MessageCommand == 0 {
			msg.MessageType = chat.MessageStart
		}

		// not an admin && not start message, we flag it to reject
		if !s.safeguardService.IsAdmin(uType) && msg.MessageType&(chat.MessageStart|chat.MessageCommand) == 0 {
			color.Yellow(fmt.Sprintf("[Manboster Processor] We detected an unknown user wants to talk with your lobster in person!"))
			msg.MessageType = chat.MessageUnknown
		}
	} else {
		// then we check other chat information, automatically reject
		resultProcess = ProcessDrop
		if msg.MessageType&chat.MessageText != 0 {
			if msg.Text != nil && strings.Contains(msg.Text.Text, "[[!@{Assistant}]]") {
				resultProcess = ProcessHandle
			}
			if msg.Reply != nil && msg.Reply.Username == "Assistant" {
				resultProcess = ProcessHandle
			}
		}
	}

	switch resultProcess {
	case ProcessHandle:
		// get message types
		sessionId, err := s.sessionService.LoadChatSession(ctx, instance, msg, s.safeguardService.IsAdmin(uType))
		// if you're not an administrator, you can not create a new session
		if errors.Is(err, session.ErrAccessDenied) {
			color.Yellow(fmt.Sprintf("[Manboster Processor] We detected an unknown user wants to start a new chat!"))
			msg.MessageType = chat.MessageUnknown
		} else if err != nil {
			color.Red(fmt.Sprintf("[Manboster Processor] We encountered an error while processing message: %q", err))
			return err
		}
		return s.engine.Distribute(ctx, instance, msg, sessionId)
	//case ProcessDrop:
	//	return nil
	//case ProcessConsider:
	//	return nil // TODO: manboster active mode
	default:
		return nil
	}
}
