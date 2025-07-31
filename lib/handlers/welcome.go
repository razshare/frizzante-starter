package handlers

import (
	"github.com/razshare/frizzante/servers"
	"github.com/razshare/frizzante/views"
)

func Welcome(con *servers.Connection) {
	con.SendView(views.View{Name: "Welcome"})
}
