package handlers

import "github.com/razshare/frizzante/connections"

func Default(connection *connections.Connection) {
	connection.SendFileOrElse(func() { Welcome(connection) })
}
