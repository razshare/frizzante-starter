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
		AddRoute(frz.Route{Pattern: "GET /", Handler: handlers.Default}).
		AddRoute(frz.Route{Pattern: "GET /welcome", Handler: handlers.Welcome}).
		AddRoute(frz.Route{Pattern: "GET /todos", Handler: handlers.Todos}).
		AddRoute(frz.Route{Pattern: "GET /check", Handler: handlers.Check}).
		AddRoute(frz.Route{Pattern: "GET /uncheck", Handler: handlers.Uncheck}).
		AddRoute(frz.Route{Pattern: "GET /add", Handler: handlers.Add}).
		AddRoute(frz.Route{Pattern: "GET /remove", Handler: handlers.Remove}).
		Start()
}
