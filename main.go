package main

import (
	"embed"
	"main/lib/handlers"

	"github.com/razshare/frizzante/frz"
)

//go:embed app/dist
var efs embed.FS

func main() {
	frz.NewServer().
		WithEfs(efs).
		WithPublicRoot("app/dist/client").
		WithViewServer("app/dist/server.js").
		WithViewIndex("app/dist/client/index.html").
		AddRoute(frz.Route{Pattern: "GET /", Handler: handlers.Default}).
		AddRoute(frz.Route{Pattern: "GET /welcome", Handler: handlers.Welcome}).
		AddRoute(frz.Route{Pattern: "GET /todos", Handler: handlers.Todos}).
		AddRoute(frz.Route{Pattern: "GET /check", Handler: handlers.Check}).
		AddRoute(frz.Route{Pattern: "GET /uncheck", Handler: handlers.Uncheck}).
		Start()
}
