package main

import (
	"embed"
	"github.com/razshare/frizzante/route"
	"github.com/razshare/frizzante/server"
	"main/lib/routes/handlers"
)

//go:embed app/dist
var efs embed.FS
var conf = server.Default()

func main() {
	defer server.Start(conf)
	conf.Container.Efs = efs
	conf.Routes = []route.Route{
		{Pattern: "GET /", Handler: handlers.Fallback},
		{Pattern: "GET /welcome", Handler: handlers.Welcome},
		{Pattern: "GET /todos", Handler: handlers.Todos},
		{Pattern: "GET /check", Handler: handlers.Check},
		{Pattern: "GET /uncheck", Handler: handlers.Uncheck},
		{Pattern: "GET /add", Handler: handlers.Add},
		{Pattern: "GET /remove", Handler: handlers.Remove},
	}
}
