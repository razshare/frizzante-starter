package handler

import (
	"github.com/razshare/frizzante/act"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
)

func Welcome(c *connections.Connection) {
	act.SendView(c, views.View{Name: "Welcome"})
}
