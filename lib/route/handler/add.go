package handler

import (
	"github.com/razshare/frizzante/act"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
	"main/lib/sessions"
)

func Add(c *connections.Connection) {
	s := sessions.Start(act.ReceiveSessionId(c))
	d := act.ReceiveQuery(c, "description")

	if "" == d {
		act.SendView(c, views.View{Name: "Todos", Data: map[string]any{
			"todos": s.Todos,
			"error": "todo description cannot be empty",
		}})
		return
	}

	s.Todos = append(s.Todos, sessions.Todo{
		Checked:     false,
		Description: d,
	})

	act.SendNavigate(c, "/todos")
}
