package handlers

import "github.com/razshare/frizzante"

func GetWelcome(c *frizzante.Connection) {
	c.SendView(frizzante.View{Name: "Welcome"})
}
