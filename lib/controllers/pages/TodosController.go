package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/guards"
	"main/lib/sessions"
	"strconv"
)

type TodosData struct {
	Items []lib.Item `json:"items"`
}

type TodosController struct {
	f.PageController
}

func (_ TodosController) Configure() f.PageConfiguration {
	return f.PageConfiguration{
		Path: "/todos",
		Guards: []f.GuardFunction{
			guards.SessionIsValid,
		},
	}
}

func (_ TodosController) Base(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, sessions.Archived)

	res.SendView(f.NewViewWithData(f.RenderModeFull, TodosData{
		Items: session.Data.Items,
	}))
}

func (_ TodosController) Action(req *f.Request, res *f.Response) {
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

	res.SendView(f.NewViewWithData(f.RenderModeFull, TodosData{
		Items: session.Data.Items,
	}))
}
