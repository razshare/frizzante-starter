package sessions

import (
	"encoding/json"
	"github.com/razshare/frizzante"
	"time"
)

var notifier = frizzante.NewNotifier()

func Adapter(session *frizzante.Session[Data]) {
	session.WithExistsHandler(func() bool {
		return Archive.Has(session.Id, Key)
	})

	session.WithLoadHandler(func() {
		data := Archive.Get(session.Id, Key)
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
		Archive.Set(session.Id, Key, data)
	})

	session.WithDestroyHandler(func() {
		Archive.RemoveDomain(session.Id)
	})

	if session.Exists() {
		session.Load()
		return
	}

	session.Data = Data{
		LastActivity: time.Now(),
		Todos: []Todo{
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Do laundry"},
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Cook"},
			{Checked: false, Description: "Pet the cat."},
		},
	}
}
