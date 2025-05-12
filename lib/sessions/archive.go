package sessions

import (
	f "github.com/razshare/frizzante"
	"main/lib"
	"time"
)

var archiveKey = "session.json"
var archive = f.ArchiveCreateOnDisk(".sessions", time.Hour)

func Archive(session *f.Session[lib.State]) {
	f.SessionWithLoadHandler(session, func() {
		if !f.ArchiveHas(archive, session.Id, archiveKey) {
			session.State = lib.InitializeState()
			f.ArchiveSetAsJson(archive, session.Id, archiveKey, session.State)
			return
		}

		session.State = f.ArchiveGetJson[lib.State](archive, session.Id, archiveKey)
	})

	f.SessionWithValidateHandler(session, func() bool {
		return true
	})

	f.SessionWithSaveHandler(session, func() {
		f.ArchiveSetAsJson(archive, session.Id, archiveKey, session.State)
	})

	f.SessionWithDestroyHandler(session, func() {
		f.ArchiveRemove(archive, session.Id, archiveKey)
	})
}
