package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib"
	"strconv"
)

func Remove(con *connections.Connection) {
	session := sessions.StartWithState(con, lib.InitialState())
	defer session.Save()

	count := int64(len(session.State.Todos))

	if 0 == count {
		// No index found, ignore the request.
		con.SendNavigate("/todos")
		return
	}

	indexString := con.ReceiveQuery("index")
	if "" == indexString {
		// No index found, ignore the request.
		con.SendNavigate("/todos")
		return
	}

	index, indexError := strconv.ParseInt(indexString, 10, 64)
	if nil != indexError {
		con.SendView(views.View{Name: "Todos", Data: map[string]any{
			"error": indexError.Error(),
		}})
		return
	}

	if index >= count {
		// Index is out of bounds, ignore the request.
		con.SendNavigate("/todos")
		return
	}

	session.State.Todos = append(session.State.Todos[:index], session.State.Todos[index+1:]...)

	con.SendNavigate("/todos")
}
