package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/state"
)

func Todos(connection *connections.Connection) {
	session := sessions.Start(connection, state.New())
	defer session.Save()

	connection.SendView(views.View{Name: "Todos", Data: map[string]any{
		"todos": session.State.Todos,
	}})
}
