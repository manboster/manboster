package onboard

import (
	"context"
	"fmt"

	"github.com/manboster/manboster/internal/i18n"
	"github.com/manboster/manboster/internal/i18n/keys"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/internal/repository/types"
	"github.com/manboster/manboster/spec/chat"
)

func (s *Service) Pair(ctx context.Context, instance chat.Provider, msg *chat.Message, repo repository.Repository, code int64) error {
	if !s.Active() {
		return nil
	}

	text := ""
	if code == s.pairKey {
		text = i18n.T(keys.EngineOnboardPairSuccess)
		err := repo.CreateUser(ctx, types.User{
			ID:       0,
			UserID:   msg.UserID,
			Platform: instance.Name(),
			Type:     types.UserRoot,
		})
		if err != nil {
			text += fmt.Sprintf(i18n.T(keys.EngineOnboardPairUserError), err.Error())
			return fmt.Errorf(text)
		}

		text += i18n.T(keys.EngineOnboardPairSuccessMsg)
		s.Deactivate()
		return nil
	}

	text = i18n.T(keys.EngineOnboardPairFailed)
	return fmt.Errorf(text)
}
