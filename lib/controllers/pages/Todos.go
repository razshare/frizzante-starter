package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/guards"
	"main/lib/sessions"
	"strconv"
)

var Todos = f.
	NewPageController().
	WithGuard(guards.SessionIsValid).
	WithBase(base).
	WithAction(action)

type Data struct {
	Items []lib.Item `json:"items"`
}

func base(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, sessions.Archived)
	res.SendView(f.NewViewWithData(f.RenderModeFull, Data{
		Items: session.Data.Items,
	}))
}

func action(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, sessions.Archived)
	form := req.ReceiveForm()
	if form.Has("check") {
		index, _ := strconv.ParseInt(form.Get("check"), 10, 32)
		session.Data.Items[index].Checked = true
	} else if form.Has("uncheck") {
		index, _ := strconv.ParseInt(form.Get("uncheck"), 10, 32)
		session.Data.Items[index].Checked = false
	}
	session.Save()
	res.SendView(f.NewViewWithData(f.RenderModeFull, Data{
		Items: session.Data.Items,
	}))
}
