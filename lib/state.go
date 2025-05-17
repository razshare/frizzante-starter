package lib

import "time"

type Item struct {
	Checked     bool
	Description string
}

type State struct {
	Items        []Item
	LastActivity time.Time
	Expired      bool
}

func InitialState() State {
	return State{
		Items: []Item{
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
