package handlers

import (
	"github.com/razshare/frizzante/act"
	"github.com/razshare/frizzante/server"
	"github.com/razshare/frizzante/view"
)

func Welcome(c *server.Connection) {
	act.SendView(c, view.View{Name: "Welcome"})
}
