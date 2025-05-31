package handlers

import "github.com/razshare/frizzante"

func GetDefault(c *frizzante.Connection) {
	c.SendFileOrElse(func() { GetWelcome(c) })
}
