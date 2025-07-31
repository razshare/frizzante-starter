package handlers

import "github.com/razshare/frizzante/servers"

func Default(con *servers.Connection) {
	con.SendFileOrElse(func() { Welcome(con) })
}
