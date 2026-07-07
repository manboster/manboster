package chat

import (
	"context"

	"github.com/manboster/manboster/spec/chat"
)

func (s *Service) DeleteChatSession(ctx context.Context, instance chat.Provider, msg *chat.Message, sessionId string, deleteData bool) error {
	err := s.repo.DeleteChat(ctx, msg.ChatID, instance.Name())
	if err != nil {
		return err
	}

	if deleteData {
		err = s.repo.DeleteChatData(ctx, sessionId)
		if err != nil {
			return err
		}
	}

	err = s.repo.DeleteSession(ctx, sessionId)
	if err != nil {
		return err
	}

	return nil
}
