package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/state"
)

func Add(connection *connections.Connection) {
	session := sessions.New(connection, state.Default()).Start()
	defer session.Save()

	description := connection.ReceiveQuery("description")

	if "" == description {
		connection.SendView(views.View{Name: "Todos", Data: map[string]any{
			"todos": session.State.Todos,
			"error": "todo description cannot be empty",
		}})
		return
	}

	session.State.Todos = append(session.State.Todos, state.Todo{
		Checked:     false,
		Description: description,
	})

	connection.SendNavigate("/todos")
}
