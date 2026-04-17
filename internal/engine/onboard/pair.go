package onboard

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/chat"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
)

func (s *Service) Pair(ctx context.Context, instance chat.Provider, msg *chat.Message, repo repository.Repository, code int64) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	text := ""
	if code == s.pairKey {
		text = "Successfully paired!"
		err := repo.CreateUser(ctx, types.User{
			ID:       0,
			UserID:   msg.UserID,
			Platform: instance.Name(),
			Type:     types.UserRoot,
		})
		if err != nil {
			text += " But we failed to create the user! Error: " + err.Error()
			return fmt.Errorf(text)
		}

		text += "\nEnjoy using your personal Lobster!"
		return nil
	}

	text = "Pair failed, invalid pair code, please check your code!"
	return fmt.Errorf(text)
}
