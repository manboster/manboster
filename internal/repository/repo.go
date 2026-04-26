package repository

import "gorm.io/gorm"

type Repository interface {
	UserRepository
	ChatRepository
	ChatDataRepository
	SessionRepository
	SoulRepository
	MemoryRepository
}

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}
