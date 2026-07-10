package _30

import (
	"github.com/manboster/manboster/internal/migration/patches"
)

func init() {
	patches.Register(cacheMigrationPatch)
}
