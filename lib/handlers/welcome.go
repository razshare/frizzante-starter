package handlers

import frz "github.com/razshare/frizzante"

func GetWelcome(c *frz.Connection) {
	c.SendView(frz.View{Name: "Welcome"})
}
