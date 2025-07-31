package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/state"
)

func Todos(connection *connections.Connection) {
	session := sessions.New(connection, state.Default()).Start()
	defer session.Save()

	connection.SendView(views.View{Name: "Todos", Data: map[string]any{
		"todos": session.State.Todos,
	}})
}
