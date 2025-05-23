package todos

import (
	f "github.com/razshare/frizzante"
	"main/lib/config"
	"strconv"
)

func init() {
	config.Server.LoadPageController(func(controller *f.PageController) {
		controller.
			WithGuard(config.GuardSession).
			WithBase(base).
			WithAction(action)
	})
}

type data struct {
	Items []config.Todo `json:"items"`
}

func base(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, config.SessionLoad)
	res.SendView(f.NewViewWithData(f.RenderModeFull, data{
		Items: session.Data.Todos,
	}))
}

func action(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, config.SessionLoad)
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
