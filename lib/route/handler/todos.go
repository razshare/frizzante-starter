package handler

import (
	"github.com/razshare/frizzante/act"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
	"main/lib/sessions"
)

func Todos(c *connections.Connection) {
	s := sessions.Start(act.ReceiveSessionId(c))

	act.SendView(c, views.View{Name: "Todos", Data: map[string]any{
		"todos": s.Todos,
	}})
}
