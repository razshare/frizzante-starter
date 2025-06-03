package handlers

import "github.com/razshare/frizzante/frz"

func GetWelcome(c *frz.Connection) {
	c.SendView(frz.View{Name: "Welcome"})
}
