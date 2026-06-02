package safeguard

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/manboster/manboster/internal/repository"
	"github.com/manboster/manboster/spec/schema"
)

// UserType returns current user's type
func (s *Service) UserType(ctx context.Context, name string, userId string) schema.UserType {
	uInfo, err := s.repo.UserInfo(ctx, name, userId)
	if err != nil {
		// cause error!
		if !errors.Is(err, repository.ErrNotFound) {
			color.Red(fmt.Sprintf("[Manboster Safeguard] We encountered an error while fetching user data from repository, error: %q", err))
		}
		return schema.UserUnknown
	}
	return uInfo.Type
}
