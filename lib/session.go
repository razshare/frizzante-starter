package lib

import (
	"encoding/json"
	f "github.com/razshare/frizzante"
	"time"
)

var sessionKey = "session.json"
var sessionNotifier = f.NewNotifier()
var sessionArchive = f.NewArchiveOnDisk(".sessions", time.Second/2)

func SessionAdapter(session *f.Session[SessionData]) {
	session.WithExistsHandler(func() bool {
		return sessionArchive.Has(session.Id, sessionKey)
	})

	session.WithLoadHandler(func() {
		data := sessionArchive.Get(session.Id, sessionKey)
		unmarshalError := json.Unmarshal(data, &session.Data)
		if nil != unmarshalError {
			sessionNotifier.SendError(unmarshalError)
		}
	})

	session.WithSaveHandler(func() {
		data, marshalError := json.Marshal(session.Data)
		if nil != marshalError {
			sessionNotifier.SendError(marshalError)
			return
		}
		sessionArchive.Set(session.Id, sessionKey, data)
	})

	session.WithDestroyHandler(func() {
		sessionArchive.RemoveDomain(session.Id)
	})

	if session.Exists() {
		session.Load()
		return
	}

	session.Data = NewSessionData()
}
