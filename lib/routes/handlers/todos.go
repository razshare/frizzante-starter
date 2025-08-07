package handlers

import (
	"github.com/razshare/frizzante/act"
	"github.com/razshare/frizzante/server"
	"github.com/razshare/frizzante/view"
	"main/lib/session"
)

func Todos(c *server.Connection) {
	s := session.Start(act.ReceiveSessionId(c))

	act.SendView(c, view.View{Name: "Todos", Data: map[string]any{
		"todos": s.Todos,
	}})
}
