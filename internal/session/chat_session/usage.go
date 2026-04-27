package chat_session

import (
	"github.com/manboster/manboster/spec/llm"
)

// Usage returns in, out and tot tokens
func (m *Manager) Usage(sid string) (int, int, int) {
	pTokens := 0
	cTokens := 0
	tTokens := 0

	m.Lock.RLock()
	defer m.Lock.RUnlock()

	s, avail := m.Sessions[sid]
	if !avail {
		return -1, -1, -1
	}
	for _, e := range s.Events {
		if e.EventType&llm.EventUsage != 0 && e.Usage != nil {
			pTokens += e.Usage.PromptTokens
			cTokens += e.Usage.CompletionTokens
			tTokens += e.Usage.TotalTokens
		}
	}
	return pTokens, cTokens, tTokens
}
