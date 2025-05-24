package lib

import "time"

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
