package main

import (
	"embed"
	"github.com/razshare/frizzante/frz"
	"main/lib/handlers"
)

//go:embed dist
var efs embed.FS

func main() {
	frz.NewServer().
		WithEfs(efs).
		WithPublicRoot("dist/client").
		WithViewIndex("dist/client/index.html").
		WithViewServer("dist/server.js").
		AddRoute(frz.Route{Pattern: "GET /", Handler: handlers.GetDefault}).
		AddRoute(frz.Route{Pattern: "GET /welcome", Handler: handlers.GetWelcome}).
		AddRoute(frz.Route{Pattern: "GET /todos", Handler: handlers.GetTodos}).
		AddRoute(frz.Route{Pattern: "POST /todos", Handler: handlers.PostTodos}).
		Start()
}
