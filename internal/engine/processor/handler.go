package processor

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/spec/chat"
)

func (s *Service) Process(ctx context.Context, instance chat.Provider, msg *chat.Message) error {
	var resultProcess ProcessSuggestion
	color.Cyan("[Manboster Processor] Processing message")

	// before receiving messages, we should check users' identity.
	// get user information
	uType := s.safeguardService.UserType(ctx, instance.Name(), msg.UserID)

	if !s.safeguardService.IsAdmin(uType) && msg.ChatType == chat.ChatsPersonal {
		color.Yellow(fmt.Sprintf("[Manboster Processor] We detected an unknown user wants to talk with your lobster in person!"))
		msg.MessageType = chat.MessageUnknown
		resultProcess = ProcessHandle
	}

	// get message types
	sessionId, err := s.sessionService.LoadChatSession(ctx, instance, msg, s.safeguardService.IsAdmin(uType))
	// if you're not an administrator, you can not create a new session
	if errors.Is(err, ErrAccessDenied) {
		color.Yellow(fmt.Sprintf("[Manboster Processor] We detected an unknown user wants to start a new chat!"))
		msg.MessageType = chat.MessageUnknown
		resultProcess = ProcessHandle
	}
	if err != nil {
		color.Red(fmt.Sprintf("[Manboster Processor] We encountered an error while loading sessionId, error: %q", err))
		return err
	}

	if s.onboard != nil && !s.onboard.Active() {
		msg.MessageType = chat.MessageStart
		resultProcess = ProcessHandle
	}

	if msg.ChatType == chat.ChatsGroup {
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
		return s.engine.Distribute(ctx, instance, msg, sessionId)
	case ProcessDrop:
		return nil
	case ProcessConsider:
		return nil // TODO: manboster active mode
	default:
		return nil
	}
}
