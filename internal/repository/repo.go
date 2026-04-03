package repository

import "gorm.io/gorm"

type Repository interface {
	UserRepository
}

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}
