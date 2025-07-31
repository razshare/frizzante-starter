package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/state"
)

func Add(con *connections.Connection) {
	session := sessions.Start(con, state.Default())
	defer session.Save()

	description := con.ReceiveQuery("description")

	if "" == description {
		con.SendView(views.View{Name: "Todos", Data: map[string]any{
			"todos": session.State.Todos,
			"error": "todo description cannot be empty",
		}})
		return
	}

	session.State.Todos = append(session.State.Todos, state.Todo{
		Checked:     false,
		Description: description,
	})

	con.SendNavigate("/todos")
}
