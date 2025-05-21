package lib

import "time"

func NewData() Data {
	return Data{
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

type Data struct {
	Items        []Item    `json:"items"`
	LastActivity time.Time `json:"lastActivity"`
	Expired      bool      `json:"expired"`
}

type Item struct {
	Checked     bool   `json:"checked"`
	Description string `json:"description"`
}
