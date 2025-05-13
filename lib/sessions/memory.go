package sessions

import (
	f "github.com/razshare/frizzante"
)

var memory = map[string]map[string][]byte{}

// Memory builds sessions in memory.
func Memory(session *f.Session) {
	sessionId := f.SessionId(session)
	memory[sessionId] = map[string][]byte{}

	f.SessionWithGetHandler(session, func(key string) []byte {
		return memory[sessionId][key]
	})

	f.SessionWithSetHandler(session, func(key string, value []byte) {
		memory[sessionId][key] = value
	})

	f.SessionWithHasHandler(session, func(key string) bool {
		_, hasKey := memory[sessionId][key]
		return hasKey
	})

	f.SessionWithDestroyHandler(session, func() {
		delete(memory, sessionId)
	})
}
