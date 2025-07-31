package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/state"
)

func Todos(con *connections.Connection) {
	session := sessions.Start(con, state.Default())
	defer session.Save()

	con.SendView(views.View{Name: "Todos", Data: map[string]any{
		"todos": session.State.Todos,
	}})
}
