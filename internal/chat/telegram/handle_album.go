package telegram

import (
	"context"
	"sync"
	"time"

	"github.com/manboster/manboster/spec/chat"
)

var AlbumHandlerTimerMap = make(map[string]*time.Timer)
var AlbumDataMap = make(map[string]*chat.Message)
var lock sync.RWMutex

func GetData(id string) (*chat.Message, bool) {
	lock.RLock()
	defer lock.RUnlock()
	data, ok := AlbumDataMap[id]
	return data, ok
}

func SetData(id string, data *chat.Message) {
	lock.Lock()
	defer lock.Unlock()
	AlbumDataMap[id] = data
}

func DeleteData(id string) {
	lock.Lock()
	defer lock.Unlock()
	delete(AlbumDataMap, id)
}

func SetTimer(id string, onMsg func(message *chat.Message)) (*time.Timer, bool) {
	lock.Lock()
	defer lock.Unlock()
	t, ok := AlbumHandlerTimerMap[id]
	if ok {
		t.Reset(5 * time.Second)
		return t, ok
	}
	timer := time.NewTimer(5 * time.Second)
	AlbumHandlerTimerMap[id] = timer
	go AlbumRunner(context.Background(), timer, id, onMsg)
	return timer, ok
}

func DeleteTimer(id string) {
	lock.Lock()
	defer lock.Unlock()
	delete(AlbumHandlerTimerMap, id)
}

func AlbumRunner(ctx context.Context, timer *time.Timer, id string, onMsg func(*chat.Message)) {
	for {
		select {
		case <-timer.C:
			msg, ok := GetData(id)
			if ok && onMsg != nil && msg != nil {
				onMsg(msg)
			}
			DeleteTimer(id)
			DeleteData(id)
			return
		case <-ctx.Done():
			return
		}
	}
}
