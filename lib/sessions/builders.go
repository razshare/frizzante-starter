package sessions

import (
	"encoding/json"
	f "github.com/razshare/frizzante"
	"main/lib"
	"time"
)

var key = "session.json"
var notifier = f.NotifierCreate()
var archive = f.ArchiveCreateOnDisk(".sessions", time.Second/2)

func Archived(session *f.Session[lib.State]) {
	f.SessionWithExistsHandler(session, func() bool {
		return f.ArchiveHas(archive, session.Id, key)
	})

	f.SessionWithLoadHandler(session, func() {
		data := f.ArchiveGet(archive, session.Id, key)
		unmarshalError := json.Unmarshal(data, &session.Data)
		if nil != unmarshalError {
			f.NotifierSendError(notifier, unmarshalError)
		}
	})

	f.SessionWithSaveHandler(session, func() {
		data, marshalError := json.Marshal(session.Data)
		if nil != marshalError {
			f.NotifierSendError(notifier, marshalError)
			return
		}
		f.ArchiveSet(archive, session.Id, key, data)
	})

	f.SessionWithDestroyHandler(session, func() {
		f.ArchiveRemoveDomain(archive, session.Id)
	})

	if f.SessionExists(session) {
		f.SessionLoad(session)
		return
	}

	session.Data = lib.InitialState()
}
