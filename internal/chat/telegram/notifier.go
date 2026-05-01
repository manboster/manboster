package telegram

import (
	"context"
	"strconv"
	"sync"
)

type Notifier struct {
	Active    bool
	ChatId    int64
	MessageId int
	Cancel    context.CancelFunc
}

var notifierLock sync.Mutex
var notifierMap = make(map[string]Notifier)

func notifierWrite(chatID int64, msgId int, cancel context.CancelFunc) {
	notifierLock.Lock()
	defer notifierLock.Unlock()

	chatId := strconv.FormatInt(chatID, 10)
	n, avail := notifierMap[chatId]
	if !avail {
		notifierMap[chatId] = Notifier{}
	}
	n.Active = true
	n.Cancel = cancel
	n.MessageId = msgId
	n.ChatId = chatID
	notifierMap[chatId] = n
}

func notifierCancel(chatId string) (int64, int) {
	notifierLock.Lock()
	defer notifierLock.Unlock()
	n, avail := notifierMap[chatId]
	if !avail {
		return 0, 0
	}
	if n.Cancel != nil {
		n.Cancel()
		return n.ChatId, n.MessageId
	}
	delete(notifierMap, chatId)
	return 0, 0
}
