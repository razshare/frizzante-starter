package lib

import (
	"time"
)

type UserSession struct {
	Items        []Item    `json:"items"`
	LastActivity time.Time `json:"lastActivity"`
}

type Item struct {
	Checked     bool   `json:"checked"`
	Description string `json:"description"`
}

// InitializeState initializes the session state.
func InitializeState() *UserSession {
	return &UserSession{
		Items: []Item{
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Do laundry"},
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Cook"},
			{Checked: false, Description: "Pet the cat."},
		},
	}
}
