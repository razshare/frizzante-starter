package handlers

import (
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libview"
)

func Welcome(c *libcon.Connection) {
	c.SendView(libview.View{Name: "Welcome"})
}
