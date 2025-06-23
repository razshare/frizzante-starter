package lib

type State struct {
	Todos []Todo
}

type Todo struct {
	Checked     bool
	Description string
}

func InitialTodos() []Todo {
	return []Todo{
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Do laundry"},
		{Checked: false, Description: "Pet the cat."},
		{Checked: false, Description: "Cook"},
		{Checked: false, Description: "Pet the cat."},
	}
}
