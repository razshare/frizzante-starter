package handlers

import (
	"github.com/razshare/frizzante/actions"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib"
)

func Todos(c *connections.Connection) {
	s := sessions.Start(sessions.New(c, lib.InitialState()))
	defer sessions.Save(s)

	actions.SendView(c, views.View{Name: "Todos", Data: map[string]any{
		"todos": s.State.Todos,
	}})
}
