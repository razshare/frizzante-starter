package handlers

import (
	"github.com/razshare/frizzante/actions"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib"
	"strconv"
)

func Check(c *connections.Connection) {
	s := sessions.Start(sessions.New(c, lib.InitialState()))
	defer sessions.Save(s)

	indexString := actions.ReceiveQuery(c, "index")
	if "" == indexString {
		// No index found, ignore the request.
		actions.SendNavigate(c, "/todos")
		return
	}

	index, indexError := strconv.ParseInt(indexString, 10, 64)
	if nil != indexError {
		actions.SendView(c, views.View{Name: "Todos", Data: map[string]any{
			"error": indexError.Error(),
		}})
		return
	}

	count := int64(len(s.State.Todos))
	if index >= count {
		// Index is out of bounds, ignore the request.
		actions.SendNavigate(c, "/todos")
		return
	}

	s.State.Todos[index].Checked = true

	actions.SendNavigate(c, "/todos")
}
