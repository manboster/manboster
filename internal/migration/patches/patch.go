package patches

import (
	"github.com/manboster/manboster/internal/config"
)

type DetectFunc func(conf config.Config) bool
type MigrateFunc func(conf config.Config) (config.Config, error)

type Patch struct {
	Detect      DetectFunc
	Migrate     MigrateFunc
	Description string
	Introduce   string // introduced version in this patch
	V           int    // the big V, following to config.V
}
