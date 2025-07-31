package handlers

import (
	"github.com/razshare/frizzante/actions"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/state"
	"strconv"
)

func Remove(connection *connections.Connection) {
	session := sessions.Start(connection, state.New())
	defer sessions.Save(session)

	count := int64(len(session.State.Todos))

	if 0 == count {
		// No index found, ignore the request.
		actions.SendNavigate(connection, "/todos")
		return
	}

	indexString := actions.ReceiveQuery(connection, "index")
	if "" == indexString {
		// No index found, ignore the request.
		actions.SendNavigate(connection, "/todos")
		return
	}

	index, indexError := strconv.ParseInt(indexString, 10, 64)
	if nil != indexError {
		actions.SendView(connection, views.View{
			Name: "Todos",
			Data: map[string]any{
				"error": indexError.Error(),
			},
		})
		return
	}

	if index >= count {
		// Index is out of bounds, ignore the request.
		actions.SendNavigate(connection, "/todos")
		return
	}

	session.State.Todos = append(
		session.State.Todos[:index],
		session.State.Todos[index+1:]...,
	)

	actions.SendNavigate(connection, "/todos")
}
