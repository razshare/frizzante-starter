package sessions

import (
	"encoding/json"
	"github.com/razshare/frizzante"
	"main/lib/notifiers"
	"time"
)

func Adapter(session *frizzante.Session[Data]) {
	session.WithExistsHandler(func() bool {
		return Archive.Has(session.Id, Key)
	})

	session.WithLoadHandler(func() {
		data := Archive.Get(session.Id, Key)
		unmarshalError := json.Unmarshal(data, &session.Data)
		if nil != unmarshalError {
			notifiers.Console.SendError(unmarshalError)
		}
	})

	session.WithSaveHandler(func() {
		data, marshalError := json.Marshal(session.Data)
		if nil != marshalError {
			notifiers.Console.SendError(marshalError)
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
