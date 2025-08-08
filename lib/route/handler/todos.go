package handler

import (
	"github.com/razshare/frizzante/conn"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"github.com/razshare/frizzante/view"
	"main/lib/session"
)

func Todos(c *conn.Conn) {
	s := session.Start(receive.SessionId(c))

	send.View(c, view.View{Name: "Todos", Data: map[string]any{
		"todos": s.Todos,
	}})
}
