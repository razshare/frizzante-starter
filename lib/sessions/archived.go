package sessions

import (
	"encoding/json"
	f "github.com/razshare/frizzante"
	"main/lib"
	"time"
)

var key = "session.json"
var notifier = f.NewNotifier()
var archive = f.NewArchiveOnDisk(".sessions", time.Second/2)

func Archived(session *f.Session[lib.Data]) {
	session.WithExistsHandler(func() bool {
		return archive.Has(session.Id, key)
	})

	session.WithLoadHandler(func() {
		data := archive.Get(session.Id, key)
		unmarshalError := json.Unmarshal(data, &session.Data)
		if nil != unmarshalError {
			notifier.SendError(unmarshalError)
		}
	})

	session.WithSaveHandler(func() {
		data, marshalError := json.Marshal(session.Data)
		if nil != marshalError {
			notifier.SendError(marshalError)
			return
		}
		archive.Set(session.Id, key, data)
	})

	session.WithDestroyHandler(func() {
		archive.RemoveDomain(session.Id)
	})

	if session.Exists() {
		session.Load()
		return
	}

	session.Data = lib.InitialData()
}
