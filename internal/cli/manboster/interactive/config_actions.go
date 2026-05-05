package interactive

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/cli/helper"
	"github.com/manboster/manboster/internal/repository/types"
)

func (s *databaseConfigService) printConfigDatabaseSessionList(ctx context.Context) error {
	// string is sessionID
	chatsMap := make(map[string][]types.Chat)

	sessions, err := s.repo.GetSessions(ctx)
	if err != nil {
		return err
	}
	s.sessions = sessions

	chats, err := s.repo.GetAllChats(ctx)
	if err != nil {
		return err
	}
	s.chats = chats

	for _, c := range chats {
		cm, avail := chatsMap[c.SessionID]
		if !avail {
			cm = []types.Chat{c}
			chatsMap[c.SessionID] = cm
			continue
		}
		cm = append(cm, c)
		chatsMap[c.SessionID] = cm
	}
	s.chatsMap = chatsMap

	var outputMsg strings.Builder
	for _, sess := range s.sessions {
		outputMsg.WriteString(fmt.Sprintf("%d) `%s`, used `%s:%s` created at %s, updated at %s.\n", sess.ID, sess.SessionID, sess.LLMProvider, sess.LLMProviderModel, sess.CreatedAt.Format("2006-01-02 15:04:05"), sess.UpdatedAt.Format("2006-01-02 15:04:05")))
		cm, avail := chatsMap[sess.SessionID]
		if avail {
			outputMsg.WriteString(fmt.Sprintf("Bind %d chats: ", len(cm)))
			for _, c := range cm {
				outputMsg.WriteString(fmt.Sprintf("%s:%s ", c.ChatProvider, c.ChatID))
			}
			outputMsg.WriteString("\n")
		}
	}

	outputMsg.WriteString(fmt.Sprintf("%d sessions loaded, %d sessions can be purged.", len(s.sessions), len(sessions)-len(chatsMap)))

	helper.DisplayText(outputMsg.String())
	return nil
}

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
