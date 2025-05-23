package config

import (
	"encoding/json"
	f "github.com/razshare/frizzante"
	"time"
)

type SessionData struct {
	Todos        []Todo    `json:"todos"`
	LastActivity time.Time `json:"lastActivity"`
	Expired      bool      `json:"expired"`
}

type Todo struct {
	Checked     bool   `json:"checked"`
	Description string `json:"description"`
}

func NewSessionData() SessionData {
	return SessionData{
		Todos: []Todo{
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Do laundry"},
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Cook"},
			{Checked: false, Description: "Pet the cat."},
		},
		LastActivity: time.Now(),
		Expired:      false,
	}
}

var key = "session.json"
var notifier = f.NewNotifier()
var archive = f.NewArchiveOnDisk(".sessions", time.Second/2)

func SessionLoad(session *f.Session[SessionData]) {
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

	session.Data = NewSessionData()
}
