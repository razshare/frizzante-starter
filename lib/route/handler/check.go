package handler

import (
	"github.com/razshare/frizzante/act"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
	"main/lib/sessions"
	"strconv"
)

func Check(c *connections.Connection) {
	s := sessions.Start(act.ReceiveSessionId(c))

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

	l := int64(len(s.Todos))
	if i >= l {
		// Index is out of bounds, ignore the request.
		act.SendNavigate(c, "/todos")
		return
	}

	s.Todos[i].Checked = true

	act.SendNavigate(c, "/todos")
}
