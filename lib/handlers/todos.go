package handlers

import (
	"github.com/razshare/frizzante"
	"main/lib/notifiers"
	"main/lib/sessions"
	"strconv"
)

func GetTodos(req *frizzante.Request, res *frizzante.Response) {
	session := frizzante.SessionStart(req, res, sessions.Adapter)
	res.SendView(frizzante.View{Name: "Todos", Data: session.Data.Todos})
}

func PostTodos(req *frizzante.Request, res *frizzante.Response) {
	session := frizzante.SessionStart(req, res, sessions.Adapter)
	form := req.ReceiveForm()

	if form.Has("check") {
		id, intError := strconv.ParseInt(form.Get("check"), 10, 64)
		if nil != intError {
			notifiers.Console.SendError(intError)
			res.SendView(frizzante.View{Name: "Todos", Error: intError})
			return
		}
		session.Data.Todos[id].Checked = true
	} else if form.Has("uncheck") {
		id, intError := strconv.ParseInt(form.Get("uncheck"), 10, 64)
		if nil != intError {
			notifiers.Console.SendError(intError)
			res.SendView(frizzante.View{Name: "Todos", Error: intError})
			return
		}
		session.Data.Todos[id].Checked = false
	}
	session.Save()
	res.SendView(frizzante.View{Name: "Todos", Data: session.Data.Todos})
}
