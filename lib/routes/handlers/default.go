package handlers

import (
	"github.com/razshare/frizzante/act"
	"github.com/razshare/frizzante/server"
)

func Default(c *server.Connection) {
	act.SendFileOrElse(c, func() { Welcome(c) })
}
