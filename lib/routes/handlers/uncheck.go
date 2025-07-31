package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib/state"
	"strconv"
)

func Uncheck(con *connections.Connection) {
	session := sessions.Start(con, state.Default())
	defer session.Save()

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

	count := int64(len(session.State.Todos))
	if index >= count {
		// Index is out of bounds, ignore the request.
		con.SendNavigate("/todos")
		return
	}

	session.State.Todos[index].Checked = false

	con.SendNavigate("/todos")
}
