package todos

import (
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"github.com/razshare/frizzante/view"
	"main/lib/session"
	"strconv"
)

func Remove(c *client.Client) {
	s := session.Start(receive.SessionId(c))

	l := int64(len(s.Todos))
	if 0 == l {
		// No index found, ignore the request.
		send.View(c, view.View{
			Name: "Todos",
			Data: map[string]any{
				"todos": s.Todos,
			},
		})
		return
	}

	is := receive.Query(c, "index")
	if is == "" {
		// No index found, ignore the request.
		send.View(c, view.View{
			Name: "Todos",
			Data: map[string]any{
				"todos": s.Todos,
			},
		})
		return
	}

	i, e := strconv.ParseInt(is, 10, 64)
	if nil != e {
		send.View(c, view.View{
			Name: "Todos",
			Data: map[string]any{
				"todos": s.Todos,
				"error": e.Error(),
			},
		})
		return
	}
	if i >= l {
		// Index is out of bounds, ignore the request.
		send.View(c, view.View{
			Name: "Todos",
			Data: map[string]any{
				"todos": s.Todos,
			},
		})
		return
	}

	s.Todos = append(
		s.Todos[:i],
		s.Todos[i+1:]...,
	)

	send.View(c, view.View{
		Name: "Todos",
		Data: map[string]any{
			"todos": s.Todos,
		},
	})
}
