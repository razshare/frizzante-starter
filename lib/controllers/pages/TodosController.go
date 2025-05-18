package pages

import (
	f "github.com/razshare/frizzante"
	"main/lib"
	"main/lib/guards"
	"main/lib/sessions"
	"strconv"
)

type TodosData struct {
	Items []lib.Item
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

func (_ TodosController) Base(request *f.Request, response *f.Response) {
	session := f.SessionStart(request, response, sessions.Archived)

	response.SendView(f.NewView(TodosData{
		Items: session.Data.Items,
	}))
}

func (_ TodosController) Action(request *f.Request, response *f.Response) {
	session := f.SessionStart(request, response, sessions.Archived)

	form := request.ReceiveForm()

	if form.Has("check") {
		index, _ := strconv.ParseInt(form.Get("check"), 10, 32)
		session.Data.Items[index].Checked = true
	} else if form.Has("uncheck") {
		index, _ := strconv.ParseInt(form.Get("uncheck"), 10, 32)
		session.Data.Items[index].Checked = false
	}

	session.Save()

	response.SendView(f.NewView(TodosData{
		Items: session.Data.Items,
	}))
}
