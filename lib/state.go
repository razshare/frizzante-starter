package lib

type Item struct {
	Checked     bool   `json:"checked"`
	Description string `json:"description"`
}

type State struct {
	Items []Item `json:"items"`
}

// InitializeState initialize session state.
func InitializeState() State {
	return State{
		Items: []Item{
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Do laundry"},
			{Checked: false, Description: "Pet the cat."},
			{Checked: false, Description: "Cook"},
			{Checked: false, Description: "Pet the cat."},
		},
	}
}
