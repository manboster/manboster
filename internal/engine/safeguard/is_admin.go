package safeguard

import (
	"github.com/manboster/manboster/spec/schema"
)

func (s *Service) IsAdmin(userType schema.UserType) bool {
	return userType >= schema.UserAdmin
}
