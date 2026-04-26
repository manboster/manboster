package repository

import (
	"errors"

	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound
var ErrDuplicateSoulScope = errors.New("duplicate soul scope")
var ErrDuplicateMemoryScope = errors.New("duplicate memory scope")
