package handler

import (
	"github.com/razshare/frizzante/conn"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"github.com/razshare/frizzante/view"
	"main/lib/session"
	"strconv"
)

func Uncheck(c *conn.Conn) {
	s := session.Start(receive.SessionId(c))

	is := receive.Query(c, "index")
	if "" == is {
		// No index found, ignore the request.
		send.Navigate(c, "/todos")
		return
	}

	i, e := strconv.ParseInt(is, 10, 64)
	if nil != e {
		send.View(c, view.View{Name: "Todos", Data: map[string]any{
			"error": e.Error(),
		}})
		return
	}

	l := int64(len(s.Todos))

	if i >= l {
		// Index is out of bounds, ignore the request.
		send.Navigate(c, "/todos")
		return
	}

	s.Todos[i].Checked = false

	send.Navigate(c, "/todos")
}
