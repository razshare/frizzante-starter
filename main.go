package main

import (
	"embed"
	"github.com/razshare/frizzante/route"
	"github.com/razshare/frizzante/svelte/container/server"
	"main/lib/routes/handlers/fallback"
	"main/lib/routes/handlers/todos"
	"main/lib/routes/handlers/welcome"
)

//go:embed app/dist
var efs embed.FS
var srv = server.Default(efs)

func main() {
	defer server.Start(srv)
	srv.Routes = []route.Route{
		{Pattern: "GET /", Handler: fallback.View},
		{Pattern: "GET /welcome", Handler: welcome.View},
		{Pattern: "GET /todos", Handler: todos.View},
		{Pattern: "GET /check", Handler: todos.Check},
		{Pattern: "GET /uncheck", Handler: todos.Uncheck},
		{Pattern: "GET /add", Handler: todos.Add},
		{Pattern: "GET /remove", Handler: todos.Remove},
	}
}
