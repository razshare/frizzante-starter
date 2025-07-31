package handlers

import (
	"github.com/razshare/frizzante/actions"
	"github.com/razshare/frizzante/connections"
)

func Default(c *connections.Connection) {
	actions.SendFileOrElse(c, func() { Welcome(c) })
}
