package handlers

import "github.com/razshare/frizzante/libcon"

func Default(con *libcon.Connection) {
	con.SendFileOrElse(func() { Welcome(con) })
}
