package handlers

import "github.com/razshare/frizzante/frz"

func GetDefault(c *frz.Connection) {
	c.SendFileOrElse(func() { GetWelcome(c) })
}
