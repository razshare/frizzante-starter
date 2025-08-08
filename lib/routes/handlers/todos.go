package handlers

import (
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"github.com/razshare/frizzante/view"
	"main/lib/session"
)

func Todos(c *client.Client) {
	s := session.Start(receive.SessionId(c))
	send.View(c, view.View{
		Name: "Todos",
		Data: map[string]any{
			"todos": s.Todos,
		},
	})
}
