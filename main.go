package main

import (
	"embed"
	"github.com/razshare/frizzante/route"
	"github.com/razshare/frizzante/server"
	"main/lib/route/handler"
)

//go:embed app/dist
var efs embed.FS
var conf = server.Default()

func main() {
	defer server.Start(conf)

	conf.Container.Efs = efs
	conf.Routes = []route.Route{
		{Pattern: "GET /", Handler: handler.Default},
		{Pattern: "GET /welcome", Handler: handler.Welcome},
		{Pattern: "GET /todos", Handler: handler.Todos},
		{Pattern: "GET /check", Handler: handler.Check},
		{Pattern: "GET /uncheck", Handler: handler.Uncheck},
		{Pattern: "GET /add", Handler: handler.Add},
		{Pattern: "GET /remove", Handler: handler.Remove},
	}
}
