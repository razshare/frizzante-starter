package handlers

import (
	"github.com/razshare/frizzante/actions"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/state"
)

func Todos(connection *connections.Connection) {
	session := sessions.Start(sessions.New(connection, state.New()))
	defer sessions.Save(session)

	actions.SendView(connection, views.View{
		Name: "Todos",
		Data: map[string]any{
			"todos": session.State.Todos,
		},
	})
}
