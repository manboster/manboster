package telegram

import (
	"context"
	"sync"
)

type Notifier struct {
	Active bool
	Cancel context.CancelFunc
}

var notifierLock sync.Mutex
var notifierMap = make(map[string]Notifier)

func notifierWrite(chatId string, cancel context.CancelFunc) {
	notifierLock.Lock()
	defer notifierLock.Unlock()
	n, avail := notifierMap[chatId]
	if !avail {
		notifierMap[chatId] = Notifier{}
	}
	n.Active = true
	n.Cancel = cancel
}

func notifierCancel(chatId string) {
	notifierLock.Lock()
	defer notifierLock.Unlock()
	n, avail := notifierMap[chatId]
	if !avail {
		return
	}
	n.Cancel()
	delete(notifierMap, chatId)
}
