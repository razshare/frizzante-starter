package main

import (
	"embed"
	f "github.com/razshare/frizzante"
	"main/lib/sessions"
	"strconv"
)

//go:embed .dist/*/**
var dist embed.FS

func welcome(req *f.Request, res *f.Response) {
	res.SendView(f.View{
		Name:       "Welcome",
		RenderMode: f.RenderModeFull,
		Data:       map[string]string{},
	})
}

func todos(req *f.Request, res *f.Response) {
	session := f.SessionStart(req, res, sessions.Adapter)
	res.SendView(f.View{
		Name:       "Todos",
		RenderMode: f.RenderModeFull,
		Data:       session.Data,
	})
}

func todosPost(req *f.Request, res *f.Response) {
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
}

func main() {
	f.NewServer().
		WithAddress("127.0.0.1:8080").
		WithRequestHandler("POST /todos", todosPost).
		WithRequestHandler("GET /todos", todos).
		WithRequestHandler("GET /welcome", welcome).
		WithRequestHandler("GET /", func(req *f.Request, res *f.Response) {
			res.SendFileOrElse(func() {
				welcome(req, res)
			})
		}).
		Start(dist)
}
