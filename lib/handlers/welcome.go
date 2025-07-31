package handlers

import (
	"github.com/razshare/frizzante/actions"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
)

func Welcome(c *connections.Connection) {
	actions.SendView(c, views.View{Name: "Welcome"})
}
