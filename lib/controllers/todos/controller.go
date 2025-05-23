package todos

import (
	f "github.com/razshare/frizzante"
	"main/lib/config"
	"strconv"
)

var guards = []f.Guard{config.GuardSession}

func init() {
	config.Server.LoadController(func(controller *f.Controller) {
		controller.
			WithBase(guards, base).
			WithAction(guards, action)
	})
}

type data struct {
	Items []config.Todo `json:"items"`
}

func base(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, config.SessionAdapter)
	res.SendView(f.NewViewWithData(f.RenderModeFull, data{
		Items: session.Data.Todos,
	}))
}

func action(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, config.SessionAdapter)
	form := req.ReceiveForm()
	if form.Has("check") {
		index, _ := strconv.ParseInt(form.Get("check"), 10, 32)
		session.Data.Todos[index].Checked = true
	} else if form.Has("uncheck") {
		index, _ := strconv.ParseInt(form.Get("uncheck"), 10, 32)
		session.Data.Todos[index].Checked = false
	}
	session.Save()
	res.SendView(f.NewViewWithData(f.RenderModeFull, data{
		Items: session.Data.Todos,
	}))
}
