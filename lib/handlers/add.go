package handlers

import (
	"github.com/razshare/frizzante/actions"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib"
)

func Add(c *connections.Connection) {
	s := sessions.Start(sessions.New(c, lib.InitialState()))
	defer sessions.Save(s)

	description := actions.ReceiveQuery(c, "description")

	if "" == description {
		actions.SendView(c, views.View{Name: "Todos", Data: map[string]any{
			"todos": s.State.Todos,
			"error": "todo description cannot be empty",
		}})
		return
	}

	s.State.Todos = append(s.State.Todos, lib.Todo{Checked: false, Description: description})

	actions.SendNavigate(c, "/todos")
}
