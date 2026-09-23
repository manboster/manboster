package repository

import "gorm.io/gorm"

// Repository is a universal call for repos
type Repository interface {
	UserRepository
	ChatRepository
	EventDataRepository
	SessionRepository
	SoulRepository
}

type Repo struct {
	*SessionRepo
	*SoulRepo
	*EventDataRepo
	*UserRepo
	*ChatRepo
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		SessionRepo: &SessionRepo{
			db: db,
		},
		SoulRepo: &SoulRepo{
			db: db,
		},
		EventDataRepo: &EventDataRepo{
			db: db,
		},
		UserRepo: &UserRepo{
			db: db,
		},
		ChatRepo: &ChatRepo{
			db: db,
		},
	}
}
