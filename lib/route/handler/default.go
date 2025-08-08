package handler

import (
	"github.com/razshare/frizzante/conn"
	"github.com/razshare/frizzante/send"
)

func Default(c *conn.Conn) {
	send.FileOrElse(c, func() { Welcome(c) })
}
