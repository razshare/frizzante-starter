package handlers

import frz "github.com/razshare/frizzante"

func GetDefault(c *frz.Connection) {
	c.SendFileOrElse(func() { GetWelcome(c) })
}
