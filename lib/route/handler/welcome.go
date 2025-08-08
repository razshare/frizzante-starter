package handler

import (
	"github.com/razshare/frizzante/conn"
	"github.com/razshare/frizzante/send"
	"github.com/razshare/frizzante/view"
)

func Welcome(c *conn.Conn) {
	send.View(c, view.View{Name: "Welcome"})
}
