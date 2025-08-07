package main

import (
	"embed"
	"github.com/razshare/frizzante/routes"
	"github.com/razshare/frizzante/servers"
	"main/lib/route/handler"
)

//go:embed app/dist
var fs embed.FS
var s = servers.New()

func main() {
	s.Efs = fs

	s.Routes = []routes.Route{
		{Pattern: "GET /", Handler: handler.Default},
		{Pattern: "GET /welcome", Handler: handler.Welcome},
		{Pattern: "GET /todos", Handler: handler.Todos},
		{Pattern: "GET /check", Handler: handler.Check},
		{Pattern: "GET /uncheck", Handler: handler.Uncheck},
		{Pattern: "GET /add", Handler: handler.Add},
		{Pattern: "GET /remove", Handler: handler.Remove},
	}

	servers.Start(s)
}
