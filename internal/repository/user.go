package repository

import "context"

type UserRepository interface {
	UserCounts(ctx context.Context) (int, error) // get user's counts
}

// UserCounts gets the total number of users.
func (repo *Repo) UserCounts(ctx context.Context) (int, error) {
	return 0, nil
}
