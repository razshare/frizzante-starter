package handlers

import (
	"github.com/razshare/frizzante/actions"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/state"
)

func Add(connection *connections.Connection) {
	session := sessions.Start(sessions.New(connection, state.New()))
	defer sessions.Save(session)

	description := actions.ReceiveQuery(connection, "description")

	if "" == description {
		actions.SendView(connection, views.View{
			Name: "Todos",
			Data: map[string]any{
				"todos": session.State.Todos,
				"error": "todo description cannot be empty",
			},
		})
		return
	}

	session.State.Todos = append(session.State.Todos, state.Todo{
		Checked:     false,
		Description: description,
	})

	actions.SendNavigate(connection, "/todos")
}
