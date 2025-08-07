package handler

import (
	"github.com/razshare/frizzante/act"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
	"main/lib/sessions"
	"strconv"
)

func Remove(c *connections.Connection) {
	s := sessions.Start(act.ReceiveSessionId(c))

	l := int64(len(s.Todos))

	if 0 == l {
		// No index found, ignore the request.
		act.SendNavigate(c, "/todos")
		return
	}

	is := act.ReceiveQuery(c, "index")
	if "" == is {
		// No index found, ignore the request.
		act.SendNavigate(c, "/todos")
		return
	}

	i, e := strconv.ParseInt(is, 10, 64)
	if nil != e {
		act.SendView(c, views.View{Name: "Todos", Data: map[string]any{
			"error": e.Error(),
		}})
		return
	}

	if i >= l {
		// Index is out of bounds, ignore the request.
		act.SendNavigate(c, "/todos")
		return
	}

	s.Todos = append(
		s.Todos[:i],
		s.Todos[i+1:]...,
	)

	act.SendNavigate(c, "/todos")
}
