package routes

import (
	"github.com/razshare/frizzante"
	"main/lib/guards"
	"main/lib/sessions"
	"main/lib/value"
	"strconv"
)

func GetTodos(req *frizzante.Request, res *frizzante.Response) {
	if !frizzante.AllGuardsPass(req, res, guards.NotExpired) {
		return
	}

	session := frizzante.SessionStart(req, res, sessions.Adapter)

	res.SendView(frizzante.View{Name: "Todos", Data: session.Data.Todos})
}

func PostTodos(req *frizzante.Request, res *frizzante.Response) {

	session := frizzante.SessionStart(req, res, sessions.Adapter)
	form := req.ReceiveForm()

	if form.Has("check") {
		id := value.Wrap(strconv.ParseInt(form.Get("check"), 10, 32))
		session.Data.Todos[id.Value].Checked = true
	} else if form.Has("uncheck") {
		id := value.Wrap(strconv.ParseInt(form.Get("uncheck"), 10, 32))
		session.Data.Todos[id.Value].Checked = false
	}
	session.Save()
	res.SendView(frizzante.View{
		Name: "Todos",
		Data: session.Data.Todos,
	})
}
