package sessions

import "time"

type Data struct {
	Todos        []Todo
	LastActivity time.Time
	Expired      bool
}

type Todo struct {
	Checked     bool
	Description string
}
