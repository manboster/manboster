package interactive

import (
	"context"
	"fmt"

	"github.com/fatih/color"
)

func (s *databaseConfigService) purgeConfigDatabaseSession(ctx context.Context) error {
	for _, sess := range s.sessions {
		cm, avail := s.chatsMap[sess.SessionID]
		if !avail {
			err := s.repo.DeleteSession(ctx, sess.SessionID)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Client] Error purging session %s: %q", sess.SessionID, err))
				continue
			}
			err = s.repo.DeleteChatData(ctx, sess.SessionID)
			if err != nil {
				color.Yellow(fmt.Sprintf("[Manboster Client] Error purging session %s: %q", sess.SessionID, err))
				continue
			}
			for _, c := range cm {
				err := s.repo.DeleteChat(ctx, c.ChatID, c.ChatProvider)
				if err != nil {
					color.Yellow(fmt.Sprintf("[Manboster Client] Error purging session %s: %q", sess.SessionID, err))
					continue
				}
			}
		}
	}
	return nil
}

func (s *databaseConfigService) editConfigSessionDatabase(ctx context.Context, sid string, provider string, model string) error {
	return s.repo.UpdateSession(ctx, sid, map[string]interface{}{
		"llm_provider":       provider,
		"llm_provider_model": model,
	})
}

func (s *databaseConfigService) deleteConfigSessionDatabase(ctx context.Context, sid string) error {
	err := s.repo.DeleteSession(ctx, sid)
	if err != nil {
		return err
	}
	cm, avail := s.chatsMap[sid]
	if !avail {
		return nil
	}
	for _, c := range cm {
		err := s.repo.DeleteChat(ctx, c.ChatID, c.ChatProvider)
		if err != nil {
			color.Yellow(fmt.Sprintf("[Manboster Client] Error deleting session %s: %q", sid, err))
		}
	}
	return nil
}
