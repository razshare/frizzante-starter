package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
)

func Welcome(connection *connections.Connection) {
	connection.SendView(views.View{Name: "Welcome"})
}
