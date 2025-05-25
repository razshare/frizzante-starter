package lib

import (
	"fmt"
	f "github.com/razshare/frizzante"
	"strconv"
	"time"
)

func init() {
	Server.
		WithRequestHandler("POST /todos", func(req *f.Request, res *f.Response) {
			if !f.AllGuardsPass(req, res, NotExpired) {
				return
			}

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
			if !f.AllGuardsPass(req, res, NotExpired) {
				return
			}

			session := f.SessionStart(req, res, SessionAdapter)
			res.SendView(f.View{
				Name:       "Todos",
				RenderMode: f.RenderModeFull,
				Data:       session.Data,
			})

		}).
		WithRequestHandler("GET /welcome", func(req *f.Request, res *f.Response) {
			if !f.AllGuardsPass(req, res, NotExpired) {
				return
			}

			res.SendView(f.View{
				Name:       "Welcome",
				RenderMode: f.RenderModeFull,
				Data:       map[string]string{},
			})

		}).
		WithRequestHandler("GET /api/events", func(req *f.Request, res *f.Response) {
			if !f.AllGuardsPass(req, res, NotExpired) {
				return
			}

			alive := req.IsAlive()
			res.SendSseUpgrade()
			for *alive {
				now := time.Now().Format(time.TimeOnly)
				message := fmt.Sprintf("Server time is %s", now)
				res.SendMessage(message)
				time.Sleep(time.Second)
			}

		}).
		WithRequestHandler("GET /expired", func(req *f.Request, res *f.Response) {
			res.SendView(f.View{
				Name:       "Expired",
				RenderMode: f.RenderModeFull,
				Data:       map[string]string{},
			})
		}).
		WithRequestHandler("GET /", func(req *f.Request, res *f.Response) {
			if !f.AllGuardsPass(req, res, NotExpired) {
				return
			}

			res.SendFileOrElse(func() {
				res.SendView(f.View{
					Name:       "Welcome",
					RenderMode: f.RenderModeFull,
					Data:       map[string]string{},
				})
			})
		})

	//AllGuardsPass(NotExpired) <- func(req *f.Request, res *f.Response) {
	//
	//}
}
