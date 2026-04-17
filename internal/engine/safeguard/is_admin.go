package safeguard

import "github.com/manboster/manboster/internal/repository/types"

func (s *Service) IsAdmin(userType types.UserType) bool {
	return userType >= types.UserAdmin
}
