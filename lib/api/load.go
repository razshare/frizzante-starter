package api

import (
	"github.com/razshare/frizzante"
	"main/lib/api/routes"
)

func Load(server *frizzante.Server) {
	server.
		WithRequestHandler("GET /events", routes.GetEvents).
		WithRequestHandler("GET /expired", routes.GetExpired).
		WithRequestHandler("GET /todos", routes.GetTodos).
		WithRequestHandler("POST /todos", routes.PostTodos).
		WithRequestHandler("GET /welcome", routes.GetWelcome).
		WithRequestHandler("GET /", func(req *frizzante.Request, res *frizzante.Response) {
			res.SendFileOrElse(func() {
				routes.GetWelcome(req, res)
			})
		})
}
