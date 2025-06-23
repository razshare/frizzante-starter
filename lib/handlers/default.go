package handlers

import "github.com/razshare/frizzante/libcon"

func Default(c *libcon.Connection) {
	c.SendFileOrElse(func() { Welcome(c) })
}
