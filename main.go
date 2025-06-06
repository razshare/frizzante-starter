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
		AddRoute(frz.Route{Pattern: "GET /", Handler: handlers.Default}).
		AddRoute(frz.Route{Pattern: "GET /welcome", Handler: handlers.Welcome}).
		AddRoute(frz.Route{Pattern: "GET /todos", Handler: handlers.Todos}).
		Start()
}
