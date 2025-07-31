package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/state"
	"strconv"
)

func Uncheck(connection *connections.Connection) {
	session := sessions.Start(connection, state.New())
	defer session.Save()

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

	count := int64(len(session.State.Todos))
	if index >= count {
		// Index is out of bounds, ignore the request.
		connection.SendNavigate("/todos")
		return
	}

	session.State.Todos[index].Checked = false

	connection.SendNavigate("/todos")
}
