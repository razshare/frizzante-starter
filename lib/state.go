package lib

import "time"

type Item struct {
	Checked     bool   `json:"checked"`
	Description string `json:"description"`
}

type State struct {
	Items        []Item
	LastActivity time.Time
}

// InitializeState initialize session state.
func InitializeState() State {
	return State{
		LastActivity: time.Now(),
		Items: []Item{
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Do laundry"},
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Cook"},
			{Checked: false, Description: "Pet the cat."},
		},
	}
}
