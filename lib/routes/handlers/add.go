package handlers

import (
	"github.com/razshare/frizzante/act"
	"github.com/razshare/frizzante/server"
	"github.com/razshare/frizzante/view"
	"main/lib/session"
)

func Add(c *server.Connection) {
	s := session.Start(act.ReceiveSessionId(c))
	d := act.ReceiveQuery(c, "description")

	if "" == d {
		act.SendView(c, view.View{Name: "Todos", Data: map[string]any{
			"todos": s.Todos,
			"error": "todo description cannot be empty",
		}})
		return
	}

	s.Todos = append(s.Todos, session.Todo{
		Checked:     false,
		Description: d,
	})

	act.SendNavigate(c, "/todos")
}
