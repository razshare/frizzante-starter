package handler

import (
	"github.com/razshare/frizzante/act"
	"github.com/razshare/frizzante/connections"
)

func Default(c *connections.Connection) {
	act.SendFileOrElse(c, func() { Welcome(c) })
}
