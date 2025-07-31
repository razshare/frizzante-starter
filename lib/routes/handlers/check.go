package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/state"
	"strconv"
)

func Check(connection *connections.Connection) {
	session := sessions.New(connection, state.Default()).Start()
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

	session.State.Todos[index].Checked = true

	connection.SendNavigate("/todos")
}
