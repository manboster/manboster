package patches

import "sync"

var patchRegistry []Patch
var lock sync.RWMutex

func Register(p Patch) {
	lock.Lock()
	defer lock.Unlock()

	patchRegistry = append(patchRegistry, p)
}

func Patches() []Patch {
	lock.RLock()
	defer lock.RUnlock()

	return patchRegistry
}
