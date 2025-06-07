package main

import (
	"embed"
	"github.com/razshare/frizzante/frz"
	"main/lib/handlers"
)

//go:embed app/dist
var efs embed.FS

func main() {
	frz.NewServer().
		WithEfs(efs).
		WithPublicRoot("app/dist/client").
		WithViewIndex("app/dist/client/index.html").
		WithViewServer("app/dist/server.js").
		AddRoute(frz.Route{Pattern: "GET /", Handler: handlers.Default}).
		AddRoute(frz.Route{Pattern: "GET /welcome", Handler: handlers.Welcome}).
		AddRoute(frz.Route{Pattern: "GET /todos", Handler: handlers.Todos}).
		Start()
}
