package sessions

import (
	f "github.com/razshare/frizzante"
	"time"
)

var archive = f.ArchiveCreateOnDisk(".sessions", time.Second/2)

// Archive builds sessions using a disk archive.
func Archive(session *f.Session) {
	sessionId := f.SessionId(session)

	f.SessionWithGetHandler(session, func(key string) []byte {
		return f.ArchiveGet(archive, sessionId, key)
	})

	f.SessionWithSetHandler(session, func(key string, value []byte) {
		f.ArchiveSet(archive, sessionId, key, value)
	})

	f.SessionWithHasHandler(session, func(key string) bool {
		return f.ArchiveHas(archive, sessionId, key)
	})

	f.SessionWithDestroyHandler(session, func() {
		f.ArchiveRemoveDomain(archive, sessionId)
	})
}
