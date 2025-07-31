package handlers

import (
	"github.com/razshare/frizzante/actions"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib"
	"strconv"
)

func Remove(c *connections.Connection) {
	s := sessions.Start(sessions.New(c, lib.InitialState()))
	defer sessions.Save(s)

	count := int64(len(s.State.Todos))

	if 0 == count {
		// No index found, ignore the request.
		actions.SendNavigate(c, "/todos")
		return
	}

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

	if index >= count {
		// Index is out of bounds, ignore the request.
		actions.SendNavigate(c, "/todos")
		return
	}

	s.State.Todos = append(s.State.Todos[:index], s.State.Todos[index+1:]...)

	actions.SendNavigate(c, "/todos")
}
