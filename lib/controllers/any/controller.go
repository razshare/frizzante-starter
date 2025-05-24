package any

import (
	f "github.com/razshare/frizzante"
	"main/lib/sessions"
	"strconv"
)

type Controller struct {
}

type Data struct {
	Items []sessions.Todo `json:"items"`
}

func (_ Controller) Configure(meta func() f.Metadata) f.PageConfiguration {
	return f.PageConfiguration{
		Metadata: meta(),
		GiveWay:  true,
	}
}

func (_ Controller) Base(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, sessions.Adapter)
	res.SendView(f.NewViewWithData(f.RenderModeFull, Data{
		Items: session.Data.Todos,
	}))
}

func (_ Controller) Action(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, sessions.Adapter)
	form := req.ReceiveForm()
	if form.Has("check") {
		index, _ := strconv.ParseInt(form.Get("check"), 10, 32)
		session.Data.Todos[index].Checked = true
	} else if form.Has("uncheck") {
		index, _ := strconv.ParseInt(form.Get("uncheck"), 10, 32)
		session.Data.Todos[index].Checked = false
	}
	session.Save()
	res.SendView(f.NewViewWithData(f.RenderModeFull, Data{
		Items: session.Data.Todos,
	}))
}
