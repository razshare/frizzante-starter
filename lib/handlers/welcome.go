package handlers

import "github.com/razshare/frizzante/frz"

func Welcome(c *frz.Connection) {
	c.SendView(frz.View{Name: "Welcome"})
}
