package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/state"
	"strconv"
)

func Remove(connection *connections.Connection) {
	session := sessions.Start(connection, state.Default())
	defer session.Save()

	count := int64(len(session.State.Todos))

	if 0 == count {
		// No index found, ignore the request.
		connection.SendNavigate("/todos")
		return
	}

	indexString := connection.ReceiveQuery("index")
	if "" == indexString {
		// No index found, ignore the request.
		connection.SendNavigate("/todos")
		return
	}

	index, indexError := strconv.ParseInt(indexString, 10, 64)
	if nil != indexError {
		connection.SendView(views.View{Name: "Todos", Data: map[string]any{
			"error": indexError.Error(),
		}})
		return
	}

	if index >= count {
		// Index is out of bounds, ignore the request.
		connection.SendNavigate("/todos")
		return
	}

	session.State.Todos = append(
		session.State.Todos[:index],
		session.State.Todos[index+1:]...,
	)

	connection.SendNavigate("/todos")
}
