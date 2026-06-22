package onboard

import (
	"sync"

	"github.com/manboster/manboster/internal/util"
)

// Service defines onboard instance
type Service struct {
	lock    sync.Mutex
	pairKey int64
	retry   int
	active  bool
}

// New creates an onboard service
func New() *Service {
	return &Service{
		lock:    sync.Mutex{},
		pairKey: util.RandomNumber(100000, 999999),
		retry:   0,
		active:  true,
	}
}
