package lib

import (
	f "github.com/razshare/frizzante"
	"strconv"
)

func init() {
	Server.
		WithRequestHandler("POST /todos", func(req *f.Request, res *f.Response) {
			session := f.SessionStart(req, res, SessionAdapter)
			form := req.ReceiveForm()
			if form.Has("check") {
				index, _ := strconv.ParseInt(form.Get("check"), 10, 32)
				session.Data.Todos[index].Checked = true
			} else if form.Has("uncheck") {
				index, _ := strconv.ParseInt(form.Get("uncheck"), 10, 32)
				session.Data.Todos[index].Checked = false
			}
			session.Save()
			res.SendView(f.View{
				Name:       "Todos",
				RenderMode: f.RenderModeFull,
				Data:       session.Data,
			})
		}).
		WithRequestHandler("GET /todos", func(req *f.Request, res *f.Response) {
			session := f.SessionStart(req, res, SessionAdapter)
			res.SendView(f.View{
				Name:       "Todos",
				RenderMode: f.RenderModeFull,
				Data:       session.Data,
			})
		}).
		WithRequestHandler("GET /welcome", func(req *f.Request, res *f.Response) {
			res.SendView(f.View{
				Name:       "Welcome",
				RenderMode: f.RenderModeFull,
				Data:       map[string]string{},
			})
		}).
		WithRequestHandler("GET /", func(req *f.Request, res *f.Response) {
			res.SendFileOrElse(func() {
				res.SendView(f.View{
					Name:       "Welcome",
					RenderMode: f.RenderModeFull,
					Data:       map[string]string{},
				})
			})
		})
}
