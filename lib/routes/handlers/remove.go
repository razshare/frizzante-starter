package handlers

import (
	"github.com/razshare/frizzante/act"
	"github.com/razshare/frizzante/server"
	"github.com/razshare/frizzante/view"
	"main/lib/session"
	"strconv"
)

func Remove(c *server.Connection) {
	s := session.Start(act.ReceiveSessionId(c))

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
		act.SendView(c, view.View{Name: "Todos", Data: map[string]any{
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
