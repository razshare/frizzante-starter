package main

import (
	"embed"
	f "github.com/razshare/frizzante"
	"main/lib/sessions"
	"strconv"
)

//go:embed .dist/*/**
var dist embed.FS

func main() {
	f.NewServer().
		WithAddress("127.0.0.1:8080").
		WithRequestHandler("POST /todos", func(req *f.Request, res *f.Response) {
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
			res.SendView(f.View{
				Name:       "Todos",
				RenderMode: f.RenderModeFull,
				Data:       session.Data,
			})
		}).
		WithRequestHandler("GET /todos", func(req *f.Request, res *f.Response) {
			session := f.SessionStart(req, res, sessions.Adapter)
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
		}).
		Start(dist)
}
