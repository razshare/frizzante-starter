package handlers

import (
	"github.com/razshare/frizzante/act"
	"github.com/razshare/frizzante/server"
	"github.com/razshare/frizzante/view"
	"main/lib/session"
	"strconv"
)

func Uncheck(c *server.Connection) {
	s := session.Start(act.ReceiveSessionId(c))

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

	l := int64(len(s.Todos))
	if i >= l {
		// Index is out of bounds, ignore the request.
		act.SendNavigate(c, "/todos")
		return
	}

	s.Todos[i].Checked = false

	act.SendNavigate(c, "/todos")
}
